package config

func Default() Config {
	return Config{
		UI: struct {
			Theme  string
			Border string
		}{
			Theme:  "dark",
			Border: "rounded",
		},
		Colors: struct {
			Primary   string
			Secondary string
			Error     string
		}{
			Primary:   "#89b4fa",
			Secondary: "#a6adc8",
			Error:     "#f38ba8",
		},
	}
}
