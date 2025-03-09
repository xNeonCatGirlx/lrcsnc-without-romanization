package tui

var (
	ConfigChanged       = make(chan bool)
	SongInfoChanged     = make(chan bool)
	PlayerInfoChanged   = make(chan bool)
	CurrentLyricChanged = make(chan int)
	OverwriteReceived   = make(chan string)
)
