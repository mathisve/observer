package listeners

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"gus/cloud"
	"gus/static"
	"strings"
	"log"
	"time"
)

func (l *MessageListener) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, "@yall") {
		s.ChannelMessageSend(m.ChannelID, "@yall")
	}

	if strings.Contains(m.Content, "woah") {
		s.ChannelMessageSend(m.ChannelID, "woah")
	}

	logMsg := static.LogEventMessage{
		ChannelId:     m.ChannelID,
		GuildId:       m.GuildID,
		MessageId:     m.ID,
		ContentLength: len(m.Content),
		Bot:           m.Author.Bot,
	}

	msg, err := json.Marshal(logMsg)
	if err != nil {
		log.Println(err)
	}

	d := static.LogEvent{
		Message:   string(msg),
		Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
	}

	err = cloud.PutLogEvent(d)
	if err != nil {
		log.Println(err)
	}

}
