package commands

import (
	"github.com/bwmarrin/discordgo"
	"observerBot/lockdown"
	"observerBot/static"
)

func WoahReply(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Woah!")
}

func Lockdown(s *discordgo.Session, m *discordgo.MessageCreate) {
 	switch static.LOCKDOWN {
	case true:
		go lockdown.ReleaseLockdown()
		s.ChannelMessageSend(m.ChannelID, "Lockdown released")
	case false:
		go lockdown.SetLockdown()
		s.ChannelMessageSend(m.ChannelID, "Lockdown activated")
 	}
}

func muteEverybody() {

}

func unmuteEverybody() {

}