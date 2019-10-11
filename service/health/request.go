package health

import (
	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	v1 "k8s.io/api/core/v1"
)

// Request holds the data required to generate the health response.
type Request struct {
	Cluster   v1alpha1.StatusCluster
	ClusterID string
	Nodes     []v1.Node
	Pods      []v1.Pod
}
