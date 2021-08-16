package main

import (
	"gus/cloud"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"gus/listeners"
	"gus/static"
)

func init() {
	static.SetStaticVars()
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + static.TOKEN)
	if err != nil {
		log.Println(static.ERROR_CREATING_SESSION, err)
		return
	}

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	dg.AddHandler(listeners.NewMessageListener().Handler)

	err = dg.Open()
	if err != nil {
		log.Println(static.ERROR_OPENING_CONNECTION, err)
		return
	}

	go cloud.PushLogs()

	// Wait here until CTRL-C or other term signal is received.
	log.Println(static.BOT_IS_RUNNING)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
