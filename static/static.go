package static

import "os"

var (
	FUNCTION                  string
	TOKEN                     string
	REGION                    string
	MESSAGE_TABLE_NAME        string
	VOICE_EVENTS_TABLE_NAME   string
	MEMBER_ADD_TABLE_NAME     string
	GARBAGE_COLLECTOR_CHANNEL string
	IMMUNE_ROLE_ID            string
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
	FUNCTION = os.Getenv("FUNCTION")
	REGION = os.Getenv("AWS_DEFAULT_REGION")
	MESSAGE_TABLE_NAME = os.Getenv("MESSAGE_TABLE_NAME")
	VOICE_EVENTS_TABLE_NAME = os.Getenv("VOICE_EVENTS_TABLE_NAME")
	MEMBER_ADD_TABLE_NAME = os.Getenv("MEMBER_ADD_TABLE_NAME")
	GARBAGE_COLLECTOR_CHANNEL = os.Getenv("GARBAGE")
	IMMUNE_ROLE_ID = os.Getenv("IMMUNE")
}
