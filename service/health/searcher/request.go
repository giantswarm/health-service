package searcher

// Request is the configuration for the service action.
type Request struct {
	ClusterID string `json:"clusterID"`
}

// DefaultRequest provides a default request by best effort.
func DefaultRequest() Request {
	return Request{
		ClusterID: "",
	}
}
