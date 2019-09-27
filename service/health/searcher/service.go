package searcher

import (
	"context"

	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/health-service/service/health/key"
)

const clusterNamespace = "default"

// Config represents the configuration used to create a new service object.
type Config struct {
	G8sClient versioned.Interface
	Logger    micrologger.Logger

	Provider string
}

type Service struct {
	g8sClient versioned.Interface
	logger    micrologger.Logger

	provider string
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	if config.G8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.G8sClient must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	if config.Provider == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.Provider must not be empty", config)
	}
	if !key.IsKnownProvider(config.Provider) {
		return nil, microerror.Maskf(invalidConfigError, "%T.Provider must be one of %+v", config, key.Providers)
	}

	s := &Service{
		g8sClient: config.G8sClient,
		logger:    config.Logger,

		provider: config.Provider,
	}

	return s, nil
}

// searchAWSCR searches for the cluster config in AWSClusterConfigs resources.
func (s *Service) searchAWSInfo(ctx context.Context, clusterID string) (clusterInfo, error) {
	cr, err := s.g8sClient.ProviderV1alpha1().AWSConfigs(clusterNamespace).Get(clusterID, v1.GetOptions{})
	if errors.IsNotFound(err) {
		return clusterInfo{}, microerror.Mask(clusterNotFoundError)
	} else if err != nil {
		return clusterInfo{}, microerror.Mask(err)
	}
	cluster := clusterInfo{
		status: cr.Status.Cluster,
	}
	return cluster, nil
}

// searchAzureCR searches for the cluster config in AWSClusterConfigs resources.
func (s *Service) searchAzureInfo(ctx context.Context, clusterID string) (clusterInfo, error) {
	cr, err := s.g8sClient.ProviderV1alpha1().AzureConfigs(clusterNamespace).Get(clusterID, v1.GetOptions{})
	if errors.IsNotFound(err) {
		return clusterInfo{}, microerror.Mask(clusterNotFoundError)
	} else if err != nil {
		return clusterInfo{}, microerror.Mask(err)
	}
	cluster := clusterInfo{
		status: cr.Status.Cluster,
	}
	return cluster, nil
}

// searchKVMCR abc
func (s *Service) searchKVMInfo(ctx context.Context, clusterID string) (clusterInfo, error) {
	cr, err := s.g8sClient.ProviderV1alpha1().KVMConfigs(clusterNamespace).Get(clusterID, v1.GetOptions{})
	if errors.IsNotFound(err) {
		return clusterInfo{}, microerror.Mask(clusterNotFoundError)
	} else if err != nil {
		return clusterInfo{}, microerror.Mask(err)
	}
	cluster := clusterInfo{
		status: cr.Status.Cluster,
	}
	return cluster, nil
}

func evaluateClusterStatus(cluster clusterInfo) ClusterStatus {
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

// Search searches for the cluster information.
// It try to find cluster information in CR and fallback to storage service when nothing is found.
func (s *Service) Search(ctx context.Context, request Request) (*Response, error) {
	var err error
	var cluster clusterInfo

	switch s.provider {
	case "aws":
		cluster, err = s.searchAWSInfo(ctx, request.ClusterID)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	case "azure":
		cluster, err = s.searchAzureInfo(ctx, request.ClusterID)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	case "kvm":
		cluster, err = s.searchKVMInfo(ctx, request.ClusterID)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	response := Response{
		Cluster: evaluateClusterStatus(cluster),
	}

	return &response, nil
}
