package main

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"log"
	"net/http"
)

func getContent(url string) (bytes []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return bytes, err
	}

	defer resp.Body.Close()

	bytes, err = ioutil.ReadAll(resp.Body)
	if len(bytes) >= 50000000 {
		return []byte(""), errors.New("filesize too big")
	}

	return bytes, err
}

func hash(b []byte) []byte {
	byteArray := sha256.Sum256(b)
	return byteArray[:]
}

func ToString(b []byte) string {
	return fmt.Sprintf("%x", b)
}

func uploadToS3(filename string, b []byte) error {
	_, err := s3session.PutObject(&s3.PutObjectInput{
		Key:    aws.String(filename),
		Body:   bytes.NewReader(b),
		Bucket: aws.String(BUCKET),
	})

	return err
}

func uploadToDynamo(dbEntry DBEntry) error {
	av, err := dynamodbattribute.MarshalMap(dbEntry)
	if err != nil {
		return err
	}

	_, err = dyn.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(TABLE),
		Item:      av,
	})

	return err
}

func ensureBucketExists() error {
	resp, err := s3session.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return err
	}

	var found = false
	for _, bucket := range resp.Buckets {
		if *bucket.Name == BUCKET {
			found = true
		}
	}

	if found == false {
		log.Printf(CREATING_BUCKET, BUCKET, REGION)
		_, err = s3session.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(BUCKET),
			CreateBucketConfiguration: &s3.CreateBucketConfiguration{
				LocationConstraint: aws.String(REGION),
			},
		})

		if err != nil {
			log.Println(ERR_CREATING_BUCKET, BUCKET, REGION)
			return err
		}

		err = s3session.WaitUntilBucketExists(&s3.HeadBucketInput{
			Bucket: aws.String(BUCKET),
		})
	}
	return err
}

func ensureTableExists() error {
	_, descrerr := dyn.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(TABLE),
	})

	if descrerr != nil {
		if aerr, ok := descrerr.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceNotFoundException:
				// table doesn't exist
				// creating one
				log.Printf(CREATING_TABLE, TABLE, REGION)

				_, err := dyn.CreateTable(&dynamodb.CreateTableInput{
					AttributeDefinitions: []*dynamodb.AttributeDefinition{
						{
							AttributeName: aws.String("authorId"),
							AttributeType: aws.String("S"),
						},
						{
							AttributeName: aws.String("messageId"),
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
					BillingMode: aws.String("PAY_PER_REQUEST"),
					TableName:   aws.String(TABLE),
				})
				if err != nil {
					log.Println(ERR_CREATING_TABLE, TABLE, REGION)
					return err
				}

				// wait until the table is actually available
				err = dyn.WaitUntilTableExists(&dynamodb.DescribeTableInput{
					TableName: aws.String(TABLE),
				})
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
