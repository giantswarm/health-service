package searcher

import (
	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"

	"github.com/giantswarm/health-service/service/health/key"
)

type clusterInfo struct {
	status v1alpha1.StatusCluster
}

func NewClusterStatus(cluster clusterInfo) ClusterStatus {
	health := key.Default

	desiredNodes := cluster.status.Scaling.DesiredCapacity
	currentNodes := len(cluster.status.Nodes)

	if currentNodes < desiredNodes {
		health = key.Yellow
	}

	creating := cluster.status.HasCreatingCondition()
	updating := cluster.status.HasUpdatingCondition()
	deleted := cluster.status.HasDeletedCondition()
	deleting := cluster.status.HasDeletingCondition()

	if deleted || deleting {
		health = key.Red
	} else if updating || creating { // TODO: Account for draining/scaling
		health = key.Yellow
	}

	state := key.Normal

	if creating {
		state = key.Creating
	} else if updating {
		state = key.Upgrading
	} else if deleting {
		state = key.Deleting
	}

	return ClusterStatus{
		Health:    health,
		State:     state,
		NodeCount: currentNodes,
	}
}
