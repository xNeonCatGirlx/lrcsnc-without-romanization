package config

import "errors"

// ErrFileUnreachable represents an error indicating that a file could not be reached or accessed.
var ErrFileUnreachable = errors.New("config file is unreachable")

// ErrFileInvalid represents that a TOML parsing error occurred
var ErrFileInvalid = errors.New("config file is invalid")
