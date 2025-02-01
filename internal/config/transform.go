package config

import (
	"log"
	"strings"
)

// Parse and split a list of key:value provided in the following forms :
// Example:
//
//	input: key1:value1;key2:value2;key3:value3
//	output: ["key1:value1", "key2:value2", "key3:value3"]
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

// parseStringList parses a delimited string into a slice of strings.
// The input string consists of values separated by ";".
// Example:
//
//	input:  "apple;banana;cherry"
//	output: ["apple", "banana", "cherry"]
func parseList(inputStr string) []string {
	if inputStr == "" {
		return []string{}
	}

	// Split the input string by ";" separator
	items := strings.Split(inputStr, ";")
	var result []string

	// Trim spaces and add each item to the result slice
	for _, item := range items {
		item = strings.TrimSpace(item)
		result = append(result, item)
	}

	return result
}
