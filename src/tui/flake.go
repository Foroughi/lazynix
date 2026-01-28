package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m flakeModel) View(model model) string {
	height := model.height - 2 // top + status

	return lipgloss.NewStyle().
		Width(model.width).
		Height(height).
		Render(" MAIN Flake CONTENT ")
}

func (m flakeModel) Update(msg tea.Msg) (flakeModel, tea.Cmd) {

	return m, nil
}
