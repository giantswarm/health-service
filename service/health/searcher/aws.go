package searcher

import (
	"context"

	"github.com/giantswarm/microerror"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/health-service/service/health/key"
)

// searchAWSCR searches for the cluster config in AWSClusterConfigs resources.
func (s *Service) searchAWSCR(ctx context.Context, request Request) (*Response, error) {

	awsCR, err := s.g8sClient.ProviderV1alpha1().AWSConfigs("default").Get(request.ClusterID, v1.GetOptions{})
	if err != nil {
		return nil, microerror.Mask(err)
	}

	clusterHealth := key.HealthGreen // Optimistic default

	// TODO: Move to separate function
	// Check desired/actual node counts
	desiredNodes := awsCR.Status.Cluster.Scaling.DesiredCapacity
	currentNodes := len(awsCR.Status.Cluster.Nodes)

	if currentNodes < desiredNodes {
		clusterHealth = key.HealthYellow
	}

	clusterStatus := awsCR.ClusterStatus()
	if clusterStatus.HasDeletedCondition() || clusterStatus.HasDeletingCondition() {
		clusterHealth = key.HealthRed
	} else if clusterStatus.HasUpdatingCondition() || clusterStatus.HasCreatingCondition() { // TODO: Account for draining/scaling
		clusterHealth = key.HealthYellow
	}

	response := Response{
		ClusterHealth: clusterHealth,
	}

	return &response, nil
}
