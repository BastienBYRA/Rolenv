package config

import (
	"fmt"

	"github.com/bastienbyra/rolenv/internal/docker"
	"github.com/bastienbyra/rolenv/internal/utils"
	"github.com/joho/godotenv"
)

// LoadConfig load the config file .env, and from it generate the spec of the Container to create
func LoadConfig(filename string) (*docker.ContainerConfig, error) {
	if filename == "" {
		filename = ".env"
	}

	envMap, err := godotenv.Read(filename)
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	config := &docker.ContainerConfig{
		Image:   utils.CheckEnvNotNull(envMap, "ROLENV_IMAGE"),
		Version: utils.CheckEnvNotNull(envMap, "ROLENV_VERSION"),
		Ports:   parsePorts(envMap["ROLENV_PORT"]),
	}

	fmt.Print(config)

	return config, nil
}
