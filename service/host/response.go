package host

import v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"

// Response holds the data returned from the cluster searcher endpoint.
type Response struct {
	Endpoint string                 `json:"endpoint"`
	Status   v1alpha1.StatusCluster `json:"cluster"`
}
