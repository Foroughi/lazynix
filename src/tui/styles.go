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

	inactiveSectionTitleStyle = lipgloss.NewStyle().
					Background(lipgloss.Color("#11111b")).
					Foreground(lipgloss.Color("#a6adc8"))

	inactiveSectionStyle = lipgloss.NewStyle()

	inactiveSectionItemStyle = lipgloss.NewStyle()

	activeSectionTitleStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("##96faa0")).
				Foreground(lipgloss.Color("#000000"))

	activeSectionStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("##96faa0")).
				Foreground(lipgloss.Color("#000000"))

	activeSectionItemStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("##96faa0")).
				Foreground(lipgloss.Color("#000000"))
)
