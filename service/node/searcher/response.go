package searcher

import v1 "k8s.io/api/core/v1"

// Response holds the data returned from the node searcher endpoint.
type Response struct {
	Nodes []v1.Node `json:"nodes"`
}
