package health

// Response holds the data returned from the health searcher endpoint.
type Response struct {
	Cluster ClusterStatus `json:"cluster"`
	Nodes   []NodeStatus  `json:"nodes"`
}
