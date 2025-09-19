package core

type Channel string

const (
	ChannelSMS   Channel = "SMS"
	ChannelEmail Channel = "EMAIL"
	ChannelInApp Channel = "INAPP"
)
