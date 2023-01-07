package cgroups

type Config struct {
	CPU    float64 `koanf:"cpu"`
	Memory int     `koanf:"memory"`
}
