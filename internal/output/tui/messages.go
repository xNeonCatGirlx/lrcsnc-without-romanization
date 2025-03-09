package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type (
	configChanged       bool
	songInfoChanged     bool
	playerInfoChanged   bool
	currentLyricChanged int
	overwriteReceived   string
)

func watchConfigChanges() tea.Cmd {
	return func() tea.Msg {
		return <-ConfigChanged
	}
}

func watchSongInfoChanges() tea.Cmd {
	return func() tea.Msg {
		return <-SongInfoChanged
	}
}

func watchPlayerInfoChanges() tea.Cmd {
	return func() tea.Msg {
		return <-PlayerInfoChanged
	}
}

func watchCurrentLyricChanges() tea.Cmd {
	return func() tea.Msg {
		return currentLyricChanged(<-CurrentLyricChanged)
	}
}

func watchReceivedOverwrites() tea.Cmd {
	return func() tea.Msg {
		return overwriteReceived(<-OverwriteReceived)
	}
}
