package searcher

import "github.com/giantswarm/health-service/service/health/key"

// Response holds the data returned from the health searcher endpoint.
type Response struct {
	Cluster ClusterStatus `json:"cluster"`
}

// ClusterStatus holds data about the overall health/status of a cluster.
type ClusterStatus struct {
	Health    key.Health         `json:"health"`
	State     key.LifecycleState `json:"state"`
	NodeCount int                `json:"node_count"`
}
