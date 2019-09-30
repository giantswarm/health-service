package searcher

import (
	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	v1 "k8s.io/api/core/v1"
)

// Request is the configuration for the service action.
type Request struct {
	ClusterID string
	Nodes     []v1.Node
	Cluster   v1alpha1.StatusCluster
}
