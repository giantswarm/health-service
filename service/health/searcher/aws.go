package searcher

import (
	"context"
	"fmt"

	"github.com/giantswarm/microerror"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// searchAWSCR searches for the cluster config in AWSClusterConfigs resources.
func (s *Service) searchKVMCR(ctx context.Context, request Request) (*Response, error) {
	awsClusterConfigName := fmt.Sprintf("%s-%s", request.ClusterID, "aws-cluster-config")
	awsClusterConfig, err := s.G8sClient.CoreV1alpha1().AWSClusterConfigs(metav1.NamespaceDefault).Get(awsClusterConfigName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		return nil, microerror.Mask(clusterNotFoundError)
	} else if err != nil {
		return nil, microerror.Mask(err)
	}

	response := DefaultResponse()
	response.ClusterHealth = awsClusterConfig.GetName()

	return &response, nil
}
