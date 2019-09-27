package searcher

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
	Name  string `jsno:"name"`
	Ready bool   `json:"ready"`
}
