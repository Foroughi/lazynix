package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "f":
			m.screen = ScreenFlake
		case "p":
			m.screen = ScreenProfile
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	var cmds []tea.Cmd

	// always-on components
	var topbarCmd tea.Cmd
	m.topbar, topbarCmd = m.topbar.Update(msg)
	cmds = append(cmds, topbarCmd)

	var statusbarCmd tea.Cmd
	m.statusbar, statusbarCmd = m.statusbar.Update(msg)
	cmds = append(cmds, statusbarCmd)

	switch m.screen {
	case ScreenProfile:

		var cmd tea.Cmd
		m.profile, cmd = m.profile.Update(msg)
		cmds = append(cmds, cmd)
	case ScreenFlake:

		var cmd tea.Cmd

		m.flake, cmd = m.flake.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
//
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "q":
// 			return m, tea.Quit
// 		case "f":
// 			m.screen = ScreenFlake
// 		case "p":
// 			m.screen = ScreenProfile
// 		}
//
// 	case tea.WindowSizeMsg:
// 		m.width = msg.Width
// 		m.height = msg.Height
// 	}
//
// 	return m, nil
// }
