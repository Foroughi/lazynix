package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	topBarStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#1e1e2e")).
			Foreground(lipgloss.Color("#cdd6f4")).
			Padding(0, 1)

	statusBarStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#11111b")).
			Foreground(lipgloss.Color("#a6adc8")).
			Padding(0, 1)
)
