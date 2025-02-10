package main

import (
	"github.com/bastienbyra/rolenv/internal/config"
	"github.com/bastienbyra/rolenv/internal/docker"
)

// func main() {
// 	dockerConfig, _ := config.LoadConfig("")
// 	docker.Run(dockerConfig)
// }

func main() {
	dockerConfig, _ := config.LoadConfig("")
	docker.Validate(dockerConfig)
}
