package docker

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
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

	if cc.User != "" {
		config.User = cc.User
	}

	if len(cc.EnvList) > 0 {
		config.Env = cc.EnvList
	}

	return &config
}

func createContainerHostConfig(cc *ContainerConfig) *container.HostConfig {
	config := container.HostConfig{
		Privileged:     cc.Privileged,
		RestartPolicy:  cc.RestartPolicy,
		ReadonlyRootfs: cc.ReadonlyRootFS,
	}

	if len(cc.VolumeBinds) > 0 {
		config.Binds = cc.VolumeBinds
	}

	if len(cc.VolumeMounts) > 0 {
		config.Mounts = cc.VolumeMounts
	}

	if cc.MemoryHardLimit > 0 {
		config.Memory = cc.MemoryHardLimit
	}

	if cc.CPUCoreLimit > 0 {
		config.NanoCPUs = cc.CPUCoreLimit
	}

	if len(cc.SecurityOpts) > 0 {
		config.SecurityOpt = cc.SecurityOpts
	}

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
	}

	if len(cc.VolumeBinds) > 0 {
		fmt.Printf("- Volume Binds: %s\n", strings.Join(cc.VolumeBinds, ", "))
	}

	if len(cc.VolumeMounts) > 0 {
		listMount := []string{}
		for _, v := range cc.VolumeMounts {
			str := v.Source + ":" + v.Target
			listMount = append(listMount, str)
		}
		fmt.Printf("- Volume Mounts: %s\n", strings.Join(listMount, ", "))
	}

	if cc.Network != "" {
		fmt.Printf("- Network: %s\n", cc.Network)
	}

	fmt.Printf("- Restart Policy: %s\n", cc.RestartPolicy.Name)

	if len(cc.EnvList) > 0 {
		fmt.Printf("- Environment variables : %s\n", strings.Join(cc.EnvList, ", "))
	}

	if len(cc.Hosts) > 0 {
		fmt.Printf("- Hosts: %s\n", strings.Join(cc.Hosts, ", "))
	}

	if len(cc.Entrypoint) > 0 {
		fmt.Printf("- Entrypoint: %s\n", strings.Join(cc.Entrypoint, " "))
	}

	if len(cc.Command) > 0 {
		fmt.Printf("- Command: %s\n", strings.Join(cc.Command, " "))
	}

	if cc.User != "" {
		fmt.Printf("- User: %s\n", cc.User)
	}

	if cc.CPUCoreLimit > 0 {
		fmt.Printf("- CPU Limit: %v CPU cores\n", cc.CPUCoreLimit/1000000000)
	}

	if cc.MemoryHardLimit > 0 {
		fmt.Printf("- Memory Limit: %vm\n", cc.MemoryHardLimit/1000000)
	}

	fmt.Printf("- Readonly filesystem: %t\n", cc.ReadonlyRootFS)

	if cc.Hostname != "" {
		fmt.Printf("- Hostname: %s\n", cc.Hostname)
	}

	if cc.Privileged {
		fmt.Printf("- Privileged: %t\n", cc.Privileged)
	}

	if len(cc.SecurityOpts) > 0 {
		fmt.Printf("- Security options: %s\n", strings.Join(cc.Hosts, ", "))
	}

}
