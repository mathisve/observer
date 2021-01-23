package auxilary

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"

	"observerBot/static"
)

// TODO: Change to bulk delete with messages saved to a file (or dynamodb)
func DeleteMessageEventually(s *discordgo.Session, m *discordgo.MessageCreate, tries int) {
	// makes sure it doesn't infinitely try to delete a message
	if tries > static.MAX_DELETE_RETRIES {
		return
	}

	time.Sleep(static.DELETE_AFTER_TIME * time.Second)

	err := s.ChannelMessageDelete(m.ChannelID, m.ID)
	if err != nil {
		if strings.Contains(err.Error(), "10008") {
			// message is already deleted, by user or other bot
			log.Printf(static.MESSAGE_DELETED_ALREADY, m.ID, m.Content, m.Author.Username, m.Author.ID)
			return
		}

		// some error, probably rate-limited
		// so we try again
		log.Printf(static.MESSAGE_DELETED_ERROR, m.ID, m.Content, m.Author.Username, m.Author.ID, err)
		go DeleteMessageEventually(s, m, tries+1)
		return
	}

	log.Printf(static.MESSAGE_DELETED_SUCESSFULLY, m.ID, m.Content, m.Author.Username, m.Author.ID)
}