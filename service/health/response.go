package health

import "github.com/giantswarm/health-service/service/health/key"

// Response holds the data returned from the health searcher endpoint.
type Response struct {
	Cluster ClusterStatus `json:"cluster"`
	Nodes   []NodeStatus  `json:"nodes"`
}

// ClusterStatus holds data about the overall health/status of a cluster.
type ClusterStatus struct {
	Health    key.Health         `json:"health"`
	State     key.LifecycleState `json:"state"`
	NodeCount int                `json:"node_count"`
}

// NodeStatus holds data about the health/status of a node.
type NodeStatus struct {
	Name                   string `json:"name"`
	Ready                  bool   `json:"ready"`
	IP                     string `json:"ip"`
	Hostname               string `json:"hostname"`
	InstanceType           string `json:"instance_type"`
	AvailabilityRegion     string `json:"availability_region"`
	AvailabilityZone       string `json:"availability_zone"`
	KubeletVersion         string `json:"kubelet_version"`
	OperatorVersion        string `json:"operator_version"`
	CPUCount               int64  `json:"cpu_count"`
	MemoryCapacityBytes    int64  `json:"memory_capacity_bytes"`
	MemoryAllocatableBytes int64  `json:"memory_allocatable_bytes"`
	EphemeralStorageCap    int64  `json:"ephemeral_storage_capacity_bytes"`
	EphemeralStorageAvail  int64  `json:"ephemeral_storage_allocatable_bytes"`
}
