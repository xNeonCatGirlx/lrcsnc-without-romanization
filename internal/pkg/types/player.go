package types

type LyricsState byte

const (
	LyricsStateSynced       LyricsState = 0
	LyricsStatePlain        LyricsState = 1
	LyricsStateInstrumental LyricsState = 2
	LyricsStateNotFound     LyricsState = 3
	LyricsStateInProgress   LyricsState = 4
	LyricsStateUnknown      LyricsState = 5
)
