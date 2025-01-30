package config

import (
	"log"
	"strings"
)

// Parse and split a list of key:value provided in the following forms :
// input: key1:value1;key2:value2;key3:value3
// output: ["key1:value1", "key2:value2", "key3:value3"]
func parseKeyValuePairs(portStr string) []string {
	if portStr == "" {
		return []string{}
	}

	// Separate each port in a combinaison of it source:dest
	portPairs := strings.Split(portStr, ";")
	var result []string

	// Remove space and add it to result array
	for _, portPair := range portPairs {
		portPair = strings.TrimSpace(portPair)
		result = append(result, portPair)
	}

	return result
}

// parseBoolEnv parses a string environment variable into a boolean value.
// Accepted values for true: "true", "yes", "True", "TRUE"
// Accepted values for false: "false", "no", "False", "FALSE", "" (empty string)
// Any other value causes the program to exit with an error.
func parseBoolEnv(value string) bool {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "true", "yes":
		return true
	case "false", "no", "":
		return false
	default:
		log.Fatalf("Invalid boolean value for ROLENV_PRIVILEGED: %s. Expected true/false or yes/no.", value)
		return false // Unreachable, but required for compilation
	}
}
