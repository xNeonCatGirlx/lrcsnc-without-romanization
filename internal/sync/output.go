package sync

import (
	"math"
	"time"

	"lrcsnc/internal/output"
	"lrcsnc/internal/pkg/global"

	"github.com/Endg4meZer0/go-mpris"
)

var lyricsTimer = time.NewTimer(5 * time.Second)
var writtenTimestamp float64

func resyncLyrics() {
	lyricsTimer.Reset(1)
}

func stopLyricsSync() {
	lyricsTimer.Stop()
}

func lyricsSynchronizer() {
	go func() {
		for {
			<-lyricsTimer.C
			if global.Player.P.Song.LyricsData.LyricsType >= 2 {
				output.Controllers[global.Config.C.Global.Output].DisplayLyric(-1)
			} else {
				// 5999.99s is basically the maximum limit of .lrc files' timestamps AFAIK, so 6000s is unreachable
				currentLyricTimestamp := -1.0
				nextLyricTimestamp := 6000.0
				lyricIndex := -1

				for i, lyric := range global.Player.P.Song.LyricsData.Lyrics {
					if lyric.Time+global.Config.C.Lyrics.TimestampOffset <= global.Player.P.Position && currentLyricTimestamp <= lyric.Time+global.Config.C.Lyrics.TimestampOffset {
						currentLyricTimestamp = lyric.Time + global.Config.C.Lyrics.TimestampOffset
						lyricIndex = i
					}
				}

				if lyricIndex != len(global.Player.P.Song.LyricsData.Lyrics)-1 {
					nextLyricTimestamp = global.Player.P.Song.LyricsData.Lyrics[lyricIndex+1].Time + global.Config.C.Lyrics.TimestampOffset
				}

				lyricsTimerDuration := time.Duration(int64(math.Abs(nextLyricTimestamp-global.Player.P.Position)*1000)) * time.Millisecond

				if currentLyricTimestamp == -1 || (global.Player.P.PlaybackStatus == mpris.PlaybackPlaying && writtenTimestamp != currentLyricTimestamp) {
					output.Controllers[global.Config.C.Global.Output].DisplayLyric(lyricIndex)
				}

				writtenTimestamp = currentLyricTimestamp
				global.Player.P.Position = nextLyricTimestamp
				lyricsTimer.Reset(lyricsTimerDuration)
			}
		}
	}()
}
