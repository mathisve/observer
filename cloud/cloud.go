package cloud

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/lambda"
	"log"
	"os"

	"observerBot/static"
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
	if err != nil {
		log.Panic(err)
	}

	lam = lambda.New(sess)

}

func PutMessageEvent(event static.DBMessageEvent, tableName string) error {
	p, err := dynamodbattribute.MarshalMap(event)
	if err != nil {
		return err
	}

	_, err = dyn.PutItem(&dynamodb.PutItemInput{
		Item:      p,
		TableName: aws.String(tableName),
	})

	return err
}

func PutVoiceEvent(event static.DBVoiceEvent, tableName string) error {
	p, err := dynamodbattribute.MarshalMap(event)
	if err != nil {
		return err
	}

	_, err = dyn.PutItem(&dynamodb.PutItemInput{
		Item:      p,
		TableName: aws.String(tableName),
	})

	return err
}

func PutMemberAdd(event static.DBMemberAddEvent, tableName string) error {
	p, err := dynamodbattribute.MarshalMap(event)
	if err != nil {
		return err
	}

	_, err = dyn.PutItem(&dynamodb.PutItemInput{
		Item:     p,
		TableName: aws.String(tableName),
	})

	return err
}

func InvokeLambda(e static.DBAttachmentEvent) error {
	payload, err := json.Marshal(e)
	if err != nil {
		return err
	}

	_, err = lam.Invoke(&lambda.InvokeInput{
		FunctionName:   aws.String(static.FUNCTION),
		InvocationType: aws.String(lambda.InvocationTypeRequestResponse),
		Payload:        payload,
	})

	return err
}
