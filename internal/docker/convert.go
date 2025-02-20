package docker

import (
	"log"
	"strings"

	"github.com/docker/docker/api/types/container"
)

// SetRestartPolicy converts a string representation of a Docker restart policy
// into its corresponding container.RestartPolicyMode.
//
// Accepted values:
//   - "no"              → container.RestartPolicyDisabled
//   - "always"          → container.RestartPolicyAlways
//   - "on-failure"      → container.RestartPolicyOnFailure
//   - "unless-stopped"  → container.RestartPolicyUnlessStopped
//
// If the provided policy string is not recognized, the function logs a fatal error
// and terminates the program.
//
// Example:
//
//	input:  "on-failure"
//	output: container.RestartPolicyOnFailure
func SetRestartPolicy(policy string, maxRetries int) container.RestartPolicy {

	// If policy is defined
	if policy != "" {
		validPolicies := map[string]container.RestartPolicyMode{
			string(container.RestartPolicyDisabled):      container.RestartPolicyDisabled,
			string(container.RestartPolicyAlways):        container.RestartPolicyAlways,
			string(container.RestartPolicyOnFailure):     container.RestartPolicyOnFailure,
			string(container.RestartPolicyUnlessStopped): container.RestartPolicyUnlessStopped,
		}

		if validPolicy, exists := validPolicies[policy]; exists {
			config := container.RestartPolicy{
				Name: validPolicy,
			}
			if maxRetries != -1 {
				config.MaximumRetryCount = maxRetries
			}
			return config
		}

		log.Fatalf("Unknown restart policy: %s. Available policies: no, always, on-failure, unless-stopped", policy)
	}

	// Otherwise, no policy
	return container.RestartPolicy{
		Name:              container.RestartPolicyDisabled,
		MaximumRetryCount: 0,
	}
}

func IsNamedVolume(volume string) bool {
	parts := strings.Split(volume, ":")
	if len(parts) != 2 {
		log.Fatalf("Invalid volume format: %s\n", volume)
	}

	source := parts[0]
	if strings.HasPrefix(source, "/") || strings.HasPrefix(source, "./") || strings.HasPrefix(source, "../") {
		return false
	}
	return true
}
