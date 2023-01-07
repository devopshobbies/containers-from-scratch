package internal

const (
	Project = "cfs"

	LayersPath     = "/var/lib/" + Project + "/images/layers"
	ContainersPath = "/var/run/" + Project + "/containers"
	NetNSPath      = "/var/run/" + Project + "/netns"
)
