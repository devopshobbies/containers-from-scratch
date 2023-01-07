package config

import (
	"github.com/mohammadne/zar/pkg/cgroups"
	"github.com/mohammadne/zar/pkg/log"
)

func Default() *Config {
	return &Config{
		Print: false,
		Log: &log.Config{
			Development: true,
			Level:       "info",
			Encoding:    "console",
		},
		CGroups: &cgroups.Config{
			Memory: 200,
			CPUs:   1,
			Swap:   1,
			PIDs:   128,
		},
	}
}
