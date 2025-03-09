package lyrics

import (
	"fmt"
	"lrcsnc/internal/cache"
	"lrcsnc/internal/log"
	"lrcsnc/internal/lyrics/providers"
	"lrcsnc/internal/pkg/global"
	"lrcsnc/internal/pkg/structs"
	"lrcsnc/internal/pkg/types"
)

// GetLyricsData retrieves the lyrics data for a given song.
// It first checks if caching is enabled and attempts to retrieve the lyrics from the cache.
// If the lyrics are not found in the cache, it fetches the lyrics from the configured lyrics provider.
// If the lyrics are successfully retrieved and caching is enabled, it stores the lyrics in the cache.
// If the lyrics are not found or an error occurs during retrieval, appropriate log messages are generated.
//
// Parameters:
//   - song: The song for which to retrieve the lyrics. It has to be sent via copy
//     because the original song's struct may become modified quite fast.
//
// Returns:
//   - structs.LyricsData: The retrieved lyrics data.
//   - error: An error if the lyrics could not be retrieved or another issue occurred.
func GetLyricsData(song structs.Song) (structs.LyricsData, error) {
	if global.Config.C.Cache.Enabled {
		cachedData, cacheState := cache.Get(&song)
		if cacheState == cache.CacheStateActive {
			return cachedData, nil
		}
	}

	res, err := providers.LyricsDataProviders[global.Config.C.Global.LyricsProvider].GetLyrics(song)
	if err != nil {
		if err.Error() == "the requested song was not found" {
			log.Info("lyrics/get", fmt.Sprintf("The lyrics for %s - %s were not found", song.Artists[0], song.Title))
		} else {
			log.Error("lyrics/get", fmt.Sprintf("Could not get the lyrics: %s", err))
		}

		return res, err
	}

	if global.Config.C.Cache.Enabled && res.LyricsType == types.LyricsStateSynced {
		song.LyricsData = res
		if cache.Store(&song) != nil {
			// TODO: logger :)
			// log.Println("Could not save the lyrics to the cache! Is there an issue with perms?")
		}
	}

	return res, nil
}
