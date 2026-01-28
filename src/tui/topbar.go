package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m topbarModel) View(model model) string {

	right := " Profile "
	if model.screen == ScreenFlake {
		right = " Flake "
	}

	left := " Lazy Nix "

	bar := lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.PlaceHorizontal(model.width/2, lipgloss.Left, left),
		lipgloss.PlaceHorizontal(model.width/2, lipgloss.Right, right),
	)

	return topBarStyle.Render(bar)
}

func (m topbarModel) Update(msg tea.Msg) (topbarModel, tea.Cmd) {

	return m, nil
}
