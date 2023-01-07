package config

import (
	"github.com/mohammadne/zar/pkg/cgroups"
	"github.com/mohammadne/zar/pkg/log"
)

type Config struct {
	Print    bool        `koanf:"print"`
	Log      *log.Config `koanf:"log"`
	Hostname string      `koanf:"hostname"`

	CGroups *cgroups.Config `koanf:"cgroups"`
}
