package global

import (
	"lrcsnc/internal/pkg/structs"
	"sync"
)

var CurrentSong = struct {
	Mutex sync.Mutex
	Song  structs.Song
}{
	Song: structs.Song{
		LyricsData: structs.LyricsData{
			LyricsType: structs.LyricsStateUnknown,
		},
	},
}
