package piped

import "lrcsnc/internal/pkg/global"

var currentDestination string = "stdout"

type Controller struct{}

func (Controller) OnConfigUpdate() {
	global.Config.M.Lock()
	defer global.Config.M.Unlock()

	if global.Config.C.Output.Piped.Destination != currentDestination &&
		changeOutput(global.Config.C.Output.Piped.Destination) == nil {
		currentDestination = global.Config.C.Output.Piped.Destination
	}
}

func (Controller) OnPlayerUpdate() {}

func (Controller) OnOverwrite(overwrite string) {
	Overwrite(overwrite)
}

func (Controller) DisplayLyric(lyricIndex int) {
	currentLyricChangedChan <- lyricIndex
}
