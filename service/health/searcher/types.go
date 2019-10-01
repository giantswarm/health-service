package searcher

import (
	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	v1 "k8s.io/api/core/v1"

	"github.com/giantswarm/health-service/service/health/key"
)

type clusterInfo struct {
	endpoint string
	nodes    []v1.Node
	status   v1alpha1.StatusCluster
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

func NewNodeStatus(node v1.Node) NodeStatus {
	ready := false
	for _, condition := range node.Status.Conditions {
		if condition.Type == v1.NodeReady {
			ready = condition.Status == "True"
			break
		}
	}
	return NodeStatus{
		Name:     node.Name,
		Ready:    ready,
		IP:       ipFromAddresses(node.Status.Addresses),
		Hostname: hostnameFromAddresses(node.Status.Addresses),
	}
}

func NewNodesStatus(cluster clusterInfo) []NodeStatus {
	result := []NodeStatus{}
	for _, node := range cluster.nodes {
		result = append(result, NewNodeStatus(node))
	}
	return result
}

func findInAddresses(addresses []v1.NodeAddress, address_type v1.NodeAddressType) string {
	for _, entry := range addresses {
		if entry.Type == address_type {
			return entry.Address
		}
	}
	return ""
}

func ipFromAddresses(addresses []v1.NodeAddress) string {
	return findInAddresses(addresses, v1.NodeInternalIP)
}

func hostnameFromAddresses(addresses []v1.NodeAddress) string {
	return findInAddresses(addresses, v1.NodeHostName)
}
