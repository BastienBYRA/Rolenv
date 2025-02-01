package docker

import "github.com/bastienbyra/rolenv/internal/docker/enums"

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
	RestartPolicy enums.RestartPolicy
}
