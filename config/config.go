package config

type Config struct {
	Listen []Listen
	UI     UI
}

type Listen struct {
	Addr string
}

type UI struct {
	Addr string
}
