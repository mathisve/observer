package main

const (
	MESSAGE_DELETED_SUCESSFULLY = "Deleted message with ID: %v with Content: %v from User: %v with UserID: %v"
	MESSAGE_DELETED_ERROR       = "Error deleting message with ID: %v with Content: %v from User: %v with UserID: %v, ERROR:%v"
	MESSAGE_DELETED_ALREADY     = "Message already deleted with with ID: %v with Content: %v from User: %v with UserID: %v"

	// Deletes messages in this channel after 5 minutes
	GARBAGE_COLLECTOR_CHANNEL = "795657444058464286"
	// Unless users with this roles' message starts with --
	// will be deleted:  hello
	// won't be deleted: --hello
	IMMUNE_ROLE_ID           = "795656423257669642"

	ERROR_CREATING_SESSION   = "error creating Discord session,"
	ERROR_OPENING_CONNECTION = "error opening connection,"
	BOT_IS_RUNNING           = "Bot is now running"

	CREATING_TABLE = "Creating %v table in %v"
)
