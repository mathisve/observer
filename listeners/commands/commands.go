package commands

import (
	"github.com/bwmarrin/discordgo"
)

func WoahReply(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Woah!")
}