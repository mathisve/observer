package cloud

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/bwmarrin/discordgo"
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
	err = ensureTableExists()
	if err != nil {
		log.Panic(err)
	}

	lam = lambda.New(sess)

}


func ensureTableExists() error {
	_, descrerr := dyn.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(static.TABLE_NAME),
	})

	if descrerr != nil {
		if aerr, ok := descrerr.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceNotFoundException:
				// table doesn't exist
				// creating one
				log.Printf(static.CREATING_TABLE, static.TABLE_NAME, static.REGION)

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
					TableName:   aws.String(static.TABLE_NAME),
				})
				if err != nil {
					return err
				}

				// wait until the table is actually available
				err = dyn.WaitUntilTableExists(&dynamodb.DescribeTableInput{
					TableName: aws.String(static.TABLE_NAME),
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
		TableName: aws.String(static.TABLE_NAME),
	})

	return err
}



func InvokeLambda(e static.DBMessageEvent) error {
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