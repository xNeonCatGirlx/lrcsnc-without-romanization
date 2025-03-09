package structs

import (
	"lrcsnc/internal/pkg/types"

	"github.com/Endg4meZer0/go-mpris"
)

type Player struct {
	PlaybackStatus mpris.PlaybackStatus
	Position       float64
	Rate           float64
	Song           Song
}

type Song struct {
	Title      string
	Artist     string
	Album      string
	Duration   float64
	LyricsData LyricsData
}

type LyricsData struct {
	Lyrics     []Lyric
	LyricsType types.LyricsState
}

type Lyric struct {
	Time float64
	Text string
}
