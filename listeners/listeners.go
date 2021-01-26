package listeners

type MessageListener struct {}
func NewMessageListener() *MessageListener {
	return &MessageListener{}
}

type VoiceListener struct {}
func NewVoiceListener() *VoiceListener {
	return &VoiceListener{}
}

type MemberAddListener struct {}
func NewMemberAddListener() *MemberAddListener {
	return &MemberAddListener{}
}