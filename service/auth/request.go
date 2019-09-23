package auth

// Request is the configuration for the service action.
type Request struct {
	ClusterID string   `json:"clusterID"`
	User      string   `json:"user"`
	Groups    []string `json:"groups"`
}
