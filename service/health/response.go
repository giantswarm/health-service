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
	Name         string     `json:"name"`
	Ready        bool       `json:"ready"`
	Health       key.Health `json:"health"`
	IP           string     `json:"ip"`
	Hostname     string     `json:"hostname"`
	InstanceType string     `json:"instance_type"`
	CPUCount     int64      `json:"cpu_count"`
	MemoryBytes  int64      `json:"available_memory_bytes"`
	Role         string     `json:"role"`
}
