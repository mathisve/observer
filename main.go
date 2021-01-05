package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func init() {
	Token = os.Getenv("TOKEN")
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println(ERROR_CREATING_SESSION, err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		log.Println(ERROR_OPENING_CONNECTION, err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Println(BOT_IS_RUNNING)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.ChannelID == GARBAGE_COLLECTOR_CHANNEL {
		if stringSliceContains(m.Member.Roles, IMMUNE_ROLE_ID) && strings.Contains(m.Content, "--") {
			return
		}

		go DeleteMessageEventually(s, m, 0)
	}

	if m.Author.ID == s.State.User.ID {
		return
	}

	err := PutMessage(m)
	if err != nil {
		log.Print(err)
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
