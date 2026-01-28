package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	topBarStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#1e1e2e")).
			Foreground(lipgloss.Color("#cdd6f4"))

	statusBarStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#11111b")).
			Foreground(lipgloss.Color("#a6adc8"))

	activeSectionTitleStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#04a5e5")).
				Foreground(lipgloss.Color("#ffffff"))
	inactiveSectionTitleStyle = lipgloss.NewStyle().
					Background(lipgloss.Color("#1e2030")).
					Foreground(lipgloss.Color("#ffffff"))

	activeSectionStyle   = lipgloss.NewStyle()
	inactiveSectionStyle = lipgloss.NewStyle()

	activeSectionItemStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#8aadf4")).
				Foreground(lipgloss.Color("#000000"))

	inactiveSectionItemStyle = lipgloss.NewStyle()
)
