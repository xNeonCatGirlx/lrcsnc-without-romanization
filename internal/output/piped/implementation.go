package piped

import (
	"lrcsnc/internal/pkg/global"
	"lrcsnc/internal/pkg/structs"
	"strconv"
	"strings"
)

var player structs.Player
var config structs.PipedOutputConfig

type Controller struct{}

func (Controller) OnConfigChange() {
	global.Config.M.Lock()
	defer global.Config.M.Unlock()

	config = global.Config.C.Output.Piped
}

func (Controller) OnPlayerChange() {
	global.Player.M.Lock()
	defer global.Player.M.Unlock()

	player = global.Player.P
}

func (Controller) OnOverwrite(overwrite string) {
	Overwrite(overwrite)
}

func (Controller) DisplayLyric(lyricIndex int) {
	lyric := lyricIndexToString(lyricIndex)
	global.Config.M.Lock()
	defer global.Config.M.Unlock()
	multiplier := 0
	for i := lyricIndex; i >= 0 && player.Song.LyricsData.Lyrics[i].Text != lyric; i-- {
		multiplier++
	}
	replacer := strings.NewReplacer(
		"{icon}", global.Config.C.Output.Piped.Lyric.Icon,
		"{lyric}", lyric,
		"{multiplier}", strings.ReplaceAll(global.Config.C.Output.Piped.MultiplierFormat, "{value}", strconv.Itoa(multiplier)),
	)
	Print(strings.TrimSpace(replacer.Replace(global.Config.C.Output.Piped.OutputFormat)))
}
