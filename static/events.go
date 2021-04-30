package static

type LogEvent struct {
	Message   string
	Timestamp int64
}

type LogEventMessage struct {
	ChannelId     string
	GuildId       string
	MessageId     string
	ContentLength int
	Bot           bool
}
