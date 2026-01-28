package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func column(view listModel, isActive bool) string {

	titleStyle := inactiveSectionTitleStyle.
		Bold(true).
		Width(view.width).
		Padding(0, 1)

	if isActive {

		titleStyle = activeSectionTitleStyle.
			Bold(true).
			Width(view.width).
			Padding(0, 1)
	}
	boxStyle := inactiveSectionStyle.
		Width(view.width).
		Height(view.height).
		Padding(0, 0)

	if isActive {

		boxStyle = activeSectionStyle.
			Width(view.width).
			Height(view.height).
			Padding(0, 0)
	}

	var lines []string
	lines = append(lines, titleStyle.Render(view.title))

	for i, it := range view.items {
		if i == view.activeItem {
			lines = append(lines, activeSectionItemStyle.Width(view.width).Padding(0, 1).Render(it.text))
		} else {

			lines = append(lines, inactiveSectionItemStyle.Width(view.width).Padding(0, 1).Render(it.text))
		}
	}

	return boxStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left, lines...),
	)
}

func (m flakeModel) View(model model) string {

	var sectionViews []string

	for i, it := range m.sections {

		if m.activeSection == i {
			sectionViews = append(sectionViews, column(it, true))
		} else {

			sectionViews = append(sectionViews, column(it, false))
		}
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		sectionViews...,
	)
}
func (m flakeModel) Update(msg tea.Msg) (flakeModel, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "l":
			m.activeSection++
		case "h":

			m.activeSection--

		case "j":
			m.sections[m.activeSection].activeItem++
		case "k":
			m.sections[m.activeSection].activeItem--
		}
	}

	if m.activeSection >= len(m.sections) {
		m.activeSection = 0
	}

	if m.activeSection < 0 {
		m.activeSection = len(m.sections) - 1
	}

	return m, nil
}
