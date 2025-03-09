package output

import (
	"lrcsnc/internal/output/piped"
	"lrcsnc/internal/output/tui"
	"lrcsnc/internal/pkg/types"
)

type Controller interface {
	OnConfigChange()
	OnPlayerChange()
	OnOverwrite(overwrite string)

	DisplayLyric(lyricIndex int)
}

var Controllers = map[types.OutputType]Controller{
	types.OutputPiped: piped.Controller{},
	types.OutputTUI:   tui.Controller{},
}
