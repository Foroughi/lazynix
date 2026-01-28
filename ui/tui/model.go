package tui

import "github.com/Foroughi/lazynix/ui/config"

type Screen int

const (
	ScreenProfile Screen = iota
	ScreenFlake
)

type model struct {
	cfg     config.Config
	screen  Screen
	loading bool
	cursor  int
	width   int
	height  int
}
