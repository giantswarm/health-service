package searcher

import (
	"context"
)

// searchAWSCR searches for the cluster config in AWSClusterConfigs resources.
func (s *Service) searchAWSCR(ctx context.Context, request Request) (*Response, error) {
	/*
		awsClusterConfigName := fmt.Sprintf("%s-%s", request.ClusterID, "aws-cluster-config")
		awsClusterConfig, err := s.G8sClient.CoreV1alpha1().AWSClusterConfigs(metav1.NamespaceDefault).Get(awsClusterConfigName, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			return nil, microerror.Mask(clusterNotFoundError)
		} else if err != nil {
			return nil, microerror.Mask(err)
		}
	*/

	response := DefaultResponse()
	//response.ClusterHealth = awsClusterConfig.GetName()
	response.ClusterHealth = "green"

	return &response, nil
}
