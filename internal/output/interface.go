package output

import (
	"lrcsnc/internal/output/piped"
	"lrcsnc/internal/pkg/types"
)

type Controller interface {
	OnConfigUpdate()
	OnPlayerUpdate()
	OnOverwrite(overwrite string)

	DisplayLyric(lyricIndex int)
}

var Controllers = map[types.OutputType]Controller{
	types.OutputPiped: piped.Controller{},
}
