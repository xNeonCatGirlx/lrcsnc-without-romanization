package piped

import (
	"lrcsnc/internal/pkg/structs"
	"strings"
)

func lyricIndexToString(l int) string {
	if l < 0 || l >= len(player.Song.LyricsData.Lyrics) {
		return ""
	} else {
		return player.Song.LyricsData.Lyrics[l].Text
	}
}

func getOutString(c structs.MessageOutputConfig) string {
	if !c.Enabled {
		return ""
	}

	replacer := strings.NewReplacer(
		"{icon}", c.Icon,
		"{lyric}", c.Text,
		"{multiplier}", "",
	)
	return strings.TrimSpace(replacer.Replace(config.OutputFormat))
}
