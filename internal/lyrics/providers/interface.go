package providers

import (
	lrclib "lrcsnc/internal/lyrics/providers/lrclib"

	"lrcsnc/internal/pkg/structs"
	"lrcsnc/internal/pkg/types"
)

type Provider interface {
	// Get returns the lyrics of a song in form of LyricsData
	Get(structs.Song) (structs.LyricsData, error)
}

var Providers = map[types.LyricsProviderType]Provider{
	types.LyricsProviderLrclib: lrclib.Provider{},
}
