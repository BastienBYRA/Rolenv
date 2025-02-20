package docker

import (
	"log"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
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
	VolumeBinds   []string
	VolumeMounts  []mount.Mount
}

// GuessVolumeType analyzes a list of Docker volumes provided as a semicolon-separated string
// and determines whether each volume is a named volume or a mounted volume.
//
// A mounted volume follows the pattern: "source_path:destination_path" where source_path
// is a relative or absolute path on the host.
//
// A named volume follows the pattern: "volume_name:destination_path" where volume_name
// is a named Docker volume.
//
// Example:
//
//		input:  "/home/test:/tmp/test;data-rolenv-test:/a-folder"
//		output:
//		  VolumeBinds = [data-rolenv-test:/a-folder]
//		  VolumeMounts = mount.Mount{
//				Type:   mount.TypeBind,
//				Source: /home/test,
//				Target: /tmp/test,
//	}
func (c *ContainerConfig) GuessVolumeType(volumes []string) {
	// Return "true" if the volume match with the named volume expectation, return "false" otherwise
	isNamedVolume := func(volume string) bool {
		parts := strings.Split(volume, ":")
		if len(parts) != 2 {
			log.Fatalf("Invalid volume format: %s\n", volume)
		}

		source := parts[0]
		if strings.HasPrefix(source, "/") || strings.HasPrefix(source, "./") || strings.HasPrefix(source, "../") {
			return false
		}
		return true
	}

	if len(volumes) > 0 {
		for _, volume := range volumes {

			if isNamedVolume(volume) {
				c.VolumeBinds = append(c.VolumeBinds, volume)
			} else {
				parts := strings.Split(volume, ":")
				mount := mount.Mount{
					Type:   mount.TypeBind,
					Source: parts[0],
					Target: parts[1],
				}
				c.VolumeMounts = append(c.VolumeMounts, mount)
			}
		}
	}
}
