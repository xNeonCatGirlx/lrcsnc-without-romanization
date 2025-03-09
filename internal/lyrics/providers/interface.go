package providers

import (
	lrclib "lrcsnc/internal/lyrics/providers/lrclib"

	"lrcsnc/internal/pkg/structs"
	"lrcsnc/internal/pkg/types"
)

type Provider interface {
	// GetLyrics returns the lyrics of a song in form of LyricsData
	GetLyrics(structs.Song) (structs.LyricsData, error)
}

var LyricsDataProviders = map[types.LyricsProviderType]Provider{
	types.LyricsProviderLrclib: lrclib.Provider{},
}
