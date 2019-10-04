package health

import (
	"strings"

	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/giantswarm/health-service/service/health/key"
)

func NewClusterStatus(cluster v1alpha1.StatusCluster, nodes []v1.Node) ClusterStatus {
	health := key.Default
	roleHealthCounts := map[string]map[key.Health]int{}
	workerCount := 0
	creating := cluster.HasCreatingCondition()
	updating := cluster.HasUpdatingCondition()
	deleting := cluster.HasDeletingCondition()

	for _, node := range nodes {
		role := nodeRole(node.Labels)
		nodeHealth := calculateNodeHealth(node)
		roleHealthCounts[role][nodeHealth]++
		if role == "worker" {
			workerCount++
		} else if role == "master" && nodeHealth != key.Green {
			if updating {
				health = key.Yellow
			} else {
				health = key.Red
			}
		}
	}

	if health != key.Red {
		if creating || deleting {
			health = key.Yellow
		}
		healthyWorkersProportion := float32(roleHealthCounts["worker"][key.Green]) / float32(workerCount)
		if healthyWorkersProportion < 0.75 {
			health = key.Yellow
		}
		if workerCount >= 20 && healthyWorkersProportion < 0.5 {
			health = key.Red
		}
		if roleHealthCounts["worker"][key.Green] < 3 {
			health = key.Red
		}
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
		Health: health,
		State:  state,
	}
}

func NewNodeStatus(node v1.Node) NodeStatus {
	return NodeStatus{
		Name:         node.Name,
		Ready:        nodeHasCondition(node.Status.Conditions, v1.ConditionTrue, v1.NodeReady),
		Health:       calculateNodeHealth(node),
		IP:           ipFromAddresses(node.Status.Addresses),
		Hostname:     hostnameFromAddresses(node.Status.Addresses),
		InstanceType: node.GetObjectMeta().GetLabels()["beta.kubernetes.io/instance-type"],
		CPUCount:     node.Status.Capacity.Cpu().Value(),
		MemoryBytes:  nodeMemoryToInt(node.Status.Capacity.Memory()),
		Role:         nodeRole(node.Labels),
	}
}

func NewNodesStatus(nodes []v1.Node) []NodeStatus {
	result := []NodeStatus{}
	for _, node := range nodes {
		result = append(result, NewNodeStatus(node))
	}
	return result
}

func nodeRole(labels map[string]string) string {
	roleKey := "kubernetes.io/role"
	roleKeyAlt := "node-role.kubernetes.io/"
	for key, value := range labels {
		if key == roleKey {
			return value
		} else if strings.HasPrefix(key, roleKeyAlt) && value == "" {
			runes := []rune(key)
			role := string(runes[len(roleKeyAlt):])
			return role
		}
	}
	return ""
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

func nodeHasCondition(conditions []v1.NodeCondition, s v1.ConditionStatus, t v1.NodeConditionType) bool {
	for _, c := range conditions {
		if c.Status == s && c.Type == t {
			return true
		}
	}

	return false
}

func calculateNodeHealth(node v1.Node) key.Health {
	ready := nodeHasCondition(node.Status.Conditions, v1.ConditionTrue, v1.NodeReady)
	if ready {
		return key.Green
	}
	role := nodeRole(node.ObjectMeta.Labels)
	if role == "master" {
		return key.Red
	}
	return key.Yellow
}
