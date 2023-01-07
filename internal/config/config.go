package config

import "github.com/mohammadne/zar/pkg/log"

type Config struct {
	Print bool        `koanf:"print"`
	Log   *log.Config `koanf:"log"`
}
