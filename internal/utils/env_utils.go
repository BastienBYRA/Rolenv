package utils

import (
	"log"
	"strconv"
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

func CheckEnvNotNullOrDefault(value string, defaultValue any) any {
	switch defaultValue.(type) {
	case string:
		if value != "" {
			return value
		}
		return defaultValue
	case int:
		valInt, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		return valInt
	case bool:
		valBool, err := strconv.ParseBool(value)
		if err != nil {
			return defaultValue
		}
		return valBool
	default:
		log.Fatalf("The type of defaultValue is not handled in the CheckEnvNotNullOrDefault function, type: %T", defaultValue)
		return "Shouldn't go here"
	}
}

func CheckEnvNotNullOrDefaultInt(value string, defaultValue int) int {
	res := CheckEnvNotNullOrDefault(value, defaultValue)
	resInt, isOk := res.(int)
	if !isOk {
		log.Fatalf("Error converting the interface value to int, value: %v", res)
	}
	return resInt
}

func CheckEnvNotNullOrDefaultBool(value string, defaultValue bool) bool {
	res := CheckEnvNotNullOrDefault(value, defaultValue)
	resBool, isOk := res.(bool)
	if !isOk {
		log.Fatalf("Error converting the interface value to bool, value: %v", res)
	}
	return resBool
}
