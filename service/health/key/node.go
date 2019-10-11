package key

import (
	"regexp"
	"strings"

	"github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	attachableVolumePattern = v1.ResourceAttachableVolumesPrefix + "*"
	roleKey                 = "kubernetes.io/role"
	roleKeyAlt              = "node-role.kubernetes.io/"
)

// NodeVersion returns the version of a node from an array of StatusClusterNode.
func NodeVersion(nodes []v1alpha1.StatusClusterNode, name string) string {
	for _, node := range nodes {
		if node.Name == name {
			return node.Version
		}
	}
	return ""
}

// NodeRole returns the role of a node by parsing its labels.
func NodeRole(node v1.Node) string {
	for key, value := range node.Labels {
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

// nodeAddressValue searches a list of NodeAddresses for an address of a given type and returns its value.
func nodeAddressValue(addresses []v1.NodeAddress, addressType v1.NodeAddressType) string {
	for _, entry := range addresses {
		if entry.Type == addressType {
			return entry.Address
		}
	}
	return ""
}

// NodeIP returns the value of the IP address for a node.
func NodeIP(addresses []v1.NodeAddress) string {
	return nodeAddressValue(addresses, v1.NodeInternalIP)
}

// NodeHostname returns the value of the hostname address for a node.
func NodeHostname(addresses []v1.NodeAddress) string {
	return nodeAddressValue(addresses, v1.NodeHostName)
}

// MemoryToInt returns the memory resource as an integer or 0 if conversion fails.
func MemoryToInt(nodeMemory *resource.Quantity) int64 {
	memBytes, ok := nodeMemory.AsInt64()
	if !ok {
		return 0
	}
	return memBytes
}

// NodeHasCondition returns true if the node conditions contains a condition of the given type and value.
func NodeHasCondition(conditions []v1.NodeCondition, s v1.ConditionStatus, t v1.NodeConditionType) bool {
	for _, c := range conditions {
		if c.Status == s && c.Type == t {
			return true
		}
	}
	return false
}

// NodeAttachableVolumesCount counts the total number of attachable volumes in the ResourceList
func NodeAttachableVolumesCount(nodeStatus v1.ResourceList) int64 {
	sum := int64(0)
	for _, count := range nodeAttachableVolumes(nodeStatus) {
		sum += count
	}
	return sum
}

// Searches a node's resources for attachable volumes,
// and returns a mapping of the resource types to the number of that type
func nodeAttachableVolumes(nodeStatus v1.ResourceList) map[string]int64 {
	attachableVolumeRegex, _ := regexp.Compile(attachableVolumePattern)
	var result = map[string]int64{}

	for i, resource := range nodeStatus {
		if attachableVolumeRegex.MatchString(i.String()) {
			result[i.String()] = MemoryToInt(&resource)
		}
	}
	return result
}
