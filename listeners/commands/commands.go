package commands

import (
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

func WoahReply(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Woah!")
}

func SendGif(s *discordgo.Session, m *discordgo.MessageCreate, link string) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		URL: link,
		Image: &discordgo.MessageEmbedImage{
			URL: link,
		},
	})
}

func AngryRateLimit(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Tenor is rate limiting us. please slow down a little")
}

func Woopsie(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Something related to Tenor went tits up, sowwy")
}
