package config

import (
	"log"
	"strconv"
	"strings"
)

// Parse and split a list of key:value pairs provided in the following forms :
// Example:
//
//	input: key1:value1;key2:value2;key3:value3
//	output: ["key1:value1", "key2:value2", "key3:value3"]
func parseKeyValuePairs(inputStr string) []string {
	if inputStr == "" {
		return []string{}
	}

	// Separate each pair in a combination of key:value
	pairs := strings.Split(inputStr, ";")
	var result []string

	// Remove space and add it to result array
	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		result = append(result, pair)
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

// parseList parses a delimited string into a slice of strings.
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

// parsePositiveNumber converts a string to an integer.
// If the input is empty, non-numeric, or negative, it returns -1.
func parsePositiveNumber(value string) int {
	if value == "" {
		return -1
	}

	number, err := strconv.Atoi(value)
	if err != nil || number < 0 {
		return -1
	}

	return number
}
