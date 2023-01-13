package config

import (
	"github.com/devopshobbies/containers-from-scratch/pkg/cgroup"
	"github.com/devopshobbies/containers-from-scratch/pkg/log"
)

type Config struct {
	Print    bool           `koanf:"print"`
	Log      *log.Config    `koanf:"log"`
	Hostname string         `koanf:"hostname"`
	CGroup   *cgroup.Config `koanf:"cgroup"`
}
