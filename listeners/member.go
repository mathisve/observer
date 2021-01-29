package listeners

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"observerBot/cloud"
	"observerBot/static"
	"time"
)

func (l *MemberAddListener) Handler(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	err := cloud.PutMemberAdd(newMemberAddEvent(e), "observerMemberAdd")
	if err != nil {
		log.Println(err)
	}

	if static.LOCKDOWN {
		log.Printf("member banned")
		err = s.GuildBanCreateWithReason(e.GuildID, e.User.ID, "banned kekw", 5)
		if err != nil {
			log.Println(err)
		}
	}
}

func newMemberAddEvent(e *discordgo.GuildMemberAdd) static.DBMemberAddEvent{
	return static.DBMemberAddEvent{
		UserId:   e.User.ID,
		Timestamp: time.Now().Unix(),
		Event:     *e.Member,
	}
}