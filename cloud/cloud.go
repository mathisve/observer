package cloud

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/google/uuid"
	"log"
	"os"
	"sync"
	"time"

	"gus/static"
)

var cwl *cloudwatchlogs.CloudWatchLogs

var sequenceToken string
var logStreamName string

var logQueue []*cloudwatchlogs.InputLogEvent
var logQueueLock sync.Mutex

func init() {

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(os.Getenv("AWS_DEFAULT_REGION")),
		},
	})

	if err != nil {
		log.Panic(err)
	}

	cwl = cloudwatchlogs.New(sess)

}

func PutLogEvent(e static.LogEvent) error {
	event := cloudwatchlogs.InputLogEvent{
		Timestamp: &e.Timestamp,
		Message:   &e.Message,
	}

	logQueueLock.Lock()
	defer logQueueLock.Unlock()

	logQueue = append(logQueue, &event)
	return nil
}

func CreateLogStream() error {

	name := uuid.New().String()

	_, err := cwl.CreateLogStream(&cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  &static.LOG_GROUP_NAME,
		LogStreamName: &name,
	})

	if err != nil {
		return err
	}

	logStreamName = name
	return err
}

func PushLogs() {
	for {
		time.Sleep(time.Second * 15)

		logQueueLock.Lock()

		var input cloudwatchlogs.PutLogEventsInput
		if len(logQueue) > 0 {

			if sequenceToken == "" {
				err := CreateLogStream()
				if err != nil {
					log.Println(err)
				}

				input = cloudwatchlogs.PutLogEventsInput{
					LogEvents:     logQueue,
					LogGroupName:  aws.String(os.Getenv("LOG_GROUP_NAME")),
					LogStreamName: &logStreamName,
				}
			} else {
				input = cloudwatchlogs.PutLogEventsInput{
					LogEvents:     logQueue,
					LogGroupName:  aws.String(os.Getenv("LOG_GROUP_NAME")),
					LogStreamName: &logStreamName,
					SequenceToken: &sequenceToken,
				}
			}

			resp, err := cwl.PutLogEvents(&input)
			if err != nil {
				log.Println(err)
			}

			if resp != nil {
				if resp.NextSequenceToken != nil {
					sequenceToken = *resp.NextSequenceToken
				}
			}

			logQueue = []*cloudwatchlogs.InputLogEvent{}

		}

		logQueueLock.Unlock()

	}
}
