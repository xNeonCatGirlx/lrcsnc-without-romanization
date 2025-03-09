package tui

import (
	"lrcsnc/internal/pkg/global"
	"lrcsnc/internal/pkg/structs"
)

var player structs.Player
var config structs.TUIOutputConfig

type Controller struct{}

func (Controller) OnConfigChange() {
	global.Config.M.Lock()
	defer global.Config.M.Unlock()

	config = global.Config.C.Output.TUI
}

func (Controller) OnPlayerChange() {
	global.Player.M.Lock()
	defer global.Player.M.Unlock()

	prevSongID := player.Song.ID()

	player = global.Player.P

	PlayerInfoChanged <- true
	if prevSongID != global.Player.P.Song.ID() {
		SongInfoChanged <- true
	}
}

func (Controller) OnOverwrite(overwrite string) {
	OverwriteReceived <- overwrite
}

func (Controller) DisplayLyric(lyricIndex int) {
	CurrentLyricChanged <- lyricIndex
}
