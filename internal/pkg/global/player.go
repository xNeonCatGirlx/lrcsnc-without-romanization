package global

import (
	"lrcsnc/internal/pkg/structs"
	"lrcsnc/internal/pkg/types"
	"sync"

	"github.com/Endg4meZer0/go-mpris"
)

var Player = struct {
	M sync.Mutex
	P structs.Player
}{
	P: structs.Player{
		PlaybackStatus: mpris.PlaybackStopped,
		Position:       0.0,
		Rate:           1.0,
		Song: structs.Song{
			LyricsData: structs.LyricsData{
				LyricsType: types.LyricsStateUnknown,
			},
		},
	},
}
