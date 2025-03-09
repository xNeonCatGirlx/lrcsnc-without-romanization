package tui

import (
	gloss "github.com/charmbracelet/lipgloss"
)

var (
	styleLyricGeneric = gloss.NewStyle().AlignHorizontal(gloss.Center)
	styleLyricBefore  gloss.Style
	styleLyricCurrent gloss.Style
	styleLyricAfter   gloss.Style
	styleLyricCursor  gloss.Style

	styleBorderCursor gloss.Style

	styleTimestampBefore  gloss.Style
	styleTimestampCurrent gloss.Style
	styleTimestampAfter   gloss.Style
	styleTimestampCursor  gloss.Style
)

func reloadStyles() {
	styleLyricBefore = styleLyricGeneric.Foreground(
		gloss.Color(
			config.Theme.LyricBefore.Color,
		),
	).Faint(
		config.Theme.LyricBefore.Faint,
	).Bold(
		config.Theme.LyricBefore.Bold,
	)

	styleLyricCurrent = styleLyricGeneric.Foreground(
		gloss.Color(
			config.Theme.LyricCurrent.Color,
		),
	).Faint(
		config.Theme.LyricCurrent.Faint,
	).Bold(
		config.Theme.LyricCurrent.Bold,
	)

	styleLyricAfter = styleLyricGeneric.Foreground(
		gloss.Color(
			config.Theme.LyricAfter.Color,
		),
	).Faint(
		config.Theme.LyricAfter.Faint,
	).Bold(
		config.Theme.LyricAfter.Bold,
	)

	styleLyricCursor = styleLyricGeneric.Foreground(
		gloss.Color(
			config.Theme.LyricCursor.Color,
		),
	).Faint(
		config.Theme.LyricCursor.Faint,
	).Bold(
		config.Theme.LyricCursor.Bold,
	)

	styleBorderCursor = gloss.NewStyle().Border(
		gloss.ThickBorder(),
		config.Theme.BorderCursor.Top,
		config.Theme.BorderCursor.Right,
		config.Theme.BorderCursor.Bottom,
		config.Theme.BorderCursor.Left,
	).BorderForeground(
		gloss.Color(
			config.Theme.BorderCursor.Color,
		),
	)

	styleTimestampBefore = gloss.NewStyle().Foreground(
		gloss.Color(
			config.Theme.TimestampBefore.Color,
		),
	).Faint(
		config.Theme.TimestampBefore.Faint,
	).Bold(
		config.Theme.TimestampBefore.Bold,
	)

	styleTimestampCurrent = gloss.NewStyle().Foreground(
		gloss.Color(
			config.Theme.TimestampCurrent.Color,
		),
	).Faint(
		config.Theme.TimestampCurrent.Faint,
	).Bold(
		config.Theme.TimestampCurrent.Bold,
	)

	styleTimestampAfter = gloss.NewStyle().Foreground(
		gloss.Color(
			config.Theme.TimestampAfter.Color,
		),
	).Faint(
		config.Theme.TimestampAfter.Faint,
	).Bold(
		config.Theme.TimestampAfter.Bold,
	)

	styleTimestampCursor = gloss.NewStyle().Foreground(
		gloss.Color(
			config.Theme.TimestampCursor.Color,
		),
	).Faint(
		config.Theme.TimestampCursor.Faint,
	).Bold(
		config.Theme.TimestampCursor.Bold,
	)
}
