package cache

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"lrcsnc/internal/log"
	"lrcsnc/internal/pkg/errors"
	"lrcsnc/internal/pkg/global"
	"lrcsnc/internal/pkg/structs"
	"lrcsnc/internal/pkg/util"
)

// Get retrieves the cached lyrics data for a given song.
// It first checks if the cache is enabled. If not, it returns an empty LyricsData and CacheStateDisabled.
// If the cache is enabled, it constructs the cache file path using the song ID.
// It attempts to read the cache file and unmarshal its contents into a LyricsData struct.
// If successful, it checks if the cache has expired based on the configured cache lifespan.
// It returns the cached data along with the appropriate CacheState (Active, Expired, or NonExistant).
//
// Parameters:
//   - song: A pointer to the Song struct for which the cached lyrics data is to be retrieved.
//
// Returns:
//   - LyricsData: The cached lyrics data.
//   - CacheState: The state of the cache (Disabled, Active, Expired, or NonExistant).
func Get(song *structs.Song) (structs.LyricsData, CacheState) {
	global.Config.M.Lock()
	defer global.Config.M.Unlock()

	if !global.Config.C.Cache.Enabled {
		return structs.LyricsData{}, CacheStateDisabled
	}
	cacheDirectory := getCacheDir()

	filename := getFilename(song)
	fullPath := cacheDirectory + "/" + filename + ".json"

	log.Debug("cache/Get", fmt.Sprintf("Trying to get the cached data for %v", filename))

	if file, err := os.ReadFile(fullPath); err == nil {
		var cachedData structs.LyricsData
		err = json.Unmarshal(file, &cachedData)
		if err != nil {
			log.Debug("cache/Get", fmt.Sprintf("Couldn't unmarshal the cached data for %v", filename))
			return structs.LyricsData{}, CacheStateNonExistant
		}

		log.Debug("cache/Get", fmt.Sprintf("Successfully got the cached data for %v", filename))

		if global.Config.C.Cache.CacheLifeSpan != 0 {
			cacheStats, _ := os.Lstat(fullPath)
			isExpired := time.Since(cacheStats.ModTime()).Hours() <= float64(global.Config.C.Cache.CacheLifeSpan)*24
			if isExpired {
				return cachedData, CacheStateExpired
			} else {
				return cachedData, CacheStateActive
			}
		} else {
			return cachedData, CacheStateActive
		}
	} else {
		return structs.LyricsData{}, CacheStateNonExistant
	}
}

// Store saves the lyrics data of a given song to a JSON file in the cache directory.
//
// Parameters:
//   - song: A pointer to a Song struct containing the lyrics data to be stored.
//
// Returns:
//   - error: An error if any occurs during the process of creating the cache directory,
//     marshalling the song's lyrics data, or writing the data to a file.
func Store(song *structs.Song) error {
	global.Config.M.Lock()
	defer global.Config.M.Unlock()

	cacheDirectory := getCacheDir()
	if _, err := os.ReadDir(cacheDirectory); err != nil {
		err = os.Mkdir(cacheDirectory, 0666)
		if err != nil {
			return errors.ErrDirUnwriteable
		}
	}

	filename := getFilename(song)
	fullPath := cacheDirectory + "/" + filename + ".json"

	encodedData, err := json.Marshal(song.LyricsData)
	if err != nil {
		return errors.ErrMarshalFail
	}

	if err := os.WriteFile(fullPath, []byte(encodedData), 0666); err != nil {
		return errors.ErrFileUnwriteable
	}
	return nil
}

// Remove deletes the cached data for the given song from the cache directory.
// If the cache directory is not reachable, it returns an ErrDirUnreachable error.
// If the specific cached file for the song cannot be removed,
// it logs an error message and returns an ErrFileUnreachable error.
//
// Parameters:
//   - song: A pointer to the Song struct representing the song whose cached data
//     needs to be removed.
//
// Returns:
//   - error: An error if the cache directory is not reachable or if the specific
//     cached file cannot be removed, otherwise nil.
func Remove(song *structs.Song) error {
	global.Config.M.Lock()
	defer global.Config.M.Unlock()

	cacheDirectory := getCacheDir()
	if _, err := os.ReadDir(cacheDirectory); err != nil {
		return errors.ErrDirUnreachable
	}

	filename := getFilename(song)
	fullPath := cacheDirectory + "/" + filename + ".json"

	if err := os.Remove(fullPath); err != nil {
		log.Error("cache/Remove", fmt.Sprintf("Couldn't delete the cached data for %v. Maybe the data didn't exist in the first place?", filename))
		return errors.ErrFileUnreachable
	}
	return nil
}

func getCacheDir() string {
	return os.ExpandEnv(global.Config.C.Cache.CacheDir)
}

func getFilename(song *structs.Song) string {
	return fmt.Sprintf("%v.%v.%v.%v",
		util.RemoveBadCharacters(song.Title),
		util.RemoveBadCharacters(strings.Join(song.Artists, ", ")),
		util.RemoveBadCharacters(song.Album),
		math.Round(song.Duration),
	)
}
