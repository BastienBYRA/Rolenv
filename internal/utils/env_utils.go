package utils

import "log"

// checks if a specific environment variable exists in the provided map and is not empty.
// If the variable is not found, it logs a fatal error and exits the program.
// It returns the value of the environment variable if found.
func CheckEnvNotNull(envMap map[string]string, envMapName string) string {
	value, err := envMap[envMapName]

	if !err {
		log.Fatalf("The variable %s is not defined.", envMapName)
	}

	return value
}
