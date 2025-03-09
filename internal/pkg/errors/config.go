package errors

import "errors"

// ErrConfigFileInvalid represents that a TOML parsing error occurred
var ErrConfigFileInvalid = errors.New("config file is invalid")

// ErrConfigFatalValidation represents that a fatal validation error occurred
var ErrConfigFatalValidation = errors.New("fatal validation errors")
