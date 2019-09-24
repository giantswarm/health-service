package searcher

import (
	"context"

	"github.com/giantswarm/microerror"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// searchKVMCR searches for the cluster config in AWSClusterConfigs resources.
func (s *Service) searchKVMCR(ctx context.Context, request Request) (*Response, error) {
	kvmCR, err := s.g8sClient.ProviderV1alpha1().KVMConfigs("default").Get(request.ClusterID, v1.GetOptions{})
	if err != nil {
		return nil, microerror.Mask(err)
	}
	var health string
	{
		nodeCount := len(kvmCR.Status.Cluster.Nodes)
		if nodeCount == 4 {
			health = "green"
		} else if nodeCount > 0 {
			health = "yellow"
		} else {
			health = "red"
		}
	}
	response := Response{
		ClusterHealth: health,
	}

	return &response, nil
}
