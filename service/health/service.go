package health

import (
	"context"

	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	v1 "k8s.io/api/core/v1"

	"github.com/giantswarm/health-service/service/health/key"
)

// Config represents the configuration used to create a new health service.
type Config struct {
	Logger micrologger.Logger
}

// Service is an object representing the health service.
type Service struct {
	logger micrologger.Logger
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	s := &Service{
		logger: config.Logger,
	}

	return s, nil
}

// Search accepts data from the control plane and tenant cluster and returns
// an object representing an aggregate health of the cluster and its nodes.
func (s *Service) Search(ctx context.Context, request Request) (*Response, error) {
	clusterStatus := calculateClusterStatus(request.Cluster, request.Nodes)
	nodeStatus := calculateNodeStatus(request.Cluster, request.Nodes)

	response := Response{
		Cluster: clusterStatus,
		Nodes:   nodeStatus,
	}

	return &response, nil
}

func calculateClusterStatus(cluster v1alpha1.StatusCluster, nodes []v1.Node) ClusterStatus {
	health := key.HealthDefault
	roleHealthCounts := map[string]map[key.Health]int{}
	workerCount := 0
	creating := cluster.HasCreatingCondition()
	updating := cluster.HasUpdatingCondition()
	deleting := cluster.HasDeletingCondition()

	for _, node := range nodes {
		role := key.NodeRole(node)
		nodeHealth := calculateNodeHealth(node)
		if roleHealthCounts[role] == nil {
			roleHealthCounts[role] = map[key.Health]int{}
		}
		roleHealthCounts[role][nodeHealth]++
		if role == "worker" {
			workerCount++
		} else if role == "master" && nodeHealth != key.HealthGreen {
			if updating {
				health = key.HealthYellow
			} else {
				health = key.HealthRed
			}
		}
	}

	if health != key.HealthRed {
		if creating || deleting {
			health = key.HealthYellow
		}
		healthyWorkersProportion := float32(roleHealthCounts["worker"][key.HealthGreen]) / float32(workerCount)
		if healthyWorkersProportion < 0.75 {
			health = key.HealthYellow
		}
		if workerCount >= 20 && healthyWorkersProportion < 0.5 {
			health = key.HealthRed
		}
		if roleHealthCounts["worker"][key.HealthGreen] < 3 {
			health = key.HealthRed
		}
	}

	state := key.StateNormal

	if creating {
		state = key.StateCreating
	} else if updating {
		state = key.StateUpgrading
	} else if deleting {
		state = key.StateDeleting
	}

	return ClusterStatus{
		Health: health,
		State:  state,
	}
}

func calculateNodeStatus(cluster v1alpha1.StatusCluster, nodes []v1.Node) []NodeStatus {
	result := []NodeStatus{}
	for _, node := range nodes {
		nodeStatus := NodeStatus{
			Health: calculateNodeHealth(node),
			Ready:  key.NodeHasCondition(node.Status.Conditions, v1.ConditionTrue, v1.NodeReady),
			Identity: NodeStatusIdentity{
				Name:               node.Name,
				Role:               key.NodeRole(node),
				IP:                 key.NodeIP(node.Status.Addresses),
				Hostname:           key.NodeHostname(node.Status.Addresses),
				InstanceType:       node.Labels["beta.kubernetes.io/instance-type"],
				AvailabilityZone:   node.Labels["failure-domain.beta.kubernetes.io/zone"],
				AvailabilityRegion: node.Labels["failure-domain.beta.kubernetes.io/region"],
				KubeletVersion:     node.Status.NodeInfo.KubeletVersion,
				OperatorVersion:    key.NodeVersion(cluster.Nodes, node.Name),
			},
			MachineResources: NodeStatusMachineResources{
				CPUCount:                          node.Status.Capacity.Cpu().Value(),
				MemoryCapacityBytes:               key.MemoryToInt(node.Status.Capacity.Memory()),
				MemoryAllocatableBytes:            key.MemoryToInt(node.Status.Allocatable.Memory()),
				EphemeralStorageCap:               key.MemoryToInt(node.Status.Capacity.StorageEphemeral()),
				EphemeralStorageAvail:             key.MemoryToInt(node.Status.Allocatable.StorageEphemeral()),
				AttachableVolumesAllocatableCount: key.NodeAttachableVolumesCount(node.Status.Allocatable),
				AttachableVolumesCapacityCount:    key.NodeAttachableVolumesCount(node.Status.Capacity),
			},
		}
		result = append(result, nodeStatus)
	}
	return result
}

func calculateNodeHealth(node v1.Node) key.Health {
	ready := key.NodeHasCondition(node.Status.Conditions, v1.ConditionTrue, v1.NodeReady)
	if ready {
		return key.HealthGreen
	}
	role := key.NodeRole(node)
	if role == "master" {
		return key.HealthRed
	}
	return key.HealthYellow
}
