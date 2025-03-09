package sync

import (
	"fmt"
	"time"

	"lrcsnc/internal/log"
	"lrcsnc/internal/lyrics"
	"lrcsnc/internal/output"
	"lrcsnc/internal/pkg/global"
)

var songChanged chan bool = make(chan bool, 1)
var lastDownloadStart time.Time

func lyricFetcher() {
	for {
		<-songChanged

		// This value will change on each new song changed event
		// so if the download takes too long and the song was switched
		// it can just store the necessary data in cache and forget about it
		lastDownloadStart = time.Now()

		go func() {
			thisDownloadStart := lastDownloadStart
			lyricsData, err := lyrics.GetLyricsData(global.Player.P.Song)
			if err != nil {
				log.Error("sync/fetch", fmt.Sprintf("Could not get the lyrics: %v", err))
			}

			if thisDownloadStart != lastDownloadStart {
				return
			}

			lyrics.Configure(&lyricsData)

			global.Player.M.Lock()
			global.Player.P.Song.LyricsData = lyricsData
			global.Player.M.Unlock()

			go output.Controllers[global.Config.C.Global.Output].OnPlayerChange()

			// And finally, it ends with a position sync
			AskForPositionSync()
		}()
	}
}
