package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m statusbarModel) View(model model) string {
	text := " Ready "
	if m.loading {
		text = " Loading… ⏳ "
	}

	return statusBarStyle.
		Width(model.width).
		Render(text)
}

func (m statusbarModel) Update(msg tea.Msg) (statusbarModel, tea.Cmd) {

	return m, nil
}
