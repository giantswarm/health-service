package key

// Health is an enumeration of colors which represent the overall health of a resource
type Health string

const (
	// HealthRed represents an unhealthy cluster
	HealthRed Health = "red"
	// HealthYellow represents a passable cluster
	HealthYellow Health = "yellow"
	// HealthGreen represents a healthy cluster
	HealthGreen Health = "green"
	// HealthDefault is used when a condition cannot otherwise be determined
	HealthDefault Health = HealthGreen
)

// LifecycleState is an enumeration of states correspoding to certain cluster conditions
type LifecycleState string

const (
	// StateNormal cluster state represents a cluster which is not creating, upgrading, or deleting
	StateNormal LifecycleState = "normal"
	// StateCreating cluster state represents a cluster which is still being created
	StateCreating LifecycleState = "creating"
	// StateUpgrading cluster state represents a cluster which is still being upgraded
	StateUpgrading LifecycleState = "upgrading"
	// StateDeleting cluster state represents a cluster which is still being deleting
	StateDeleting LifecycleState = "deleting"
)
