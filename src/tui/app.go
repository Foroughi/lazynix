package tui

import (
	config "github.com/Foroughi/lazynix/src/config"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

func NewModel() model {
	return model{
		cfg: config.Default(),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Tick(time.Millisecond*300, func(t time.Time) tea.Msg {
		return "Loading..."
	})
}

func Run() error {
	m := NewModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err := p.Run()
	return err
}
