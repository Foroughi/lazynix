package tui

import ()

func (m model) statusBar() string {
	text := " Ready "
	if m.loading {
		text = " Loading… ⏳ "
	}

	return statusBarStyle.
		Width(m.width).
		Render(text)
}
