package listeners

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"observerBot/auxilary"
	"observerBot/cloud"
	"regexp"
	"strings"
	"observerBot/static"
	"time"
)

type MessageListener struct {}

func NewMessageListener() *MessageListener {
	return &MessageListener{}
}

func (l *MessageListener) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.ChannelID == static.GARBAGE_COLLECTOR_CHANNEL {
		if stringSliceContains(m.Member.Roles, static.IMMUNE_ROLE_ID) && strings.Contains(m.Content, "--") {
			return
		}

		go auxilary.DeleteMessageEventually(s, m, 0)
	}

	if m.Author.ID == s.State.User.ID {
		return
	}

	messageEvent := newMessageEvent(m)

	// Don't log other bots' messages
	if m.Author.Bot == false {
		err := cloud.PutMessageEvent(messageEvent, static.MESSAGE_TABLE_NAME)
		if err != nil {
			log.Print(err)
		}

		re := regexp.MustCompile(static.LINK_REGEX)
		contentLinks := re.FindAllString(m.Content, -1)

		for _, link := range contentLinks {
			split := strings.Split(link, "/")
			event := static.DBAttachmentEvent{
				Link:      link,
				Filename:  split[len(split)-1],
				AuthorId:  m.Author.ID,
				MessageId: m.ID,
			}

			err = cloud.InvokeLambda(event)
			if err != nil {
				log.Println(err)
			}
		}

		for _, attachment := range m.Attachments {
			event := static.DBAttachmentEvent{
				Link:      attachment.URL,
				Filename:  attachment.Filename,
				AuthorId:  m.Author.ID,
				MessageId: m.ID,
			}

			err = cloud.InvokeLambda(event)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func stringSliceContains(sl []string, s string) bool {
	for _, e := range sl {
		if e == s {
			return true
		}
	}

	return false
}

func newMessageEvent(m *discordgo.MessageCreate) static.DBMessageEvent {
	return static.DBMessageEvent{
		AuthorId:  m.Author.ID,
		MessageId: m.ID,
		Timestamp: time.Now().Unix(),
		Event:     *m,
	}
}