package key

// Health is an enumeration of colors which represent the overall health of a resource
type Health string

const (
	// Red represents an unhealthy cluster
	Red Health = "red"
	// Yellow represents a passable cluster
	Yellow Health = "yellow"
	// Green represents a healthy cluster
	Green Health = "green"
	// Default is used when a condition cannot otherwise be determined
	Default Health = Green
)

// LifecycleState is an enumeration of states correspoding to certain cluster conditions
type LifecycleState string

const (
	// Normal cluster state represents a cluster which is not creating, upgrading, or deleting
	Normal LifecycleState = "normal"
	// Creating cluster state represents a cluster which is still being created
	Creating LifecycleState = "creating"
	// Upgrading cluster state represents a cluster which is still being upgraded
	Upgrading LifecycleState = "upgrading"
	// Deleting cluster state represents a cluster which is still being deleting
	Deleting LifecycleState = "deleting"
)

// Providers is an array of providers which are currently implemented for health checking
var Providers = [...]string{"aws", "azure", "kvm"}

// IsKnownProvider is a helper function returning true iff the given value is in Providers
func IsKnownProvider(value string) bool {
	for i := range Providers {
		if Providers[i] == value {
			return true
		}
	}
	return false
}
