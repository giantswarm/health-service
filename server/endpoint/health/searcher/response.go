package searcher

// GeneralStatus holds data about the overall health/status of a cluster.
type GeneralStatus struct {
	Health       string `json:"health"`
	Creating     bool   `json:"creating"`
	Upgrading    bool   `json:"upgrading"`
	Deleting     bool   `json:"deleting"`
	Normal       bool   `json:"normal"`
	NodeCount    int    `json:"node_count"`
	MaxNodeCount int    `json:"max_node_count"`
}

// NodeStatusComputeResources holds data about available or requested compute resources.
type NodeStatusComputeResources struct {
	CPU      int `json:"cpu"`
	MemoryGB int `json:"memory_gb"`
}

// NodeStatusStorage holds data about the status of storage on a cluster node.
type NodeStatusStorage struct {
	VolumeSizeGB        int `json:"volume_size_gb"`
	AttachableDiskCount int `json:"attachable_disk_count"`
	AttachedDiskCount   int `json:"attached_disk_count"`
	EphemeralStorageGB  int `json:"ephemeral_storage_gb"`
}

// NodeStatusIdentity holds data about the identity of a node which is generally static.
type NodeStatusIdentity struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	Version  string `json:"version"`
}

// NodeStatus represents the health status of a single node in a cluster.
type NodeStatus struct {
	Ready            bool                       `json:"ready"`
	Identity         NodeStatusIdentity         `json:"identity"`
	MachineResources NodeStatusComputeResources `json:"machine"`
	LimitTotals      NodeStatusComputeResources `json:"limit_totals"`
	RequestTotals    NodeStatusComputeResources `json:"request_totals"`
}

// Response holds the data returned from the health searcher endpoint.
type Response struct {
	General GeneralStatus `json:"general"`
	Nodes   []NodeStatus  `json:"nodes,omitempty"`
}
