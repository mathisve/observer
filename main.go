package main

import (
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	FUNCTION string
	TOKEN    string
	REGION	string
	TABLE_NAME	string
	GARBAGE_COLLECTOR_CHANNEL string
	IMMUNE_ROLE_ID string
)

func init() {
	TOKEN = os.Getenv("TOKEN")
	FUNCTION = os.Getenv("FUNCTION")
	REGION = os.Getenv("AWS_DEFAULT_REGION")
	TABLE_NAME = os.Getenv("TABLE_NAME")
	GARBAGE_COLLECTOR_CHANNEL = os.Getenv("GARBAGE")
	IMMUNE_ROLE_ID = os.Getenv("IMMUNE")

	log.Println(REGION)
}

type Event struct {
	Link      string `json:"link"`
	Filename  string `json:"filename"`
	AuthorId  string `json:"authorId"`
	MessageId string `json:"messageId"`
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + TOKEN)
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

	// Don't log other bots' messages
	if m.Author.Bot == false {
		err := PutMessage(m)
		if err != nil {
			log.Print(err)
		}

		re := regexp.MustCompile(`(http|https)\:\/\/[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,3}(\/\S*)?`)
		contentLinks := re.FindAllString(m.Content,-1)

		for _, link := range contentLinks {
			split := strings.Split(link, "/")
			event := Event{
				Link:      link,
				Filename:  split[len(split)-1],
				AuthorId:  m.Author.ID,
				MessageId: m.ID,
			}

			err = InvokeLambda(event)
			if err != nil {
				log.Println(err)
			}
		}



		for _, attachment := range m.Attachments {
			event := Event{
				Link:      attachment.URL,
				Filename:  attachment.Filename,
				AuthorId:  m.Author.ID,
				MessageId: m.ID,
			}

			err = InvokeLambda(event)
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