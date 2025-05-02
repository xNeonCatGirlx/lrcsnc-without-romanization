package sync

import (
	"math"
	"time"

	"lrcsnc/internal/output"
	"lrcsnc/internal/pkg/global"
	"lrcsnc/internal/pkg/types"

	"github.com/Endg4meZer0/go-mpris"
)

var lyricsTimer = time.NewTimer(5 * time.Minute)
var lyricIndex = -1
var writtenTimestamp float64

func resyncLyrics() {
	lyricsTimer.Reset(1)
}

func stopLyricsSync() {
	lyricsTimer.Stop()
}

func outputUpdate() {
	if global.Player.P.Song.LyricsData.LyricsState != types.LyricsStateSynced {
		output.Controllers[global.Config.C.Output.Type].DisplayLyric(-1)
	} else {
		output.Controllers[global.Config.C.Output.Type].DisplayLyric(lyricIndex)
	}
}

func lyricsSynchronizer() {
	for {
		<-lyricsTimer.C
		if global.Player.P.Song.LyricsData.LyricsState != types.LyricsStateSynced {
			output.Controllers[global.Config.C.Output.Type].DisplayLyric(-1)
		} else {
			// 5999.99s is basically the maximum limit of .lrc files' timestamps AFAIK, so 6000s is unreachable
			currentLyricTimestamp := -1.0
			nextLyricTimestamp := 6000.0
			newLyricIndex := -1

			for i, lyric := range global.Player.P.Song.LyricsData.Lyrics {
				if lyric.Time+global.Config.C.Lyrics.TimestampOffset <= global.Player.P.Position && currentLyricTimestamp <= lyric.Time+global.Config.C.Lyrics.TimestampOffset {
					currentLyricTimestamp = lyric.Time + global.Config.C.Lyrics.TimestampOffset
					newLyricIndex = i
				}
			}

			if newLyricIndex != len(global.Player.P.Song.LyricsData.Lyrics)-1 {
				nextLyricTimestamp = global.Player.P.Song.LyricsData.Lyrics[newLyricIndex+1].Time + global.Config.C.Lyrics.TimestampOffset
			}

			lyricsTimerDuration := time.Duration(int64(math.Abs(nextLyricTimestamp-global.Player.P.Position)*1000)) * time.Millisecond

			if currentLyricTimestamp == -1 || (global.Player.P.PlaybackStatus == mpris.PlaybackPlaying && writtenTimestamp != currentLyricTimestamp) {
				output.Controllers[global.Config.C.Output.Type].DisplayLyric(newLyricIndex)
			}

			lyricIndex = newLyricIndex
			writtenTimestamp = currentLyricTimestamp
			global.Player.P.Position = nextLyricTimestamp
			lyricsTimer.Reset(lyricsTimerDuration)
		}
	}
}
