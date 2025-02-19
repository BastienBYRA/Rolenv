package docker

import (
	"github.com/docker/docker/api/types/container"
)

type ContainerConfig struct {
	Name          string
	Image         string
	Ports         []string
	Network       string
	Hosts         []string // --add-host
	Entrypoint    []string
	Command       []string
	Hostname      string
	Privileged    bool
	RestartPolicy container.RestartPolicy
	Volumes       []string
}
