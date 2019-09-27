package searcher

// Response holds the data returned from the health searcher endpoint.
type Response struct {
	Cluster ClusterStatus `json:"cluster"`
}

// ClusterStatus holds data about the overall health/status of a cluster.
type ClusterStatus struct {
	Health    string `json:"health"`
	Creating  bool   `json:"creating"`
	Upgrading bool   `json:"upgrading"`
	Deleting  bool   `json:"deleting"`
	Normal    bool   `json:"normal"`
	NodeCount int    `json:"node_count"`
}
