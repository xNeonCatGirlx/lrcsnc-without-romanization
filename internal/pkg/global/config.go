package global

import (
	"lrcsnc/internal/pkg/structs"
	"sync"
)

var CurrentConfig struct {
	Mutex  sync.Mutex
	Config structs.Config
}
