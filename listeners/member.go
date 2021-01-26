package listeners

import (
	"github.com/bwmarrin/discordgo"
	"observerBot/cloud"
	"observerBot/static"
	"time"
)

func (l *MemberAddListener) Handler(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	err := cloud.PutMemberAdd(newMemberAddEvent(e), "")
	if err !=
	return
}

func newMemberAddEvent(e *discordgo.GuildMemberAdd) static.DBMemberAddEvent{
	return static.DBMemberAddEvent{
		GuildId:   e.GuildID,
		Timestamp: time.Now().Unix(),
		Event:     *e.Member,
	}
}