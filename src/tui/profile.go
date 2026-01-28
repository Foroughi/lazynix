package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m profileModel) View(model model) string {
	height := model.height - 2 // top + status

	return lipgloss.NewStyle().
		Width(model.width).
		Height(height).
		Render(" MAIN Profile CONTENT ")
}

func (m profileModel) Update(msg tea.Msg) (profileModel, tea.Cmd) {

	return m, nil
}
