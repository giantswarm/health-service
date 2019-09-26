package searcher

import (
	"context"

	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
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
func (s *Service) searchAWSInfo(ctx context.Context, clusterID string) (v1alpha1.StatusCluster, error) {
	cr, err := s.g8sClient.ProviderV1alpha1().AWSConfigs(clusterNamespace).Get(clusterID, v1.GetOptions{})
	if errors.IsNotFound(err) {
		return v1alpha1.StatusCluster{}, microerror.Mask(clusterNotFoundError)
	} else if err != nil {
		return v1alpha1.StatusCluster{}, microerror.Mask(err)
	}
	return cr.Status.Cluster, nil
}

// searchAzureCR searches for the cluster config in AWSClusterConfigs resources.
func (s *Service) searchAzureInfo(ctx context.Context, clusterID string) (v1alpha1.StatusCluster, error) {
	cr, err := s.g8sClient.ProviderV1alpha1().AzureConfigs(clusterNamespace).Get(clusterID, v1.GetOptions{})
	if errors.IsNotFound(err) {
		return v1alpha1.StatusCluster{}, microerror.Mask(clusterNotFoundError)
	} else if err != nil {
		return v1alpha1.StatusCluster{}, microerror.Mask(err)
	}
	return cr.Status.Cluster, nil
}

// searchKVMCR abc
func (s *Service) searchKVMInfo(ctx context.Context, clusterID string) (v1alpha1.StatusCluster, error) {
	cr, err := s.g8sClient.ProviderV1alpha1().KVMConfigs(clusterNamespace).Get(clusterID, v1.GetOptions{})
	if errors.IsNotFound(err) {
		return v1alpha1.StatusCluster{}, microerror.Mask(clusterNotFoundError)
	} else if err != nil {
		return v1alpha1.StatusCluster{}, microerror.Mask(err)
	}
	return cr.Status.Cluster, nil
}

func evaluateGeneralHealth(status v1alpha1.StatusCluster) GeneralStatus {
	clusterHealth := key.HealthGreen // Optimistic default

	desiredNodes := status.Scaling.DesiredCapacity
	currentNodes := len(status.Nodes)

	if currentNodes < desiredNodes {
		clusterHealth = key.HealthYellow
	}

	creating := status.HasCreatingCondition()
	updating := status.HasUpdatingCondition()
	deleted := status.HasDeletedCondition()
	deleting := status.HasDeletingCondition()

	if deleted || deleting {
		clusterHealth = key.HealthRed
	} else if updating || creating { // TODO: Account for draining/scaling
		clusterHealth = key.HealthYellow
	}

	return GeneralStatus{
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
	var status v1alpha1.StatusCluster

	switch s.provider {
	case "aws":
		status, err = s.searchAWSInfo(ctx, request.ClusterID)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	case "azure":
		status, err = s.searchAzureInfo(ctx, request.ClusterID)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	case "kvm":
		status, err = s.searchKVMInfo(ctx, request.ClusterID)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	response := Response{
		General: evaluateGeneralHealth(status),
	}

	return &response, nil
}
