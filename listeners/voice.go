package listeners

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"observerBot/cloud"
	"observerBot/static"
	"time"
)

var voiceStateCache = map[string]discordgo.VoiceState{}

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
	var emptyVoiceState = discordgo.VoiceState{}
	if old == emptyVoiceState {
		event := newVoiceStateEvent(e.VoiceState, voiceJoinAction)
		err := cloud.PutVoiceEvent(event, static.VOICE_EVENTS_TABLE_NAME)
		if err != nil {
			log.Println(err)
		}
		return
	}

	// User joined channel
	if old.ChannelID == "" {
		event := newVoiceStateEvent(e.VoiceState, voiceJoinAction)
		err := cloud.PutVoiceEvent(event, static.VOICE_EVENTS_TABLE_NAME)
		if err != nil {
			log.Println(err)
		}
		return
	}

	// User left channel
	if old.ChannelID != "" && cur.ChannelID == "" {
		event := newVoiceStateEvent(e.VoiceState, voiceLeaveAction)
		err := cloud.PutVoiceEvent(event, static.VOICE_EVENTS_TABLE_NAME)
		if err != nil {
			log.Println(err)
		}
		return
	}

	// User switched channel
	if old.ChannelID != "" && cur.ChannelID != "" {
		event := newVoiceStateEvent(e.VoiceState, voiceSwitchAction)
		err := cloud.PutVoiceEvent(event, static.VOICE_EVENTS_TABLE_NAME)
		if err != nil {
			log.Println(err)
		}
		return
	}
}

func updateVoiceStateCache(e *discordgo.VoiceStateUpdate) {
	voiceStateCache[e.UserID] = *e.VoiceState
}

func newVoiceStateEvent(e *discordgo.VoiceState, action string) static.DBVoiceEvent {
	return static.DBVoiceEvent{
		UserId:    e.UserID,
		Timestamp: time.Now().Unix(),
		Action:    action,
		Event:     *e,
	}
}
