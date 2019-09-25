package searcher

import (
	"context"

	"github.com/giantswarm/microerror"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AWSGreenConditions contains the values for a cluster state which are considered 'green'
var AWSGreenConditions = map[string]struct{}{
	"Creating": {}, "Created": {}, "Updated": {}}

// AWSYellowConditions contains the values for a cluster state which are considered 'yellow'
var AWSYellowConditions = map[string]struct{}{
	"Scaling": {}, "Draining": {}, "Drained": {}, "Updating": {}}

// AWSRedConditions contains the values for a cluster state which are considered 'red'
var AWSRedConditions = map[string]struct{}{
	"Deleting": {}, "Deleted": {}}

const (
	// HealthRed represents an unhealthy cluster
	HealthRed = "red"
	// HealthYellow represents a passable cluster
	HealthYellow = "yellow"
	//HealthGreen represents a healthy cluster
	HealthGreen = "green"
)

// searchAWSCR searches for the cluster config in AWSClusterConfigs resources.
func (s *Service) searchAWSCR(ctx context.Context, request Request) (*Response, error) {

	awsCR, err := s.g8sClient.ProviderV1alpha1().AWSConfigs("default").Get(request.ClusterID, v1.GetOptions{})
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var clusterHealth string = HealthGreen // Optimistic default

	// Check desired/actual node counts
	var desiredNodes int = awsCR.Status.Cluster.Scaling.DesiredCapacity
	var currentNodes int = len(awsCR.Status.Cluster.Nodes)

	if currentNodes < desiredNodes {
		clusterHealth = HealthYellow
	}

	// Examine Cluster conditions and use the worst as my state
	for _, condition := range awsCR.Status.Cluster.Conditions {

		_, isGreen := AWSGreenConditions[condition.Type]
		if isGreen {
			// Don't set back to green from Yellow or Red state
		}

		_, isYellow := AWSYellowConditions[condition.Type]
		if isYellow {
			// Set health to Yellow if it isn't already Red
			if clusterHealth != HealthRed {
				clusterHealth = HealthYellow
			}
		}

		_, isRed := AWSRedConditions[condition.Type]
		if isRed {
			clusterHealth = HealthRed
		}
	}

	response := Response{
		ClusterHealth: clusterHealth,
	}

	return &response, nil
}
