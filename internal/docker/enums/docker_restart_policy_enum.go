package enums

// RestartPolicy represents all the restart strategy available in the Docker CLI
type RestartPolicy string

const (
	RestartNo            RestartPolicy = "no"
	RestartAlways        RestartPolicy = "always"
	RestartUnlessStopped RestartPolicy = "unless-stopped"
	RestartOnFailure     RestartPolicy = "on-failure"
)
