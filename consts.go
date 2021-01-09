package main

const (
	MESSAGE_DELETED_SUCESSFULLY = "Deleted message with ID: %v with Content: %v from User: %v with UserID: %v"
	MESSAGE_DELETED_ERROR       = "Error deleting message with ID: %v with Content: %v from User: %v with UserID: %v, ERROR:%v"
	MESSAGE_DELETED_ALREADY     = "Message already deleted with with ID: %v with Content: %v from User: %v with UserID: %v"

	ERROR_CREATING_SESSION   = "error creating Discord session,"
	ERROR_OPENING_CONNECTION = "error opening connection,"
	BOT_IS_RUNNING           = "Bot is now running"

	CREATING_TABLE = "Creating %v table in %v"

	DELETE_AFTER_TIME  = 60 * 5
	MAX_DELETE_RETRIES = 3
)
