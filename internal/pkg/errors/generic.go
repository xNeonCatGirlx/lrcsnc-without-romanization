package errors

import "errors"

// ErrMarshalFail is returned when the request body could not be marshalled
var ErrMarshalFail = errors.New("encountered marshal error")

// ErrUnmarshalFail is returned when the response body could not be unmarshalled
var ErrUnmarshalFail = errors.New("encountered unmarshal error")

// ErrFileUnreachable represents an error indicating that a file could not be reached or accessed.
var ErrFileUnreachable = errors.New("file is unreachable")

// ErrFileUnreadable represents an error indicating that a file could not be read.
var ErrFileUnreadable = errors.New("file is unreadable")

// ErrFileUnwriteable represents an error indicating that a file could not be created or written in.
var ErrFileUnwriteable = errors.New("file is unwriteable")

// ErrDirUnreachable represents an error indicating that a directory could not be reached or accessed.
var ErrDirUnreachable = errors.New("directory is unreachable")

// ErrDirUnreadable represents an error indicating that a directory could not be read.
var ErrDirUnreadable = errors.New("directory is unreadable")

// ErrDirUnwriteable represents an error indicating that a directory could not be created or written in.
var ErrDirUnwriteable = errors.New("directory is unwriteable")
