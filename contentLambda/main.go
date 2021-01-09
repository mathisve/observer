package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"os"
	"strings"
)

var REGION, BUCKET, TABLE string
var s3session *s3.S3
var dyn *dynamodb.DynamoDB
var s3InitError, dynInitError error

type Event struct {
	Link      string `json:"link"`
	Filename  string `json:"filename"`
	AuthorId  string `json:"authorId"`
	MessageId string `json:"messageId"`
}

type DBEntry struct {
	Link       string `json:"link"`
	Filename   string `json:"filename"`
	S3Filename string `json:"s3filename"`
	Hash       string `json:"hash"`
	S3Link     string `json:"s3link"`
	AuthorId   string `json:"authorId"`
	MessageId  string `json:"messageId"`
}

func init() {
	REGION = os.Getenv("REGION")
	BUCKET = os.Getenv("BUCKET")
	TABLE = os.Getenv("TABLE")

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	}))

	s3session = s3.New(sess)
	dyn = dynamodb.New(sess)

	err := ensureBucketExists()
	if err != nil {
		s3InitError = err
	}

	err = ensureTableExists()
	if err != nil {
		dynInitError = err
	}
}

func main() {
	lambda.Start(handler)
}

func handler(event Event) (string, error) {
	if s3InitError != nil {
		// log any S3 related errors that occurred in init function
		log.Println(s3InitError)
	}

	if dynInitError != nil {
		// log any DynamoDB related errors that occurred in init function
		log.Println(dynInitError)
	}

	bytes, err := getContent(event.Link)
	if err != nil {
		log.Println(err)
		return "", err
	}

	bytesHash := hash(bytes)

	split := strings.Split(event.Filename, ".")
	s3filename := fmt.Sprintf("%v.%v", ToString(bytesHash), split[len(split)-1])

	var dbEntry = DBEntry{
		Link:       event.Link,
		Filename:   event.Filename,
		S3Filename: s3filename,
		Hash:       ToString(bytesHash),
		S3Link:     fmt.Sprintf("https://%s.%s.amazonaws.com/%s", BUCKET, REGION, s3filename),
		AuthorId:   event.AuthorId,
		MessageId:  event.MessageId,
	}

	log.Printf("%+v", dbEntry)

	err = uploadToS3(s3filename, bytesHash)
	if err != nil {
		log.Println(err)
		return "", err
	}

	err = uploadToDynamo(dbEntry)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return "", err
}
