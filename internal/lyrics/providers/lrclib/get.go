package lrclib

import (
	"net/url"

	"lrcsnc/internal/pkg/errors"
	"lrcsnc/internal/pkg/structs"
	"lrcsnc/internal/pkg/types"
)

func (l Provider) GetLyrics(song structs.Song) (structs.LyricsData, error) {
	var getURL *url.URL
	var body []byte
	var err error
	var res structs.LyricsData

	// Try to get the lyrics with everything exact: artists, album, duration
	if song.Duration != 0 {
		getURL = makeURL(song, lrcLibURLTypeGet)
		body, err = sendRequest(getURL)
	}
	if err == nil {
		res, err = dtoListToLyricsData(song, body)
	}
	if err != errors.ErrLyricsNotFound {
		return res, err
	}

	// Try to get the lyrics with everything exact except pick only the first artist
	getURL = makeURL(song, lrcLibURLTypeGetWithSingleArtist)
	body, err = sendRequest(getURL)
	if err == nil {
		res, err = dtoListToLyricsData(song, body)
	}
	if err != errors.ErrLyricsNotFound {
		return res, err
	}

	// Try to search for lyrics with exact album and artists
	getURL = makeURL(song, lrcLibURLTypeSearchWithAlbum)
	body, err = sendRequest(getURL)
	if err == nil {
		res, err = dtoListToLyricsData(song, body)
	}
	if err != errors.ErrLyricsNotFound {
		return res, err
	}

	// Try to search for lyrics with exact album, but only the first artist
	getURL = makeURL(song, lrcLibURLTypeSearchWithSingleArtistAndAlbum)
	body, err = sendRequest(getURL)
	if err == nil {
		res, err = dtoListToLyricsData(song, body)
	}
	if err != errors.ErrLyricsNotFound {
		return res, err
	}

	// Try to search for lyrics with only the title and all artists
	getURL = makeURL(song, lrcLibURLTypeSearch)
	body, err = sendRequest(getURL)
	if err == nil {
		res, err = dtoListToLyricsData(song, body)
	}
	if err != errors.ErrLyricsNotFound {
		return res, err
	}

	// Try to search for lyrics with only the title and the first artist
	getURL = makeURL(song, lrcLibURLTypeSearchWithSingleArtist)
	body, err = sendRequest(getURL)
	if err == nil {
		res, err = dtoListToLyricsData(song, body)
	}
	if err != errors.ErrLyricsNotFound {
		return res, err
	}

	// If nothing is found, return a not found state
	return structs.LyricsData{LyricsType: types.LyricsStateNotFound}, errors.ErrLyricsNotFound
}
