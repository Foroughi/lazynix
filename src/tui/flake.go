package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func column(title string, items []string, width, height int) string {

	titleStyle := inactiveSectionTitleStyle.
		Bold(true).
		Width(width).
		Padding(0, 1)

	boxStyle := inactiveSectionStyle.
		Width(width).
		Height(height).
		Padding(0, 0)

	var lines []string
	lines = append(lines, titleStyle.Render(title))

	for _, it := range items {
		lines = append(lines, inactiveSectionItemStyle.Padding(0, 1).Render(it))
	}

	return boxStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left, lines...),
	)
}

func (m flakeModel) View(model model) string {
	height := model.height - 2

	colWidth1 := model.width * 2 / 12
	colWidth2 := model.width * 5 / 12
	colWidth3 := (model.width*5 + 11) / 12

	left := column(
		"Inputs",
		[]string{"nixpkgs", "home-manager", "flake-utils"},
		colWidth1,
		height,
	)

	middle := column(
		"Outputs",
		[]string{"packages", "apps", "devShells"},
		colWidth2,
		height,
	)

	right := column(
		"Systems",
		[]string{"x86_64-linux", "aarch64-linux"},
		colWidth3,
		height,
	)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		middle,
		right,
	)
}
func (m flakeModel) Update(msg tea.Msg) (flakeModel, tea.Cmd) {

	return m, nil
}
