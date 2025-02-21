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

	// Ensure we can read the file
	envMap, err := godotenv.Read(filename)
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	// Generate the container configuration from it
	checkEnvNotNull := utils.CheckEnvNotNullFromEnvFile(envMap)
	config := &docker.ContainerConfig{
		Name:          checkEnvNotNull("ROLENV_NAME"),
		Image:         checkEnvNotNull("ROLENV_IMAGE"),
		Ports:         parseKeyValuePairs(envMap["ROLENV_PORTS"]),
		Network:       envMap["ROLENV_NETWORK"],
		Hosts:         parseKeyValuePairs(envMap["ROLENV_HOSTS"]),
		Entrypoint:    parseList(envMap["ROLENV_ENTRYPOINT"]),
		Command:       parseList(envMap["ROLENV_COMMAND"]),
		Hostname:      envMap["ROLENV_HOSTNAME"],
		Privileged:    parseBoolEnv(envMap["ROLENV_PRIVILEGED"]),
		RestartPolicy: docker.SetRestartPolicy(envMap["ROLENV_RESTART_POLICY"], parsePositiveNumber(envMap["ROLENV_RESTART_POLICY_MAX_RETRIES"])),
		User:          envMap["ROLENV_USER"],
	}
	// Initialize the volumes
	config.GuessVolumeType(parseKeyValuePairs(envMap["ROLENV_VOLUMES"]))

	return config, nil
}
