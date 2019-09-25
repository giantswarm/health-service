package searcher

import (
	"context"

	"github.com/giantswarm/health-service/service/health/key"

	"github.com/giantswarm/microerror"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// searchAWSCR searches for the cluster config in AWSClusterConfigs resources.
func (s *Service) searchAWSCR(ctx context.Context, request Request) (*Response, error) {

	awsCR, err := s.g8sClient.ProviderV1alpha1().AWSConfigs("default").Get(request.ClusterID, v1.GetOptions{})
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var clusterHealth = key.HealthGreen // Optimistic default

	// TODO: Move to separate function
	// Check desired/actual node counts
	var desiredNodes int = awsCR.Status.Cluster.Scaling.DesiredCapacity
	var currentNodes int = len(awsCR.Status.Cluster.Nodes)

	if currentNodes < desiredNodes {
		clusterHealth = key.HealthYellow
	}

	var clusterStatus = awsCR.ClusterStatus()
	if clusterStatus.HasDeletedCondition() || clusterStatus.HasDeletingCondition() {
		clusterHealth = key.HealthRed
	} else if clusterStatus.HasUpdatingCondition() || clusterStatus.HasCreatingCondition() { // TODO: Account for draining/scaling
		clusterHealth = key.HealthYellow
	} else if clusterStatus.HasCreatedCondition() || clusterStatus.HasUpdatedCondition() {
		// Don't set back to Green from Yellow or Red
	}

	response := Response{
		ClusterHealth: clusterHealth,
	}

	return &response, nil
}
