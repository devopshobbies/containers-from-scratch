package config

import "github.com/mohammadne/zar/pkg/log"

func Default() *Config {
	return &Config{
		Print: false,
		Log: &log.Config{
			Development: true,
			Level:       "info",
			Encoding:    "console",
		},
	}
}
