package utils

import (
	"log"
	"strings"
)

// checks if a specific environment variable exists in the provided map and is not empty.
// If the variable is not found, it logs a fatal error and exits the program.
// It returns the value of the environment variable if found.
func CheckEnvNotNullFromEnvFile(envMap map[string]string) func(string) string {
	return func(envVar string) string {
		value, exists := envMap[envVar]
		if !exists || strings.TrimSpace(value) == "" {
			log.Fatalf("The variable %s is not defined.", envVar)
		}
		return value
	}
}
