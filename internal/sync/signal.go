package sync

import (
	"lrcsnc/internal/mpris"
	"lrcsnc/internal/output"
	"lrcsnc/internal/pkg/global"

	mprislib "github.com/Endg4meZer0/go-mpris"
	"github.com/godbus/dbus/v5"
)

func signalWatcher() {
	for {
		signal, ok := <-mpris.PlayerSignalReceiver
		if !ok {
			// TODO: logger
			continue
		}
		switch mprislib.GetSignalType(signal) {
		case mprislib.SignalSeeked:
			v, ok := signal.Body[0].(int64)
			if !ok {
				// TODO: logger
				continue
			}

			// On seeked signal we just update the position...
			global.Player.M.Lock()
			global.Player.P.Position = float64(v) / 1000 / 1000
			global.Player.M.Unlock()
			// ...and, of course, ask for a position sync
			AskForPositionSync()
		case mprislib.SignalPropertiesChanged:
			v, ok := signal.Body[1].(map[string]dbus.Variant)
			if !ok {
				// TODO: logger
				continue
			}

			global.Player.M.Lock()

			// If the playback status has changed...
			playbackStatus, ok := v["PlaybackStatus"]
			if ok {
				global.Player.P.PlaybackStatus = mprislib.PlaybackStatus(playbackStatus.Value().(string))

				// ...and if it's stopped then also reset the position
				if playbackStatus.Value().(string) == string(mprislib.PlaybackStopped) {
					global.Player.P.Position = 0
				}
				global.Player.M.Unlock()

				// And ask for a position sync to be sure
				AskForPositionSync()
			}

			// If the rate has changed...
			rate, ok := v["Rate"]
			if ok {
				global.Player.P.Rate = rate.Value().(float64)
				global.Player.M.Unlock()

				// ...we will absolutely ask for a sync, since the old sync is now invalid
				AskForPositionSync()
			}

			// If the metadata has changed...
			metadata, ok := v["Metadata"]
			if ok {
				// ...that could mean only one thing - the song changed
				md := mprislib.Metadata(metadata.Value().(map[string]dbus.Variant))

				// Apply the changes onto the global player struct
				err := mpris.ApplyMetadataOntoGlobal(md)
				if err != nil {
					// TODO: logger
					continue
				}
				global.Player.M.Unlock()

				// Now we can send a signal to the output module
				// that the info has changed
				output.Controllers[global.Config.C.Global.Output].OnPlayerChange()

				// And also send a signal that the song has changed and we need
				// to fetch some new lyrics
				songChanged <- true

				// If the song changed (and we know it did), the position was probably set to 0
				// and we can also ask for a position sync
				AskForPositionSync()
			}
		}
	}
}
