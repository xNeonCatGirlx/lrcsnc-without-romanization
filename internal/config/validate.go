package config

import (
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"strconv"

	"lrcsnc/internal/pkg/structs"
)

type ValidationError struct {
	Path    string
	Message string
	Fatal   bool
}

func (v ValidationError) Error() string {
	return v.Message
}

type ValidationErrors []ValidationError

func Validate(c *structs.Config) (errs ValidationErrors) {
	errs = make(ValidationErrors, 0)

	// Check whether output value is allowed
	if c.Output.Type != "piped" && c.Output.Type != "tui" {
		errs = append(errs, ValidationError{
			Path:    "global/output",
			Message: fmt.Sprintf("'%s' is not a valid value. Allowed values are 'piped' and 'tui'", c.Output.Type),
			Fatal:   true,
		})
	}

	// Check whether lrclib is set as the lyrics provider
	if c.Lyrics.Provider != "lrclib" {
		errs = append(errs, ValidationError{
			Path:    "global/lyrics-provider",
			Message: fmt.Sprintf("'%s' is not a valid value. Allowed values are 'lrclib' (sure hope there will be more in the future)", c.Lyrics.Provider),
			Fatal:   true,
		})
	}

	// Check if piped output's destination is writeable if it's not stdout
	if c.Output.Type == "piped" && c.Output.Piped.Destination != "stdout" && !isPathWriteable(c.Output.Piped.Destination) {
		errs = append(errs, ValidationError{
			Path:    "output/piped/destination",
			Message: fmt.Sprintf("'%s' is not a writeable path. Please make sure the path exists and is writeable", c.Output.Piped.Destination),
			Fatal:   true,
		})
	}

	// Check if the instrumental interval is set to <0.1s
	if c.Output.Type == "piped" && c.Output.Piped.Instrumental.Interval < 0.1 {
		errs = append(errs, ValidationError{
			Path:    "output/piped/instrumental/interval",
			Message: fmt.Sprintf("'%f' is not a valid value. Using the possible minimum instead (0.1s)", c.Output.Piped.Instrumental.Interval),
			Fatal:   false,
		})
		c.Output.Piped.Instrumental.Interval = 0.1
	}

	// Check if max symbols is less than 1
	if c.Output.Type == "piped" && c.Output.Piped.Instrumental.MaxSymbols < 1 {
		errs = append(errs, ValidationError{
			Path:    "output/piped/instrumental/max-symbols",
			Message: fmt.Sprintf("'%d' is not a valid value. Using the possible minimum instead (1)", c.Output.Piped.Instrumental.MaxSymbols),
			Fatal:   false,
		})
		c.Output.Piped.Instrumental.MaxSymbols = 1
	}

	if c.Output.Type == "tui" {
		// Check if the colors in the theme are not allowed
		if !isValidColor(c.Output.TUI.Theme.ProgressBarColor) {
			errs = append(errs, ValidationError{
				Path:    "output/tui/theme/progress-bar-color",
				Message: fmt.Sprintf("'%s' is not a valid color. Using ANSI white instead (15)", c.Output.TUI.Theme.ProgressBarColor),
				Fatal:   false,
			})
			c.Output.TUI.Theme.ProgressBarColor = "15"
		}
		if !isValidColor(c.Output.TUI.Theme.LyricBefore.Color) {
			errs = append(errs, ValidationError{
				Path:    "output/tui/theme/lyric-before/color",
				Message: fmt.Sprintf("'%s' is not a valid color. Using ANSI white instead (15)", c.Output.TUI.Theme.LyricBefore.Color),
				Fatal:   false,
			})
			c.Output.TUI.Theme.LyricBefore.Color = "15"
		}
		if !isValidColor(c.Output.TUI.Theme.LyricCurrent.Color) {
			errs = append(errs, ValidationError{
				Path:    "output/tui/theme/lyric-current/color",
				Message: fmt.Sprintf("'%s' is not a valid color. Using ANSI white instead (15)", c.Output.TUI.Theme.LyricCurrent.Color),
				Fatal:   false,
			})
			c.Output.TUI.Theme.LyricCurrent.Color = "15"
		}
		if !isValidColor(c.Output.TUI.Theme.LyricAfter.Color) {
			errs = append(errs, ValidationError{
				Path:    "output/tui/theme/lyric-after/color",
				Message: fmt.Sprintf("'%s' is not a valid color. Using ANSI white instead (15)", c.Output.TUI.Theme.LyricAfter.Color),
				Fatal:   false,
			})
			c.Output.TUI.Theme.LyricAfter.Color = "15"
		}
		if !isValidColor(c.Output.TUI.Theme.LyricCursor.Color) {
			errs = append(errs, ValidationError{
				Path:    "output/tui/theme/lyric-cursor/color",
				Message: fmt.Sprintf("'%s' is not a valid color. Using ANSI white instead (15)", c.Output.TUI.Theme.LyricCursor.Color),
				Fatal:   false,
			})
			c.Output.TUI.Theme.LyricCursor.Color = "15"
		}
		if !isValidColor(c.Output.TUI.Theme.BorderCursor.Color) {
			errs = append(errs, ValidationError{
				Path:    "output/tui/theme/border-cursor/color",
				Message: fmt.Sprintf("'%s' is not a valid color. Using ANSI white instead (15)", c.Output.TUI.Theme.BorderCursor.Color),
				Fatal:   false,
			})
			c.Output.TUI.Theme.BorderCursor.Color = "15"
		}
		if !isValidColor(c.Output.TUI.Theme.TimestampBefore.Color) {
			errs = append(errs, ValidationError{
				Path:    "output/tui/theme/timestamp-before/color",
				Message: fmt.Sprintf("'%s' is not a valid color. Using ANSI white instead (15)", c.Output.TUI.Theme.TimestampBefore.Color),
				Fatal:   false,
			})
			c.Output.TUI.Theme.TimestampBefore.Color = "15"
		}
		if !isValidColor(c.Output.TUI.Theme.TimestampCurrent.Color) {
			errs = append(errs, ValidationError{
				Path:    "output/tui/theme/timestamp-current/color",
				Message: fmt.Sprintf("'%s' is not a valid color. Using ANSI white instead (15)", c.Output.TUI.Theme.TimestampCurrent.Color),
				Fatal:   false,
			})
			c.Output.TUI.Theme.TimestampCurrent.Color = "15"
		}
		if !isValidColor(c.Output.TUI.Theme.TimestampAfter.Color) {
			errs = append(errs, ValidationError{
				Path:    "output/tui/theme/timestamp-after/color",
				Message: fmt.Sprintf("'%s' is not a valid color. Using ANSI white instead (15)", c.Output.TUI.Theme.TimestampAfter.Color),
				Fatal:   false,
			})
			c.Output.TUI.Theme.TimestampAfter.Color = "15"
		}
		if !isValidColor(c.Output.TUI.Theme.TimestampCursor.Color) {
			errs = append(errs, ValidationError{
				Path:    "output/tui/theme/timestamp-cursor/color",
				Message: fmt.Sprintf("'%s' is not a valid color. Using ANSI white instead (15)", c.Output.TUI.Theme.TimestampCursor.Color),
				Fatal:   false,
			})
			c.Output.TUI.Theme.TimestampCursor.Color = "15"
		}
	}

	return
}

func isPathWriteable(p string) bool {
	p = path.Clean(p)
	f, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE, 0666)
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
