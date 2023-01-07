package config

import (
	"github.com/devopshobbies/containers-from-scratch/pkg/cgroups"
	"github.com/devopshobbies/containers-from-scratch/pkg/log"
)

type Config struct {
	Print    bool        `koanf:"print"`
	Log      *log.Config `koanf:"log"`
	Hostname string      `koanf:"hostname"`

	CGroups *cgroups.Config `koanf:"cgroups"`
}
