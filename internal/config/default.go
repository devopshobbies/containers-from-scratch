package config

import (
	"github.com/devopshobbies/containers-from-scratch/pkg/cgroups"
	"github.com/devopshobbies/containers-from-scratch/pkg/log"
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
