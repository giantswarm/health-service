package health

import (
	"github.com/giantswarm/health-service/service/health/key"
)

// ClusterStatus holds data about the overall health/status of a cluster.
type ClusterStatus struct {
	Health key.Health         `json:"health"`
	State  key.LifecycleState `json:"state"`
}

// NodeStatus represents the health status of a single node in a cluster.
type NodeStatus struct {
	Ready                             bool       `json:"ready"`
	Health                            key.Health `json:"health"`
	Name                              string     `json:"name"`
	KubeletVersion                    string     `json:"kubelet_version"`
	OperatorVersion                   string     `json:"operator_version"`
	Role                              string     `json:"role"`
	IP                                string     `json:"ip"`
	Hostname                          string     `json:"hostname"`
	InstanceType                      string     `json:"instance_type"`
	AvailabilityRegion                string     `json:"availability_region"`
	AvailabilityZone                  string     `json:"availability_zone"`
	CPUCount                          int64      `json:"cpu_count"`
	MemoryCapacityBytes               int64      `json:"memory_capacity_bytes"`
	MemoryAllocatableBytes            int64      `json:"memory_allocatable_bytes"`
	EphemeralStorageCap               int64      `json:"ephemeral_storage_capacity_bytes"`
	EphemeralStorageAvail             int64      `json:"ephemeral_storage_allocatable_bytes"`
	AttachableVolumesAllocatableCount int64      `json:"attachable_volume_allocatable_count"`
	AttachableVolumesCapacityCount    int64      `json:"attachable_volume_capacity_count"`
}
