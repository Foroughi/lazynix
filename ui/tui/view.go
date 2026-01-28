package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) mainView() string {

	if m.screen == ScreenFlake {
		return m.flakeView()
	} else {
		return m.profileView()
	}
}

func (m model) View() string {
	if m.width == 0 {
		return ""
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		m.topBar(),
		m.mainView(),
		m.statusBar(),
	)
}
