package config

// A type of output to use.
//
// Possible values: "piped", "tui".
type OutputType string

const (
	OutputPiped OutputType = "piped"
	OutputTUI   OutputType = "tui"
)

// Sets which lyrics provider to use.
//
// Possible values: "lrclib".
type LyricsProviderType string

const (
	LyricsProviderLrclib LyricsProviderType = "lrclib"
)

// Represents the log level to use in logger.
//
// Possible values: "debug", "info", "warn", "error", "fatal".
type LogLevelType string

const (
	LogLevelDebug LogLevelType = "debug"
	LogLevelInfo  LogLevelType = "info"
	LogLevelWarn  LogLevelType = "warn"
	LogLevelError LogLevelType = "error"
	LogLevelFatal LogLevelType = "fatal"
)

func (l LogLevelType) ToInt() int {
	switch l {
	case LogLevelDebug:
		return 4
	case LogLevelInfo:
		return 3
	case LogLevelWarn:
		return 2
	case LogLevelError:
		return 1
	case LogLevelFatal:
		return 0
	default:
		return 3
	}
}
