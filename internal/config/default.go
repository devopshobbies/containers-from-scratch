package config

import "github.com/mohammadne/zar/pkg/log"

func Default() *Config {
	return &Config{
		Print: true,
		Log: &log.Config{
			Development: true,
			Level:       "info",
			Encoding:    "console",
		},
	}
}
