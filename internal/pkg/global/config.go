package global

import (
	"sync"

	"lrcsnc/internal/pkg/structs"
)

var Config = struct {
	M sync.Mutex
	C structs.Config

	Path string
}{}
