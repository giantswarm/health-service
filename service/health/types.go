package health

import (
	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/giantswarm/health-service/service/health/key"
)

func NewClusterStatus(cluster v1alpha1.StatusCluster) ClusterStatus {
	health := key.Default

	desiredNodes := cluster.Scaling.DesiredCapacity
	currentNodes := len(cluster.Nodes)

	if currentNodes < desiredNodes {
		health = key.Yellow
	}

	creating := cluster.HasCreatingCondition()
	updating := cluster.HasUpdatingCondition()
	deleted := cluster.HasDeletedCondition()
	deleting := cluster.HasDeletingCondition()

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
		Name:                   node.Name,
		Ready:                  ready,
		IP:                     ipFromAddresses(node.Status.Addresses),
		Hostname:               hostnameFromAddresses(node.Status.Addresses),
		InstanceType:           node.GetLabels()["beta.kubernetes.io/instance-type"],
		AvailabilityZone:       node.GetLabels()["failure-domain.beta.kubernetes.io/zone"],
		AvailabilityRegion:     node.GetLabels()["failure-domain.beta.kubernetes.io/region"],
		KubeletVersion:         node.Status.NodeInfo.KubeletVersion,
		CPUCount:               node.Status.Capacity.Cpu().Value(),
		MemoryCapacityBytes:    nodeMemoryToInt(node.Status.Capacity.Memory()),
		MemoryAllocatableBytes: nodeMemoryToInt(node.Status.Allocatable.Memory()),
		EphemeralStorageCap:    nodeMemoryToInt(node.Status.Capacity.StorageEphemeral()),
		EphemeralStorageAvail:  nodeMemoryToInt(node.Status.Allocatable.StorageEphemeral()),
	}
}

// FillNodeVersions takes node version information available at the cluster level and stores it in the associated node-level status
func FillNodeVersions(nodes []NodeStatus, versions []v1alpha1.StatusClusterNode) []NodeStatus {
	for i, node := range nodes {
		nodes[i].OperatorVersion = findVersionForNode(node.Name, versions)
	}
	return nodes
}

func findVersionForNode(name string, nodes []v1alpha1.StatusClusterNode) string {
	for _, node := range nodes {
		if node.Name == name {
			return node.Version
		}
	}
	return ""
}

func NewNodesStatus(nodes []v1.Node) []NodeStatus {
	result := []NodeStatus{}
	for _, node := range nodes {
		result = append(result, NewNodeStatus(node))
	}
	return result
}

// findInAddresses searches a list of NodeAddresses for an address of a given type and returns its value.
func findInAddresses(addresses []v1.NodeAddress, addressType v1.NodeAddressType) string {
	for _, entry := range addresses {
		if entry.Type == addressType {
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

func nodeMemoryToInt(nodeMemory *resource.Quantity) int64 {
	memBytes, ok := nodeMemory.AsInt64()
	if !ok {
		return 0
	}
	return memBytes
}
