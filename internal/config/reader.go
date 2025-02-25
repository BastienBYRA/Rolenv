package config

import (
	"fmt"
	"strings"

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
		Name:            checkEnvNotNull("ROLENV_NAME"),
		Image:           checkEnvNotNull("ROLENV_IMAGE"),
		Ports:           parseKeyValuePairs(envMap["ROLENV_PORTS"]),
		Network:         envMap["ROLENV_NETWORK"],
		Hosts:           parseKeyValuePairs(envMap["ROLENV_HOSTS"]),
		Entrypoint:      parseList(envMap["ROLENV_ENTRYPOINT"]),
		Command:         parseList(envMap["ROLENV_COMMAND"]),
		Hostname:        envMap["ROLENV_HOSTNAME"],
		Privileged:      parseBoolEnv(envMap["ROLENV_PRIVILEGED"]),
		RestartPolicy:   docker.SetRestartPolicy(envMap["ROLENV_RESTART_POLICY"], parsePositiveNumber(envMap["ROLENV_RESTART_POLICY_MAX_RETRIES"])),
		User:            envMap["ROLENV_USER"],
		EnvList:         getContainerEnvVars(envMap),
		MemoryHardLimit: int64(utils.CheckEnvNotNullOrDefaultInt(envMap["ROLENV_MEMORY_LIMIT"], 0)),
		CPUCoreLimit:    int64(utils.CheckEnvNotNullOrDefaultInt(envMap["ROLENV_CPU_CORE_LIMIT"], 0)),
		ReadonlyRootFS:  utils.CheckEnvNotNullOrDefaultBool(envMap["ROLENV_READONLY"], false),
		SecurityOpts:    parseList(envMap["ROLENV_SECURITY_OPTS"]),
	}

	// GuessVolumeType identifies user-defined volumes as named or bind mounts and assigns them to VolumeBinds or VolumeMounts.
	config.GuessVolumeType(parseKeyValuePairs(envMap["ROLENV_VOLUMES"]))

	// converts the CPU core limit to the equivalent value in nanocores as expected by Docker
	config.ConvertCpuToNanoCores()
	// converts the RAM limit to the equivalent value in bytes as expected by Docker
	config.ConvertMemoryMegabytesToBytes()

	return config, nil
}

// getContainerEnvVars filters environment variables from a provided map,
// separating standard environment variables from those prefixed with "ROLENV_".
//
// Standard environment variables are those that do not start with the "ROLENV_"
// prefix.
//
// Example:
//
//	input:  map[string]string{
//		"ROLENV_IMAGE": "nginx:latest",
//		"ROLENV_PORT":  "80:80",
//		"DB_HOST":      "localhost",
//		"DB_USER":      "root",
//		"APP_ENV":      "production",
//	}
//	output: []string{
//		"DB_HOST=localhost",
//		"DB_USER=root",
//		"APP_ENV=production",
//	}
func getContainerEnvVars(envMap map[string]string) []string {
	containerEnvList := []string{}

	for k, v := range envMap {
		if !strings.HasPrefix(k, "ROLENV_") {
			containerEnvList = append(containerEnvList, k+"="+v)
		}
	}

	return containerEnvList
}
