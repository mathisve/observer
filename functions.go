package main

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var dyn *dynamodb.DynamoDB
var lam *lambda.Lambda

func init() {

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(os.Getenv("AWS_DEFAULT_REGION")),
		},
	})

	if err != nil {
		log.Panic(err)
	}

	dyn = dynamodb.New(sess)
	err = ensureTableExists()
	if err != nil {
		log.Panic(err)
	}

	lam = lambda.New(sess)

}

func ensureTableExists() error {
	_, descrerr := dyn.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(TABLE_NAME),
	})

	if descrerr != nil {
		if aerr, ok := descrerr.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceNotFoundException:
				// table doesn't exist
				// creating one
				log.Printf(CREATING_TABLE, TABLE_NAME, REGION)

				_, err := dyn.CreateTable(&dynamodb.CreateTableInput{
					AttributeDefinitions: []*dynamodb.AttributeDefinition{
						{
							AttributeName: aws.String("authorId"),
							AttributeType: aws.String("S"),
						},
						{
							AttributeName: aws.String("authorId"),
							AttributeType: aws.String("S"),
						},
					},
					KeySchema: []*dynamodb.KeySchemaElement{
						{
							AttributeName: aws.String("authorId"),
							KeyType:       aws.String("HASH"),
						},
						{
							AttributeName: aws.String("messageId"),
							KeyType:       aws.String("RANGE"),
						},
					},
					BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
					TableName:   aws.String(TABLE_NAME),
				})
				if err != nil {
					return err
				}

				// wait until the table is actually available
				err = dyn.WaitUntilTableExists(&dynamodb.DescribeTableInput{
					TableName: aws.String(TABLE_NAME),
				})
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func PutMessage(m *discordgo.MessageCreate) error {
	p, err := dynamodbattribute.MarshalMap(m)
	if err != nil {
		return err
	}

	// Primary Partition Key
	p["authorId"] = &dynamodb.AttributeValue{S: aws.String(m.Author.ID)}
	// Primary Sort Key
	p["messageId"] = &dynamodb.AttributeValue{S: aws.String(m.ID)}

	_, err = dyn.PutItem(&dynamodb.PutItemInput{
		Item:      p,
		TableName: aws.String(TABLE_NAME),
	})

	return err
}

// TODO: Change to bulk delete with messages saved to a file (or dynamodb)
func DeleteMessageEventually(s *discordgo.Session, m *discordgo.MessageCreate, tries int) {
	// makes sure it doesn't infinitely try to delete a message
	if tries > MAX_DELETE_RETRIES {
		return
	}

	time.Sleep(DELETE_AFTER_TIME * time.Second)

	err := s.ChannelMessageDelete(m.ChannelID, m.ID)
	if err != nil {
		if strings.Contains(err.Error(), "10008") {
			// message is already deleted, by user or other bot
			log.Printf(MESSAGE_DELETED_ALREADY, m.ID, m.Content, m.Author.Username, m.Author.ID)
			return
		}

		// some error, probably rate-limited
		// so we try again
		log.Printf(MESSAGE_DELETED_ERROR, m.ID, m.Content, m.Author.Username, m.Author.ID, err)
		go DeleteMessageEventually(s, m, tries+1)
		return
	}

	log.Printf(MESSAGE_DELETED_SUCESSFULLY, m.ID, m.Content, m.Author.Username, m.Author.ID)
}

func InvokeLambda(e Event) error {
	payload, err := json.Marshal(e)
	if err != nil {
		return err
	}

	_, err = lam.Invoke(&lambda.InvokeInput{
		FunctionName:   aws.String(FUNCTION),
		InvocationType: aws.String(lambda.InvocationTypeRequestResponse),
		Payload:        payload,
	})

	return err
}
