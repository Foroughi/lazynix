package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) mainView() string {

	if m.screen == ScreenFlake {
		return m.flake.View(m)
	} else {
		return m.profile.View(m)
	}
}

func (m model) View() string {
	if m.width == 0 {
		return ""
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		m.topbar.View(m),
		m.mainView(),
		m.statusbar.View(m),
	)
}
