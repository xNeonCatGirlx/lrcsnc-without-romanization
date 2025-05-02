package json

import (
	"encoding/json"
	"fmt"
	"lrcsnc/internal/pkg/global"
	"lrcsnc/internal/pkg/log"
	"lrcsnc/internal/pkg/types"
	"strings"
)

func FormatToJSON(text string) string {
	var jsonOutput any

	switch global.Config.C.Output.Piped.JSON {
	case types.JSONOutputGeneric:
		jsonOutput = JSONOutput{
			Text:     text,
			Title:    global.Player.P.Song.Title,
			Artist:   strings.Join(global.Player.P.Song.Artists, ", "),
			Album:    global.Player.P.Song.Album,
			Player:   global.Player.P.Name,
			Position: fmt.Sprintf("%02d:%02d", int(global.Player.P.Position)/60, int(global.Player.P.Position)%60),
			Duration: fmt.Sprintf("%02d:%02d", int(global.Player.P.Song.Duration)/60, int(global.Player.P.Song.Duration)%60),
		}
	case types.JSONOutputWaybar:
		altTooltipReplacer := strings.NewReplacer(
			"{text}", text,
			"{artist}", global.Player.P.Song.Artists[0],
			"{artists}", strings.Join(global.Player.P.Song.Artists, ", "),
			"{title}", global.Player.P.Song.Title,
			"{album}", global.Player.P.Song.Album,
			"{position}", fmt.Sprintf("%02d:%02d", int(global.Player.P.Position)/60, int(global.Player.P.Position)%60),
			"{duration}", fmt.Sprintf("%02d:%02d", int(global.Player.P.Song.Duration)/60, int(global.Player.P.Song.Duration)%60),
		)

		classReplacer := strings.NewReplacer(
			"{playback-status}", strings.ToLower(string(global.Player.P.PlaybackStatus)),
			"{lyrics-status}", strings.ToLower(global.Player.P.Song.LyricsData.LyricsState.String()),
		)

		jsonOutput = WaybarJSONOutput{
			Text:       text,
			Alt:        strings.TrimSpace(altTooltipReplacer.Replace(global.Config.C.Output.Piped.JSONWaybar.Alt)),
			Tooltip:    strings.TrimSpace(altTooltipReplacer.Replace(global.Config.C.Output.Piped.JSONWaybar.Tooltip)),
			Class:      strings.Split(strings.TrimSpace(classReplacer.Replace(global.Config.C.Output.Piped.JSONWaybar.Class)), " "),
			Percentage: 0,
		}
	}

	jsonData, err := json.Marshal(jsonOutput)
	if err != nil {
		log.Error("output/piped", "Error marshalling JSON output: "+err.Error())
		return "{}"
	}

	return string(jsonData)
}
