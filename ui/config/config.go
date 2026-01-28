package config

type Config struct {
	UI struct {
		Theme  string
		Border string
	}

	Colors struct {
		Primary   string
		Secondary string
		Error     string
	}
}
