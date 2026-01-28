package tui

import (
	// "github.com/Foroughi/lazynix/ui/utils"
	"github.com/charmbracelet/lipgloss"
)

func (m model) topBar() string {

	right := " Profile "
	if m.screen == ScreenFlake {
		right = " Flake "
	}

	left := " Lazy Nix "

	bar := lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		left,
	)

	bar += lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Right,
		right,
	)

	return topBarStyle.Width(m.width).Render(bar)
}
