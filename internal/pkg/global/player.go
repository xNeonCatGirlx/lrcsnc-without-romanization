package global

import (
	"lrcsnc/internal/pkg/structs"
	"sync"
)

var CurrentPlayer = struct {
	Mutex  sync.Mutex
	Player structs.Player
}{
	Player: structs.Player{
		IsPlaying: false,
		Position:  0,
	},
}
