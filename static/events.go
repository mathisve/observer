package static

import "github.com/bwmarrin/discordgo"

type DBVoiceEvent struct {
	UserId    string               `json:"userId"`
	Timestamp int64                `json:"timestamp"`
	Action    string               `json:"action"`
	Event     discordgo.VoiceState `json:"event"`
}

type DBMessageEvent struct {
	AuthorId  string `json:"authorId"`
	MessageId string `json:"messageId"`
	Timestamp int64  `json:"timestamp"`
	Event     discordgo.MessageCreate
}

type DBAttachmentEvent struct {
	Link      string `json:"link"`
	Filename  string `json:"filename"`
	AuthorId  string `json:"authorId"`
	MessageId string `json:"messageId"`
}

type DBMemberAddEvent struct {
	UserId   string `json:"guildId"`
	Timestamp int64  `json:"timestamp"`
	Event     discordgo.Member `json:"event"`
}
