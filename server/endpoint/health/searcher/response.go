package searcher

type GeneralStatus struct {
	Health       string `json:"health"`
	Creating     bool   `json:"creating"`
	Upgrading    bool   `json:"upgrading"`
	Deleting     bool   `json:"deleting"`
	Normal       bool   `json:"normal"`
	NodeCount    int    `json:"nodeCount"`
	MaxNodeCount int    `json:"maxNodeCount"`
}

type NodeStatus struct {
	Ready                bool   `json:"ready"`
	IP                   string `json:"ip"`
	Hostname             string `json:"hostname"`
	Version              string `json:"version"`
	CPUs                 int    `json:"cpus"`
	VolumeSizeGB         int    `json:"volumeSizeGb"`
	MemoryGB             int    `json:"memoryGb"`
	CPULimitTotal        int    `json:"cpuLimitTotal"`
	MemoryGBLimitTotal   int    `json:"memoryGbLimitTotal"`
	CPURequestTotal      int    `json:"cpuRequestTotal"`
	MemoryGBRequestTotal int    `json:"memoryGbRequestTotal"`
	AttachableDiskCount  int    `json:"attachableDiskCount"`
	AttachedDiskCount    int    `json:"attachedDiskCount"`
	EphemeralStorageGB   int    `json:"ephemeralStorageGb"`
}

type Response struct {
	General GeneralStatus `json:"general"`
	Nodes   []NodeStatus  `json:"nodes"`
}
