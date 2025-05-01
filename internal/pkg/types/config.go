package types

// LyricsProviderType sets which lyrics provider to use.
//
// Possible values: "lrclib".
type LyricsProviderType string

const (
	LyricsProviderLrclib LyricsProviderType = "lrclib"
)

// CacheStoreConditionType is a bit flag that sets the condition for when to save cache
//
// Possible values:
// - first 1 for when the lyrics are synced,
// - second 1 for when the lyrics are plain,
// - third 1 for when the song is instrumental,
type CacheStoreConditionType uint8

const (
	CacheStoreConditionSynced       CacheStoreConditionType = 0b100
	CacheStoreConditionPlain        CacheStoreConditionType = 0b010
	CacheStoreConditionInstrumental CacheStoreConditionType = 0b001
	CacheStoreConditionNone         CacheStoreConditionType = 0b000
)

// OutputType is a type of output to use.
//
// Possible values: "piped".
type OutputType string

const (
	OutputPiped OutputType = "piped"
)

// LogLevelType represents the log level to use in logger.
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
		return -1
	}
}
