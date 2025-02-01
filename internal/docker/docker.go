package docker

import (
	"context"
	"fmt"
	"io"
	"os"

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
	resp, err := cli.ContainerCreate(ctx, contConfig, nil, nil, nil, cc.Name)
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

	// config.ExposedPorts = nat.PortSet{"truc": {}}
	return &config

}
