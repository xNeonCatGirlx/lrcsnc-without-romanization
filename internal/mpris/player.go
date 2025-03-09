package mpris

import (
	"slices"
	"strings"

	"lrcsnc/internal/log"
	"lrcsnc/internal/pkg/global"
	"lrcsnc/internal/pkg/types"

	"github.com/Endg4meZer0/go-mpris"
	"github.com/godbus/dbus/v5"
)

func UpdatePlayer() error {
	global.Config.M.Lock()
	defer global.Config.M.Unlock()

	// Getting the players list
	players, err := mpris.List(conn)
	if err != nil {
		log.Error("mpris/UpdatePlayer/List", err.Error())
		return err
	}

	// A little helper function to filter players
	playerInFilter := func(player string) bool {
		if len(global.Config.C.Player.IncludedPlayers) != 0 {
			for _, includedPlayer := range global.Config.C.Player.IncludedPlayers {
				if strings.Contains(player, includedPlayer) {
					return true
				}
			}
		}

		if len(global.Config.C.Player.ExcludedPlayers) != 0 {
			for _, excludedPlayer := range global.Config.C.Player.ExcludedPlayers {
				if strings.Contains(player, excludedPlayer) {
					return false
				}
			}
		}

		return true
	}

	// Check if the current player is still alive and kicking
	if player != nil {
		currentPlayer := player.GetName()
		if slices.Contains(players, currentPlayer) && playerInFilter(currentPlayer) {
			return nil
		} else {
			// Remove the signal handler from the current player before assigning to a new player
			err = player.UnregisterSignalReceiver(PlayerSignalReceiver)
			if err != nil {
				log.Error(
					"mpris/UpdatePlayer/UnregisterSignalReceiver",
					"An error occurred. Reassigning channel. More: "+err.Error(),
				)
				PlayerSignalReceiver = make(chan *dbus.Signal)
			} else {
				log.Debug(
					"mpris/UpdatePlayer/UnregisterSignalReceiver",
					"The current player is no longer alive; successfully unregistered signal receiver.",
				)
			}
		}
	}

	// Find a new player that passes the filters and supports MPRIS to the extent that we need
	for _, p := range players {
		if playerInFilter(p) {
			log.Debug("mpris/UpdatePlayer", "Found new player: '"+p+"', trying to connect...")
			player = mpris.New(conn, p)

			// Lock the player mutex while we're updating the data
			global.Player.M.Lock()

			pass := true

			global.Player.P.PlaybackStatus, err = GetPlaybackStatus()
			if err != nil {
				log.Error("mpris/UpdatePlayer/GetPlaybackStatus", err.Error())
				pass = false
			}
			global.Player.P.Position, err = GetPosition()
			if err != nil {
				log.Error("mpris/UpdatePlayer/GetPosition", err.Error())
				pass = false
			}
			global.Player.P.Rate, err = GetRate()
			if err != nil {
				log.Error("mpris/UpdatePlayer/GetRate", err.Error())
				pass = false
			}
			md, err := GetMetadata()
			if err != nil {
				log.Error("mpris/UpdatePlayer/GetMetadata", err.Error())
				pass = false
			}

			err = ApplyMetadataOntoGlobal(md)
			if err != nil {
				log.Error("mpris/UpdatePlayer/ApplyMetadataOntoGlobal", err.Error())
				pass = false
			}

			global.Player.M.Unlock()

			if !pass {
				log.Debug("mpris/UpdatePlayer", "Failed to gather necessary data from player '"+p+"'. Skipping...")
				continue
			}

			// Register signal receiver for the new player
			err = player.RegisterSignalReceiver(PlayerSignalReceiver)
			if err != nil {
				log.Fatal("mpris/UpdatePlayer/RegisterSignalReceiver", "Cannot watch for player signals. More: "+err.Error())
			}

			log.Debug("mpris/UpdatePlayer", "Successfully connected to player '"+p+"' and gathered necessary data.")
			log.Info("mpris/UpdatePlayer", "Switched to player '"+p+"'")

			return nil
		}
	}

	// If no player is found, set the player to nil
	log.Info("mpris/UpdatePlayer", "No active player found. Zzz")
	player = nil
	return nil
}

func ApplyMetadataOntoGlobal(md mpris.Metadata) (err error) {
	global.Player.P.Song.Title, err = md.Title()
	if err != nil {
		return err
	}
	global.Player.P.Song.Artists, err = md.Artist()
	if err != nil {
		return err
	}
	global.Player.P.Song.Album, err = md.Album()
	if err != nil {
		return err
	}
	dur, err := md.Length()
	if err != nil {
		return err
	}
	global.Player.P.Song.Duration = float64(dur) / 1000 / 1000
	global.Player.P.Song.LyricsData.LyricsType = types.LyricsStateInProgress

	return nil
}

func GetPlaybackStatus() (mpris.PlaybackStatus, error) {
	if player == nil {
		return mpris.PlaybackStopped, nil
	}

	return player.GetPlaybackStatus()
}

func GetPosition() (float64, error) {
	if player == nil {
		return 0, nil
	}

	val, err := player.GetPosition()
	if err != nil {
		return 0, err
	}

	return float64(val) / 1000 / 1000, nil
}

func GetRate() (float64, error) {
	if player == nil {
		return 0, nil
	}

	return player.GetRate()
}

func GetMetadata() (mpris.Metadata, error) {
	if player == nil {
		return nil, nil
	}

	return player.GetMetadata()
}

func SetPosition(pos float64) error {
	if player == nil {
		return nil
	}

	return player.SetPosition(int64(pos * 1000 * 1000))
}
