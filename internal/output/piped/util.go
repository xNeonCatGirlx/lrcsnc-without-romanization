package piped

import (
	"lrcsnc/internal/pkg/structs"
	"strings"
)

func lyricIndexToString(l int, lyricsData []structs.Lyric) string {
	if l < 0 || l >= len(lyricsData) {
		return ""
	} else {
		return lyricsData[l].Text
	}
}

func getInstrumentalMessage(c structs.MessageOutputConfig, outputFormat string) string {
	if !c.Enabled {
		return ""
	}

	replacer := strings.NewReplacer(
		"{icon}", c.Icon,
		"{lyric}", c.Text,
		"{multiplier}", "",
	)
	return strings.TrimSpace(replacer.Replace(outputFormat))
}

func getInstrumentalString(c structs.LyricOutputConfig, outputFormat string) string {
	replacer := strings.NewReplacer(
		"{icon}", c.Icon,
		"{lyric}", "",
		"{multiplier}", "",
	)
	return strings.TrimSpace(replacer.Replace(outputFormat))
}
