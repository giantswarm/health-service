package health

import "github.com/giantswarm/health-service/service/health/key"

// Response holds the data returned from the health searcher endpoint.
type Response struct {
	Cluster ClusterStatus `json:"cluster"`
	Nodes   []NodeStatus  `json:"nodes"`
}

// ClusterStatus holds data about the overall health/status of a cluster.
type ClusterStatus struct {
	Health key.Health         `json:"health"`
	State  key.LifecycleState `json:"state"`
}

// NodeStatus holds data about the health/status of a node.
type NodeStatus struct {
	// Identity
	Name               string `json:"name"`
	Role               string `json:"role"`
	IP                 string `json:"ip"`
	Hostname           string `json:"hostname"`
	InstanceType       string `json:"instance_type"`
	AvailabilityRegion string `json:"availability_region"`
	AvailabilityZone   string `json:"availability_zone"`
	// Versions
	KubeletVersion  string `json:"kubelet_version"`
	OperatorVersion string `json:"operator_version"`
	// Resources
	CPUCount                          int64 `json:"cpu_count"`
	MemoryCapacityBytes               int64 `json:"memory_capacity_bytes"`
	MemoryAllocatableBytes            int64 `json:"memory_allocatable_bytes"`
	EphemeralStorageCap               int64 `json:"ephemeral_storage_capacity_bytes"`
	EphemeralStorageAvail             int64 `json:"ephemeral_storage_allocatable_bytes"`
	AttachableVolumesAllocatableCount int64 `json:"attachable_volume_allocatable_count"`
	AttachableVolumesCapacityCount    int64 `json:"attachable_volume_capacity_count"`
	// Status
	Ready  bool       `json:"ready"`
	Health key.Health `json:"health"`
	// Pods
	LimitTotals   NodeStatusComputeResources `json:"limit_totals"`
	RequestTotals NodeStatusComputeResources `json:"request_totals"`
}

// NodeStatusComputeResources holds data about available or requested compute resources.
type NodeStatusComputeResources struct {
	CPU         int64 `json:"cpu"`
	MemoryBytes int64 `json:"memory_bytes"`
}
