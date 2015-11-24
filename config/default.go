package config

var DefaultConfig = &Config{
	Listen: []Listen{
		{
			Addr: ":5858",
		},
	},
	UI: UI{
		Addr: ":5859",
	},
}
