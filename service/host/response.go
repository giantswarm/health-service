package host

import v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"

// Response holds the data returned from the cluster searcher endpoint.
type Response struct {
	Endpoint string                 `json:"endpoint"`
	Status   v1alpha1.StatusCluster `json:"cluster"`
	Spec     ProviderSpec           `json:"provider_spec"`
}

type ProviderSpec struct {
	Workers []Worker
}

type Worker struct {
	DockerVolumeSizeGB int64 `json:"docker_volume_size_gb"`
}
