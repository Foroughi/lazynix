package tui

import "github.com/Foroughi/lazynix/src/config"

type Screen int

const (
	ScreenProfile Screen = iota
	ScreenFlake
)

type model struct {
	cfg    config.Config
	cursor int
	width  int
	height int

	screen Screen

	topbar    topbarModel
	statusbar statusbarModel
	profile   profileModel
	flake     flakeModel
}

type topbarModel struct{}
type statusbarModel struct {
	loading bool
}
type profileModel struct{}

type flakeModel struct {
	activeSection int
	sections      []listModel
}

type listModel struct {
	title  string
	width  int
	height int
	items  []itemModel
}

type itemModel struct {
	text     string
	selected bool
}
