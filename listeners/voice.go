package listeners

import (
	"github.com/bwmarrin/discordgo"
	"observerBot/static"
	"time"
)

var voiceStateCache = map[string]*discordgo.VoiceState{}

type VoiceListener struct {}

const (
	voiceJoinAction   = "JOINED"
	voiceLeaveAction  = "LEFT"
	voiceSwitchAction = "SWITCHED"
)

func NewVoiceListener() *VoiceListener {
	return &VoiceListener{}
}

func (l *VoiceListener) Handler(s *discordgo.Session, e *discordgo.VoiceStateUpdate) {
	defer updateVoiceStateCache(e)

	// current voice update
	cur := e.VoiceState

	// previous voice state update
	old := voiceStateCache[e.UserID]

	// User joined channel
	if old == new(discordgo.VoiceState) {
		NewBDEvent(e.VoiceState, voiceJoinAction)
		return
	}

	// User joined channel
	if old.ChannelID == "" {
		NewBDEvent(e.VoiceState, voiceJoinAction)
		return
	}

	// User left channel
	if old.ChannelID != "" && cur.ChannelID == "" {
		NewBDEvent(e.VoiceState, voiceLeaveAction)
		return
	}

	// User switched channel
	if old.ChannelID != "" && cur.ChannelID != "" {
		NewBDEvent(e.VoiceState, voiceSwitchAction)
		return
	}
}

func updateVoiceStateCache(e *discordgo.VoiceStateUpdate) {
	voiceStateCache[e.UserID] = e.VoiceState
}

func NewBDEvent(e *discordgo.VoiceState, action string) static.DBVoiceEvent {
	return static.DBVoiceEvent{
		Userid:    e.UserID,
		Timestamp: time.Now().Unix(),
		Action:    action,
		Event:     *e,
	}
}
