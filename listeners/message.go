package listeners

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"gus/cloud"
	"gus/static"
	"log"
	"strings"
	"time"
)

type KeyValue struct {
	key string
	value string
}

func (l *MessageListener) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	var keyValues = []KeyValue{
		{
			key:   "@yall",
			value: "@yall",
		},
		{
			key:   "woah",
			value: "haow",
		},
		{
			key: ":<>",
			value: "chimken",
		},
		{
			key: "chimken",
			value: ":<>",
		},
		{
			key: "banana",
			value: "hahahaha",
		},
		{
			key: "observer",
			value: "https://github.com/mathisve/observer",
		},
	}

	for _, keyvalue := range keyValues {

		if strings.Contains(m.Content, keyvalue.key) {
			_, err := s.ChannelMessageSend(m.ChannelID, keyvalue.value)
			if err != nil {
				log.Println(err)
			}

			break
		}


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
