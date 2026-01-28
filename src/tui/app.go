package tui

import (
	config "github.com/Foroughi/lazynix/src/config"

	tea "github.com/charmbracelet/bubbletea"
	"time"
)

func NewModel() model {

	return model{
		screen: ScreenFlake,
		cfg:    config.Default(),
		flake: flakeModel{

			sections: []listModel{
				{
					title:      "Inputs",
					widthRatio: 2,
					activeItem: 0,
					items: []itemModel{

						{
							text:     "item1",
							selected: false,
						},

						{
							text:     "item2",
							selected: false,
						},

						{
							text:     "item3",
							selected: false,
						},
					},
				},

				{
					title:      "Outputs",
					activeItem: 0,
					widthRatio: 5,
					items: []itemModel{

						{
							text:     "item1",
							selected: false,
						},

						{
							text:     "item2",
							selected: false,
						},

						{
							text:     "item3",
							selected: false,
						},
					},
				},

				{
					title:      "Systems",
					activeItem: 1,
					widthRatio: 0,
					items: []itemModel{
						{
							text:     "item1",
							selected: false,
						},

						{
							text:     "item2",
							selected: false,
						},

						{
							text:     "item3",
							selected: false,
						},
					},
				},
			},
		},
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
