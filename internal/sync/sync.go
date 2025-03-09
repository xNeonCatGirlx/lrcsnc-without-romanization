package sync

func Loop() {
	// Goroutine of the position synchronizer
	go positionSynchronizer()

	// Goroutine to watch for DBus signals
	go signalWatcher()

	// Goroutine to check for changes in currently playing song
	go lyricFetcher()

	// Goroutine to actively synchronize the lyrics with the song
	go lyricsSynchronizer()
}
