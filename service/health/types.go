package health

import (
	"regexp"
	"strings"

	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/giantswarm/health-service/service/health/key"
)

const (
	attachableVolumePattern = v1.ResourceAttachableVolumesPrefix + "*"
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
		if roleHealthCounts[role] == nil {
			roleHealthCounts[role] = map[key.Health]int{}
		}
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

func NewNodeStatus(node v1.Node, pods []v1.Pod) NodeStatus {
	limits := NodeStatusComputeResources{}
	requests := NodeStatusComputeResources{}
	for _, pod := range pods {
		if pod.Spec.NodeName == node.Name {
			for _, container := range pod.Spec.Containers {
				limits.CPU += container.Resources.Limits.Cpu().MilliValue()
				limits.MemoryBytes += container.Resources.Limits.Memory().Value()
				requests.CPU += container.Resources.Requests.Cpu().MilliValue()
				requests.MemoryBytes += container.Resources.Requests.Memory().Value()
			}
		}
	}
	return NodeStatus{
		Name:                              node.Name,
		Role:                              nodeRole(node.Labels),
		Health:                            calculateNodeHealth(node),
		Ready:                             nodeHasCondition(node.Status.Conditions, v1.ConditionTrue, v1.NodeReady),
		IP:                                ipFromAddresses(node.Status.Addresses),
		Hostname:                          hostnameFromAddresses(node.Status.Addresses),
		InstanceType:                      node.GetLabels()["beta.kubernetes.io/instance-type"],
		AvailabilityZone:                  node.GetLabels()["failure-domain.beta.kubernetes.io/zone"],
		AvailabilityRegion:                node.GetLabels()["failure-domain.beta.kubernetes.io/region"],
		KubeletVersion:                    node.Status.NodeInfo.KubeletVersion,
		CPUCount:                          node.Status.Capacity.Cpu().Value(),
		MemoryCapacityBytes:               nodeMemoryToInt(node.Status.Capacity.Memory()),
		MemoryAllocatableBytes:            nodeMemoryToInt(node.Status.Allocatable.Memory()),
		EphemeralStorageCap:               nodeMemoryToInt(node.Status.Capacity.StorageEphemeral()),
		EphemeralStorageAvail:             nodeMemoryToInt(node.Status.Allocatable.StorageEphemeral()),
		AttachableVolumesAllocatableCount: countAttachableVolumes(node.Status.Allocatable),
		AttachableVolumesCapacityCount:    countAttachableVolumes(node.Status.Capacity),
		LimitTotals:                       limits,
		RequestTotals:                     requests,
	}
}

// Counts the total number of Attachable volumes in the ResourceList
func countAttachableVolumes(nodeStatus v1.ResourceList) int64 {
	sum := int64(0)
	for _, count := range getAttachableVolumes(nodeStatus) {
		sum += count
	}
	return sum
}

// Searches a node's resources for attachable volumes,
// and returns a mapping of the resource types to the number of that type
func getAttachableVolumes(nodeStatus v1.ResourceList) map[string]int64 {
	attachableVolumeRegex, _ := regexp.Compile(attachableVolumePattern)
	var result = map[string]int64{}

	for i, resource := range nodeStatus {
		if attachableVolumeRegex.MatchString(i.String()) {
			result[i.String()] = nodeMemoryToInt(&resource)
		}
	}
	return result
}

// FillNodeVersions takes node version information available at the cluster level and stores it in the associated node-level status
func FillNodeVersions(nodes []NodeStatus, versions []v1alpha1.StatusClusterNode) []NodeStatus {
	for i, node := range nodes {
		nodes[i].OperatorVersion = findVersionForNode(node.Name, versions)
	}
	return nodes
}

// findVersionForNode searches a list of cluster status nodes for a node with the given name, and returns the node's version
func findVersionForNode(name string, nodes []v1alpha1.StatusClusterNode) string {
	for _, node := range nodes {
		if node.Name == name {
			return node.Version
		}
	}
	return ""
}

func NewNodesStatus(nodes []v1.Node, pods []v1.Pod) []NodeStatus {
	result := []NodeStatus{}
	for _, node := range nodes {
		result = append(result, NewNodeStatus(node, pods))
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

// Accepts a node memory (or any resource Quantity) and
// returns it as an int, if it can be represented as one, else 0.
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
