package lyrics

import (
	"lrcsnc/internal/pkg/global"
	"lrcsnc/internal/pkg/structs"
	"lrcsnc/internal/romanization"
)

// Configure sets up the lyrics data by applying necessary configurations.
// Currently, it only applies romanization to the lyrics data.
//
// Parameters:
//   - lyricsData (*structs.LyricsData): A pointer to the LyricsData struct containing the lyrics to be configured.
func Configure(lyricsData *structs.LyricsData) {
	global.Config.M.Lock()
	defer global.Config.M.Unlock()

	// Romanization
	romanization.RomanizeLyrics(lyricsData.Lyrics)
}
