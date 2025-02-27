package main

import (
	"github.com/bastienbyra/rolenv/cmd"
)

func main() {
	cmd.Execute()
}

// func main() {
// 	dockerConfig, _ := config.LoadConfig("")
// 	docker.Run(dockerConfig)
// }

// func validate() {
// 	dockerConfig, _ := config.LoadConfig("")
// 	docker.Validate(dockerConfig)
// }
