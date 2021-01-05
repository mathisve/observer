package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var dyn *dynamodb.DynamoDB

const (
	TABLE_NAME = "discordObserver"
	REGION     = "eu-central-1"
)

func init() {
	dyn = getDynamo()
}

func getDynamo() *dynamodb.DynamoDB {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(REGION),
		},
	})

	if err != nil {
		log.Panic(err)
	}

	return dynamodb.New(sess)
}

func PutMessage(m *discordgo.MessageCreate) error {
	p, err := dynamodbattribute.MarshalMap(m)
	if err != nil {
		return err
	}

	p["authorId"] = &dynamodb.AttributeValue{S: aws.String(m.Author.ID)}
	p["messageId"] = &dynamodb.AttributeValue{S: aws.String(m.ID)}

	_, err = dyn.PutItem(&dynamodb.PutItemInput{
		Item:      p,
		TableName: aws.String(TABLE_NAME),
	})

	return err
}

func DeleteMessageEventually(s *discordgo.Session, m *discordgo.MessageCreate, tries int) {
	if tries > 3 {
		return
	}

	time.Sleep(5 * time.Second)
	err := s.ChannelMessageDelete(m.ChannelID, m.ID)
	if err != nil {
		if strings.Contains(err.Error(), "10008") {
			log.Printf(MESSAGE_DELETED_ALREADY, m.ID, m.Content, m.Author.Username, m.Author.ID)
			return
		}

		log.Printf(MESSAGE_DELETED_ERROR, m.ID, m.Content, m.Author.Username, m.Author.ID, err)
		go DeleteMessageEventually(s, m, tries+1)
		return
	}

	log.Printf(MESSAGE_DELETED_SUCESSFULLY, m.ID, m.Content, m.Author.Username, m.Author.ID)
}
