package docker

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
)

type ContainerConfig struct {
	Name            string // Docker CLI equivalent : --name ; ROLENV variable name : ROLENV_NAME
	Image           string
	Ports           []string
	Network         string
	Hosts           []string // --add-host
	Entrypoint      []string
	Command         []string
	Hostname        string
	Privileged      bool
	RestartPolicy   container.RestartPolicy
	VolumeBinds     []string
	VolumeMounts    []mount.Mount
	User            string
	EnvList         []string
	MemoryHardLimit int64
	CPUCoreLimit    int64
	ReadonlyRootFS  bool
	SecurityOpts    []string
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

			// If the volume is a named one
			if isNamedVolume(volume) {
				c.VolumeBinds = append(c.VolumeBinds, volume)
			} else {
				// If the volume is a local one (mounted)
				parts := strings.Split(volume, ":")

				absolutePath, err := filepath.Abs(parts[0])
				if err != nil {
					log.Fatalf("Error occurred during path conversion : %v", err)
				}

				mount := mount.Mount{
					Type:   mount.TypeBind,
					Source: absolutePath,
					Target: parts[1],
				}
				c.VolumeMounts = append(c.VolumeMounts, mount)
			}
		}
	}
}

// convertCpuToNanoCores converts the CPU core limit from a user-provided value
// to the equivalent value in nanocores, as expected by Docker.
//
// Example:
//
//	input:  c.CPUCoreLimit = 2.5 (representing 2.5 CPU cores)
//	output: c.CPUCoreLimit = 2500000000 (2.5 billion nanocores)
func (c *ContainerConfig) ConvertCpuToNanoCores() {
	if c.CPUCoreLimit > 0 {
		// Convert CPU core limit to nanocores by multiplying by 1 billion (10^9)
		c.CPUCoreLimit = c.CPUCoreLimit * 1000000000
	}
}

func (c *ContainerConfig) ConvertMemoryMegabytesToBytes() {
	if c.MemoryHardLimit > 0 {
		c.MemoryHardLimit = c.MemoryHardLimit * 1000000
	}
}
