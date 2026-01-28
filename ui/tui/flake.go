package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) flakeView() string {
	height := m.height - 2 // top + status

	return lipgloss.NewStyle().
		Width(m.width).
		Height(height).
		Render(" MAIN Flake CONTENT ")
}
