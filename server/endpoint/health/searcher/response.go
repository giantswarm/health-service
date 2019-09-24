package searcher

// GeneralStatus holds data about the overall health/status of a cluster
type GeneralStatus struct {
	Health       string `json:"health"`
	Creating     bool   `json:"creating"`
	Upgrading    bool   `json:"upgrading"`
	Deleting     bool   `json:"deleting"`
	Normal       bool   `json:"normal"`
	NodeCount    int    `json:"nodeCount"`
	MaxNodeCount int    `json:"maxNodeCount"`
}

// NodeStatusComputeResources holds data about available or requested compute resources
type NodeStatusComputeResources struct {
	CPU      int `json:"cpu"`
	MemoryGB int `json:"memoryGb"`
}

// NodeStatusStorage holds data about the status of storage on a cluster node
type NodeStatusStorage struct {
	VolumeSizeGB        int `json:"volumeSizeGb"`
	AttachableDiskCount int `json:"attachableDiskCount"`
	AttachedDiskCount   int `json:"attachedDiskCount"`
	EphemeralStorageGB  int `json:"ephemeralStorageGb"`
}

// NodeStatusIdentity holds data about the identity of a node which is generally static
type NodeStatusIdentity struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	Version  string `json:"version"`
}

// NodeStatus represents the health status of a single node in a cluster
type NodeStatus struct {
	Ready            bool                       `json:"ready"`
	Identity         NodeStatusIdentity         `json:"identity"`
	MachineResources NodeStatusComputeResources `json:"machine"`
	LimitTotals      NodeStatusComputeResources `json:"limitTotals"`
	RequestTotals    NodeStatusComputeResources `json:"requestTotals"`
}

// Response holds the data returned from the health searcher endpoint
type Response struct {
	General GeneralStatus `json:"general"`
	Nodes   []NodeStatus  `json:"nodes,omitempty"`
}
