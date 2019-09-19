package project

var (
	description string = "This is a service to get the health of Giant Swarm tenant clusters."
	gitSHA             = "n/a"
	name        string = "health-service"
	source      string = "https://github.com/giantswarm/health-service"
	version            = "n/a"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
