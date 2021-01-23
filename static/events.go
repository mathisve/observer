package static

import "github.com/bwmarrin/discordgo"

type DBVoiceEvent struct {
	Userid    string `json:"userId"`
	Timestamp int64 `json:"timestamp"`
	Action    string `json:"action"`
	Event     discordgo.VoiceState `json:"event"`
}

type DBMessageEvent struct {
	Link      string `json:"link"`
	Filename  string `json:"filename"`
	AuthorId  string `json:"authorId"`
	MessageId string `json:"messageId"`
}
