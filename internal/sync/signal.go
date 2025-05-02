package sync

import (
	"fmt"
	"lrcsnc/internal/mpris"
	"lrcsnc/internal/output"
	"lrcsnc/internal/pkg/global"
	"lrcsnc/internal/pkg/log"

	mprislib "github.com/Endg4meZer0/go-mpris"
)

func mprisMessageReceiver() {
	for msg := range mpris.MPRISMessageChannel {
		log.Debug("sync/mprisMessageReceiver", fmt.Sprintf("Received message: %v", msg))
		switch msg.Type {
		case mpris.SignalReady, mpris.SignalPlayerChanged:
			if global.Player.P.PlaybackStatus != mprislib.PlaybackStopped {
				songChanged <- true
			} else {
				output.Controllers[global.Config.C.Output.Type].DisplayLyric(-1)
			}
		case mpris.SignalSeeked:
			// On seeked signal we just update the position...
			global.Player.M.Lock()
			global.Player.P.Position = msg.Data.(float64) / 1000 / 1000
			global.Player.M.Unlock()
			// ...and, of course, ask for a position sync
			AskForPositionSync()
		case mpris.SignalPlaybackStatusChanged:
			// If the playback status has changed...
			global.Player.M.Lock()
			global.Player.P.PlaybackStatus = msg.Data.(mprislib.PlaybackStatus)
			// ...and if it's stopped then also reset the position
			if global.Player.P.PlaybackStatus == mprislib.PlaybackStopped {
				global.Player.P.Position = 0
			}
			global.Player.M.Unlock()

			outputUpdate()

			// And ask for a position sync to be sure
			AskForPositionSync()

		case mpris.SignalRateChanged:
			// If the rate has changed...
			global.Player.M.Lock()
			global.Player.P.Rate = msg.Data.(float64)
			global.Player.M.Unlock()
			// ...we will absolutely ask for a sync, since the old sync is now invalid
			AskForPositionSync()

		case mpris.SignalMetadataChanged:
			// If the metadata has changed...
			global.Player.M.Lock()
			err := mpris.ApplyMetadataOntoGlobal(msg.Data.(mprislib.Metadata))
			if err != nil {
				log.Error("sync/mprisMessageReceiver", "Couldn't parse metadata: "+err.Error())
			}
			global.Player.M.Unlock()
			// Now we can send a signal to the output module
			// that the info has changed
			output.Controllers[global.Config.C.Output.Type].OnPlayerUpdate()

			// And also send a signal that the song has changed and we need
			// to fetch some new lyrics
			songChanged <- true

			// If the song changed (and we know it did), the position was probably set to 0
			// and we can also ask for a position sync
			AskForPositionSync()
		}
	}
	log.Error("sync/mprisMessageReceiver", "The MPRIS message channel was closed. What happened?")
}
