package searcher

import (
	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"

	"github.com/giantswarm/health-service/service/health/key"
)

type clusterInfo struct {
	status v1alpha1.StatusCluster
}

func NewClusterStatus(cluster clusterInfo) ClusterStatus {
	clusterHealth := key.HealthGreen // Optimistic default

	desiredNodes := cluster.status.Scaling.DesiredCapacity
	currentNodes := len(cluster.status.Nodes)

	if currentNodes < desiredNodes {
		clusterHealth = key.HealthYellow
	}

	creating := cluster.status.HasCreatingCondition()
	updating := cluster.status.HasUpdatingCondition()
	deleted := cluster.status.HasDeletedCondition()
	deleting := cluster.status.HasDeletingCondition()

	if deleted || deleting {
		clusterHealth = key.HealthRed
	} else if updating || creating { // TODO: Account for draining/scaling
		clusterHealth = key.HealthYellow
	}

	return ClusterStatus{
		Health:    clusterHealth,
		Creating:  creating,
		Upgrading: updating,
		Deleting:  deleting,
		Normal:    !creating && !updating && !deleting,
		NodeCount: currentNodes,
	}
}
