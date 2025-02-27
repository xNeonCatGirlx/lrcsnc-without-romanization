package config

import (
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"strconv"
)

func ValidateConfig(c *Config) (errs []error, fatal bool) {
	// Check if output value is not allowed
	if c.Global.Output != "piped" && c.Global.Output != "tui" {
		fatal = true
		errs = append(errs, fmt.Errorf(`[config: global/output] ERROR: Output should either be "piped" or "tui"`))
	}

	// Check if lrclib is not set as the lyrics provider
	if c.Global.LyricsProvider != "lrclib" {
		c.Global.LyricsProvider = "lrclib"
		errs = append(errs, fmt.Errorf(`[config: global/lyrics-provider] WARNING: For now, 'lrclib' is the only lyrics provider, so lrcsnc will always use 'lrclib' until there are new lyrics providers introduced`))
	}

	// Check if piped output's destination is writeable if it's not stdout
	if c.Global.Output == "piped" && c.Output.Piped.Destination != "stdout" && !isPathWriteable(c.Output.Piped.Destination) {
		fatal = true
		errs = append(errs, fmt.Errorf("[config: output/piped/destination] ERROR: The destination is not writeable. An issue with the path itself, or maybe permissions?"))
	}

	// Check if the instrumental interval is set to <0.1s
	if c.Global.Output == "piped" && c.Output.Piped.Instrumental.Interval < 0.1 {
		c.Output.Piped.Instrumental.Interval = 0.1
		errs = append(errs, fmt.Errorf("[config: output/piped/instrumental/interval] WARNING: Instrumental interval is set to a value less than 0.1. Using the possible minimum instead (0.1)"))
	}

	// Check if max symbols is less than 1
	if c.Global.Output == "piped" && c.Output.Piped.Instrumental.MaxSymbols < 1 {
		c.Output.Piped.Instrumental.MaxSymbols = 1
		errs = append(errs, fmt.Errorf("[config: output/piped/instrumental/max-symbols] WARNING: Max symbols in instrumental section is set to a value less than 1. Using the possible minimum instead (1)"))
	}

	if c.Global.Output == "tui" {
		// Check if the colors in the theme are not allowed
		if !isValidColor(c.Output.TUI.Theme.ProgressBarColor) {
			c.Output.TUI.Theme.ProgressBarColor = "15"
			errs = append(errs, fmt.Errorf("[config: output/tui/theme/progress-bar-color] WARNING: Color is not valid. Falling back to ANSI white"))
		}
		if !isValidColor(c.Output.TUI.Theme.LyricBefore.Color) {
			c.Output.TUI.Theme.LyricBefore.Color = "15"
			errs = append(errs, fmt.Errorf("[config: output/tui/theme/lyric-before/color] WARNING: Color is not valid. Falling back to ANSI white"))
		}
		if !isValidColor(c.Output.TUI.Theme.LyricCurrent.Color) {
			c.Output.TUI.Theme.LyricCurrent.Color = "15"
			errs = append(errs, fmt.Errorf("[config: output/tui/theme/lyric-current/color] WARNING: Color is not valid. Falling back to ANSI white"))
		}
		if !isValidColor(c.Output.TUI.Theme.LyricAfter.Color) {
			c.Output.TUI.Theme.LyricAfter.Color = "15"
			errs = append(errs, fmt.Errorf("[config: output/tui/theme/lyric-after/color] WARNING: Color is not valid. Falling back to ANSI white"))
		}
		if !isValidColor(c.Output.TUI.Theme.LyricCursor.Color) {
			c.Output.TUI.Theme.LyricCursor.Color = "15"
			errs = append(errs, fmt.Errorf("[config: output/tui/theme/lyric-cursor/color] WARNING: Color is not valid. Falling back to ANSI white"))
		}
		if !isValidColor(c.Output.TUI.Theme.BorderCursor.Color) {
			c.Output.TUI.Theme.BorderCursor.Color = "15"
			errs = append(errs, fmt.Errorf("[config: output/tui/theme/border-cursor/color] WARNING: Color is not valid. Falling back to ANSI white"))
		}
		if !isValidColor(c.Output.TUI.Theme.TimestampBefore.Color) {
			c.Output.TUI.Theme.TimestampBefore.Color = "15"
			errs = append(errs, fmt.Errorf("[config: output/tui/theme/timestamp-before/color] WARNING: Color is not valid. Falling back to ANSI white"))
		}
		if !isValidColor(c.Output.TUI.Theme.TimestampCurrent.Color) {
			c.Output.TUI.Theme.TimestampCurrent.Color = "15"
			errs = append(errs, fmt.Errorf("[config: output/tui/theme/timestamp-current/color] WARNING: Color is not valid. Falling back to ANSI white"))
		}
		if !isValidColor(c.Output.TUI.Theme.TimestampAfter.Color) {
			c.Output.TUI.Theme.TimestampAfter.Color = "15"
			errs = append(errs, fmt.Errorf("[config: output/tui/theme/timestamp-after/color] WARNING: Color is not valid. Falling back to ANSI white"))
		}
		if !isValidColor(c.Output.TUI.Theme.TimestampCursor.Color) {
			c.Output.TUI.Theme.TimestampCursor.Color = "15"
			errs = append(errs, fmt.Errorf("[config: output/tui/theme/timestamp-cursor/color] WARNING: Color is not valid. Falling back to ANSI white"))
		}
	}

	return
}

func isPathWriteable(p string) bool {
	p = path.Clean(p)
	f, err := os.OpenFile(p, 0777, os.ModeExclusive)
	if err != nil {
		return false
	} else {
		f.Close()
		return true
	}
}

func isValidColor(c string) bool {
	if _, err := strconv.ParseUint(c, 10, 8); err == nil {
		return true
	} else if len(c) == 7 && c[0] == '#' {
		if _, err := hex.DecodeString(c[1:]); err == nil {
			return true
		}
	}

	return false
}
