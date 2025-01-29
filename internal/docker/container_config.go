package docker

type ContainerConfig struct {
	Image   string
	Version string
	Ports   []string
}
