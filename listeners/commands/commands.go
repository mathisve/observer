package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type KeyValue struct {
	Key   string
	Value string
}

var KeyValues = []KeyValue{
	{
		Key:   "@yall",
		Value: "@yall",
	},
	{
		Key:   "woah",
		Value: "haow",
	},
	{
		Key:   ":<>",
		Value: "chimken",
	},
	{
		Key:   "chimken",
		Value: ":<>",
	},
	{
		Key:   "banana",
		Value: "hahahaha",
	},
}

// WoahReply replies to a message with "Woah!"
func WoahReply(s *discordgo.Session, m *discordgo.MessageCreate) {

	_, err := s.ChannelMessageSendReply(m.ChannelID, "Woah!", m.Reference())
	if err != nil {
		log.Panicln(err)
	}
}

func SendGif(s *discordgo.Session, m *discordgo.MessageCreate, link string) {
	_, err := s.ChannelMessageSendReply(m.ChannelID, link, m.Reference())
	if err != nil {
		log.Println(err)
	}
}

func AngryRateLimit(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSendReply(m.ChannelID, "Tenor is rate limiting us. please slow down a little", m.Reference())
	if err != nil {
		log.Println(err)
	}
}

func Woopsie(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSendReply(m.ChannelID, "Something related to Tenor went tits up, sowwy", m.Reference())
	if err != nil {
		log.Println(err)
	}
}
