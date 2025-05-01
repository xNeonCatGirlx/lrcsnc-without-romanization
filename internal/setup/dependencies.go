package setup

import (
	"os/exec"

	"lrcsnc/internal/pkg/global"
	"lrcsnc/internal/pkg/log"
)

// CheckDependencies checks if the required dependencies are installed
// and makes the appropriate adjustments to the config.
func CheckDependencies() {
	global.Config.M.Lock()
	defer global.Config.M.Unlock()

	// kakasi - for japanese romanization
	if _, err := exec.LookPath("kakasi"); err != nil {
		log.Info("setup/dependencies", "kakasi not found, disabling Japanese romanization")
		global.Config.C.Lyrics.Romanization.Japanese = false
	}
}
