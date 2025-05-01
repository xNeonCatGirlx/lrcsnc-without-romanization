package mpris

type SignalType uint8

const (
	SignalReady SignalType = iota
	SignalSeeked
	SignalPlaybackStatusChanged
	SignalRateChanged
	SignalMetadataChanged
	SignalPlayerChanged
)

type Message struct {
	Type SignalType
	Data any
}

var MPRISMessageChannel = make(chan Message)
