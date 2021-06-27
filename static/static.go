package static

import "os"

var (
	TOKEN          string
	REGION         string
	LOG_GROUP_NAME string
)

const (
	MESSAGE_DELETED_SUCESSFULLY = "Deleted message with ID: %v with Content: %v from User: %v with UserID: %v"
	MESSAGE_DELETED_ERROR       = "Error deleting message with ID: %v with Content: %v from User: %v with UserID: %v, ERROR:%v"
	MESSAGE_DELETED_ALREADY     = "Message already deleted with with ID: %v with Content: %v from User: %v with UserID: %v"

	ERROR_CREATING_SESSION   = "error creating Discord session,"
	ERROR_OPENING_CONNECTION = "error opening connection,"
	BOT_IS_RUNNING           = "Bot is now running"

	DELETE_AFTER_TIME  = 60 * 5
	MAX_DELETE_RETRIES = 3

	LINK_REGEX = `(http|https)\:\/\/[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,3}(\/\S*)?`
)

func SetStaticVars() {
	TOKEN = os.Getenv("TOKEN")
	REGION = os.Getenv("AWS_DEFAULT_REGION")
	LOG_GROUP_NAME = os.Getenv("LOG_GROUP_NAME")
}
