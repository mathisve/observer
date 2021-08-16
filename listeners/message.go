package listeners

import (
	"encoding/json"
	"gus/cloud"
	"gus/cloud/gifs"
	"gus/listeners/commands"
	"gus/static"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	gifPrefix = "gif:"
)

func (l *MessageListener) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	go func() {
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

	}()

	if len(m.Content) > 4 && len(m.Content) < 50 && static.TENOR_API_KEY != "" {
		if m.Content[0:4] == gifPrefix {
			searchTerm := m.Content[4:len(m.Content)]
			results, err := gifs.SearchForGif(searchTerm)
			if err != nil {
				if err.Error() == static.ERROR_TENOR_RATE_LIMIT {
					commands.AngryRateLimit(s, m)
				} else {
					commands.Woopsie(s, m)
				}
			}

			commands.SendGif(s, m, results[rand.Intn(len(results))])
		}
	}

	for _, keyvalue := range commands.KeyValues {
		if strings.Contains(m.Content, keyvalue.Key) {
			_, err := s.ChannelMessageSendReply(m.ChannelID, keyvalue.Value, m.MessageReference)

			if err != nil {
				log.Println(err)
			}

			break
		}

	}
}
