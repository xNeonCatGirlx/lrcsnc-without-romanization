package sync

import (
	"errors"
	"time"

	"lrcsnc/internal/lyrics"
	"lrcsnc/internal/output"
	errs "lrcsnc/internal/pkg/errors"
	"lrcsnc/internal/pkg/global"
)

var songChanged chan bool = make(chan bool)
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

			lyricsData, err := lyrics.Fetch()
			if err != nil && !errors.Is(err, errs.ErrLyricsNotFound) {
				return
			}

			if thisDownloadStart != lastDownloadStart {
				return
			}

			lyrics.Configure(&lyricsData)

			global.Player.M.Lock()
			global.Player.P.Song.LyricsData = lyricsData
			global.Player.M.Unlock()

			go output.Controllers[global.Config.C.Output.Type].OnPlayerUpdate()

			// And finally, it ends with a position sync
			AskForPositionSync()
		}()
	}
}
