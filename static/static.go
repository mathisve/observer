package static

import "os"

var (
	TOKEN          string
	REGION         string
	LOG_GROUP_NAME string
	TENOR_API_KEY  string
)

const (
	MESSAGE_DELETED_SUCESSFULLY = "Deleted message with ID: %v with Content: %v from User: %v with UserID: %v"
	MESSAGE_DELETED_ERROR       = "Error deleting message with ID: %v with Content: %v from User: %v with UserID: %v, ERROR:%v"
	MESSAGE_DELETED_ALREADY     = "Message already deleted with with ID: %v with Content: %v from User: %v with UserID: %v"

	ERROR_CREATING_SESSION   = "error creating Discord session,"
	ERROR_OPENING_CONNECTION = "error opening connection,"
	BOT_IS_RUNNING           = "Bot is now running"

	ERROR_TENOR_RATE_LIMIT = "tenor is being ratelimited. please wait 30 seconds"
	ERROR_TENOR_NO_GIFS    = "tenor didn't send any gifs"

	DELETE_AFTER_TIME  = 60 * 5
	MAX_DELETE_RETRIES = 3

	LINK_REGEX = `(http|https)\:\/\/[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,3}(\/\S*)?`

	TENOR_URL = "https://g.tenor.com/v1/search?q=%s&key=%s&limit=8"

	ALPHABET = "abcdefghijklmnopqrstuvwxyz"
)

func SetStaticVars() {
	TOKEN = os.Getenv("TOKEN")
	REGION = os.Getenv("AWS_DEFAULT_REGION")
	LOG_GROUP_NAME = os.Getenv("LOG_GROUP_NAME")
	TENOR_API_KEY = os.Getenv("TENOR_API_KEY")
}
