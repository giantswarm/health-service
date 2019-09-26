package key

const (
	// HealthRed represents an unhealthy cluster
	HealthRed = "red"
	// HealthYellow represents a passable cluster
	HealthYellow = "yellow"
	//HealthGreen represents a healthy cluster
	HealthGreen = "green"
	// HealthDefault is used when a condition cannot otherwise be determined
	HealthDefault = "green"
)

var Providers = [...]string{"aws", "azure", "kvm"}

func IsKnownProvider(value string) bool {
	for i := range Providers {
		if Providers[i] == value {
			return true
		}
	}
	return false
}
