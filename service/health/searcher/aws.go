package searcher

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// searchAWSCR searches for the cluster config in AWSClusterConfigs resources.
func (s *Service) searchAWSCR(ctx context.Context, request Request) (*Response, error) {
	response := Response{
		ClusterHealth: "greenio",
	}
	// Use clusteroperator snippet as a model to create k8s client
	// Use Go client for K8s to query AWSClusterConfigs
	awsCR, err := s.G8sClient.ProviderV1alpha1().AWSConfigs("default").Get(request.ClusterID, v1.GetOptions{})

	fmt.Println(awsCR, err)
	return &response, nil
}
