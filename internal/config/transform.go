package config

import "strings"

func parsePorts(portStr string) []string {
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
