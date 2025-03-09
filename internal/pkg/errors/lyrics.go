package errors

import "errors"

// ErrLyricsNotFound is returned when the requested song's lyrics are not found
var ErrLyricsNotFound = errors.New("the requested song's lyrics were not found")

// ErrLyricsServerError is returned when the server returns any other error code than 404
var ErrLyricsServerError = errors.New("a server error occurred")

// ErrLyricsBodyReadFail is returned when the body of the response could not be read
var ErrLyricsBodyReadFail = errors.New("failed to read the response body")