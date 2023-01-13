package cgroup

type Config struct {
	CPUs   float64 `koanf:"cpus"`
	Memory int     `koanf:"memory"`
	Swap   int     `koanf:"swap"`
	PIDs   int     `koanf:"pids"`
}
