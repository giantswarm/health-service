package health

import (
	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	v1 "k8s.io/api/core/v1"

	"github.com/giantswarm/health-service/service/host"
)

// Request is the configuration for the service action.
type Request struct {
	ClusterID string
	Cluster   Cluster
}

// Cluster represents the control plane and tenant cluster data received for performing health calculations
type Cluster struct {
	Nodes  []v1.Node
	Pods   []v1.Pod
	Status v1alpha1.StatusCluster
	Spec   host.ProviderSpec
}
