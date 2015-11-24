package config

import (
	"github.com/magiconair/properties"
)

func FromFile(filename string) (*Config, error) {
	p, err := properties.LoadFile(filename, properties.UTF8)
	if err != nil {
		return nil, err
	}
	return FromProperties(p)
}

func FromProperties(p *properties.Properties) (*Config, error) {
	var cfg *Config = &Config{}
	var l Listen

	l.Addr = p.GetString("receiver.addr", DefaultConfig.Listen[0].Addr)
	cfg.Listen = append(cfg.Listen, l)

	return cfg, nil
}
