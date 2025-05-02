package piped

import (
	"encoding/json"
	"fmt"
	"lrcsnc/internal/pkg/global"
	"lrcsnc/internal/pkg/log"
	"strings"
)

type JSONOutput struct {
	Text     string `json:"text"`
	Title    string `json:"song"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Player   string `json:"player"`
	Position string `json:"position"`
	Duration string `json:"duration"`
}

func FormatToJSON(text string) string {
	global.Player.M.Lock()
	defer global.Player.M.Unlock()

	jsonOutput := JSONOutput{
		Text:     text,
		Title:    global.Player.P.Song.Title,
		Artist:   strings.Join(global.Player.P.Song.Artists, ", "),
		Album:    global.Player.P.Song.Album,
		Player:   global.Player.P.Name,
		Position: fmt.Sprintf("%02d:%02d", int(global.Player.P.Position)/60, int(global.Player.P.Position)%60),
		Duration: fmt.Sprintf("%02d:%02d", int(global.Player.P.Song.Duration)/60, int(global.Player.P.Song.Duration)%60),
	}

	jsonData, err := json.Marshal(jsonOutput)
	if err != nil {
		log.Error("output/piped", "Error marshalling JSON output: "+err.Error())
		return "{}"
	}

	return string(jsonData)
}
