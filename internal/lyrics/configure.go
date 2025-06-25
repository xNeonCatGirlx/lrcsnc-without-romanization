package lyrics

import (
	"lrcsnc/internal/pkg/log"
	"lrcsnc/internal/pkg/structs"
)

// Configure sets up the lyrics data by applying necessary configurations.
// Currently, it only applies romanization to the lyrics data.
// May be extended in the future.
//
// Every function/method/module/whatever needs to lock the mutex
// by themselves and only themselves.
// No locking a mutex in THIS function.
func Configure(lyricsData *structs.LyricsData) {
	log.Debug("lyrics/configure", "Starting configuring the received lyrics")
	log.Debug("lyrics/configure", "Done")
}
