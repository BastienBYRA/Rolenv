package docker

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)

func Run(cc *ContainerConfig) {
	// Init
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// Pull the docker image
	imageName := cc.Image
	out, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)

	// Create the container
	contConfig := createContainerConfig(cc)
	hostConfig := createContainerHostConfig(cc)
	// netconf := network.NetworkingConfig

	resp, err := cli.ContainerCreate(ctx, contConfig, hostConfig, nil, nil, cc.Name)
	if err != nil {
		panic(err)
	}

	// Run the container
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
}

func createContainerConfig(cc *ContainerConfig) *container.Config {
	config := container.Config{
		Image: cc.Image,
	}

	if cc.Hostname != "" {
		config.Hostname = cc.Hostname
	}

	if len(cc.Entrypoint) != 0 {
		config.Entrypoint = cc.Entrypoint
	}

	if len(cc.Command) != 0 {
		config.Cmd = cc.Command
	}

	// mySet := make(map[string]struct{})
	// mySet["what"] = struct{}{}
	// config.Volumes = map[string]struct{}{
	// 	"/data": {},
	// }

	// config.ExposedPorts = nat.PortSet{"truc": {}}
	return &config

}

func createContainerHostConfig(cc *ContainerConfig) *container.HostConfig {
	config := container.HostConfig{
		Privileged:    cc.Privileged,
		RestartPolicy: cc.RestartPolicy,
	}

	if len(cc.Volumes) > 0 {

		binds := []string{}
		mounts := []mount.Mount{}

		for _, volume := range cc.Volumes {

			if IsNamedVolume(volume) {
				binds = append(binds, volume)
			} else {
				parts := strings.Split(volume, ":")
				mount := mount.Mount{
					Type:   mount.TypeBind,
					Source: parts[0],
					Target: parts[1],
				}
				mounts = append(mounts, mount)
			}
		}

		config.Binds = binds
		config.Mounts = mounts
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path + "\\fold")

	return &config
}

// While the name is "validate", the configuration is actually validate at it creation in LoadConfig
// We just print the configuration
func Validate(cc *ContainerConfig) {
	fmt.Println("Your container configuration:")
	fmt.Printf("- Name: %s\n", cc.Name)
	fmt.Printf("- Image: %s\n", cc.Image)

	if len(cc.Ports) > 0 {
		fmt.Printf("- Open Ports: %s\n", strings.Join(cc.Ports, ", "))
	} else {
		// fmt.Println("- Open Ports: None")
	}

	if cc.Network != "" {
		fmt.Printf("- Network: %s\n", cc.Network)
	} else {
		// fmt.Println("- Network: None")
	}

	if len(cc.Hosts) > 0 {
		fmt.Printf("- Hosts: %s\n", strings.Join(cc.Hosts, ", "))
	} else {
		// fmt.Println("- Hosts: None")
	}

	if len(cc.Entrypoint) > 0 {
		fmt.Printf("- Entrypoint: %s\n", strings.Join(cc.Entrypoint, " "))
	} else {
		// fmt.Println("- Entrypoint: Default")
	}

	if len(cc.Command) > 0 {
		fmt.Printf("- Command: %s\n", strings.Join(cc.Command, " "))
	} else {
		// fmt.Println("- Command: None")
	}

	if cc.Hostname != "" {
		fmt.Printf("- Hostname: %s\n", cc.Hostname)
	} else {
		// fmt.Println("- Hostname: None")
	}

	if cc.Privileged {
		fmt.Printf("- Privileged: %t\n", cc.Privileged)
	}

	fmt.Printf("- Restart Policy: %s\n", cc.RestartPolicy.Name)
}
