package structs

import (
	"lrcsnc/internal/pkg/types"
)

// LEVEL 0

type Config struct {
	// Player config is for player related things. Currently it is used
	// for specifying included/excluded players for the watcher.
	Player PlayerConfig `toml:"player"`
	// Lyrics config currently has stuff to do with lyrics providers,
	// time offset and romanization
	Lyrics LyricsConfig `toml:"lyrics"`
	// Cache config has an "enabled" toggle, dir path and life span
	Cache CacheConfig `toml:"cache"`
	// Output config has... a lot of personalized settings.
	Output OutputConfig `toml:"output"`
}

// LEVEL 1

type PlayerConfig struct {
	IncludedPlayers []string `toml:"included-players"`
	ExcludedPlayers []string `toml:"excluded-players"`
}

type LyricsConfig struct {
	Provider        types.LyricsProviderType `toml:"provider"`
	TimestampOffset float64                  `toml:"timestamp-offset"`
	Romanization    RomanizationConfig       `toml:"romanization"`
}

type CacheConfig struct {
	Enabled        bool                          `toml:"enabled"`
	Dir            string                        `toml:"dir"`
	LifeSpan       uint                          `toml:"life-span"`
	StoreCondition types.CacheStoreConditionType `toml:"store-condition"`
}

type OutputConfig struct {
	Type  types.OutputType  `toml:"type"`
	Piped PipedOutputConfig `toml:"piped"`
	TUI   TUIOutputConfig   `toml:"tui"`
}

// LEVEL 2

type RomanizationConfig struct {
	Japanese bool `toml:"japanese"`
	Chinese  bool `toml:"chinese"`
	Korean   bool `toml:"korean"`
}

func (r *RomanizationConfig) IsEnabled() bool {
	return r.Japanese || r.Chinese || r.Korean
}

type PipedOutputConfig struct {
	Destination    string              `toml:"destination"`
	JSON           bool                `toml:"json"`
	InsertNewline  bool                `toml:"insert-newline"`
	Text           FormatOutputConfig  `toml:"text"`
	Multiplier     FormatOutputConfig  `toml:"multiplier"`
	Lyric          LyricOutputConfig   `toml:"lyric"`
	SongNotFound   MessageOutputConfig `toml:"song-not-found"`
	NoSyncedLyrics MessageOutputConfig `toml:"no-synced-lyrics"`
	GettingLyrics  MessageOutputConfig `toml:"getting-lyrics"`
	ErrorMessage   MessageOutputConfig `toml:"error-message"`
	Instrumental   InstrumentalConfig  `toml:"instrumental"`
}

// TODO: Move ShowTimestamps and ShowProgressBar to a state file
type TUIOutputConfig struct {
	ShowTimestamps  bool           `toml:"show-timestamps"`
	ShowProgressBar bool           `toml:"show-progress-bar"`
	Theme           TUIThemeConfig `toml:"theme"`
}

// LEVEL 3

type FormatOutputConfig struct {
	Format string `toml:"format"`
}

type LyricOutputConfig struct {
	Icon string `toml:"icon"`
}

type MessageOutputConfig struct {
	Enabled bool   `toml:"enabled"`
	Icon    string `toml:"icon"`
	Text    string `toml:"text"`
}

type InstrumentalConfig struct {
	Interval   float64 `toml:"interval"`
	Symbol     string  `toml:"symbol"`
	MaxSymbols uint    `toml:"max-symbols"`
}

type TUIThemeConfig struct {
	LyricBefore      TUIThemeTextConfig   `toml:"lyric-before"`
	LyricCurrent     TUIThemeTextConfig   `toml:"lyric-current"`
	LyricAfter       TUIThemeTextConfig   `toml:"lyric-after"`
	LyricCursor      TUIThemeTextConfig   `toml:"lyric-cursor"`
	BorderCursor     TUIThemeBorderConfig `toml:"border-cursor"`
	TimestampBefore  TUIThemeTextConfig   `toml:"timestamp-before"`
	TimestampCurrent TUIThemeTextConfig   `toml:"timestamp-current"`
	TimestampAfter   TUIThemeTextConfig   `toml:"timestamp-after"`
	TimestampCursor  TUIThemeTextConfig   `toml:"timestamp-cursor"`
	ProgressBarColor string               `toml:"progress-bar-color"`
}

// LEVEL 4

type TUIThemeTextConfig struct {
	Color string `toml:"color"`
	Bold  bool   `toml:"bold"`
	Faint bool   `toml:"faint"`
}

type TUIThemeBorderConfig struct {
	Color  string `toml:"color"`
	Top    bool   `toml:"top"`
	Right  bool   `toml:"right"`
	Bottom bool   `toml:"bottom"`
	Left   bool   `toml:"left"`
}
