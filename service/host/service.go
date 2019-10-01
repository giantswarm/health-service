package host

import (
	"context"

	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/health-service/pkg/errors"
	"github.com/giantswarm/health-service/service/host/key"
)

const clusterNamespace = "default"

// Config represents the configuration used to create a new service object.
type Config struct {
	G8sClient versioned.Interface
	Logger    micrologger.Logger

	Provider string
}

// Service is an object representing the health searcher service.
type Service struct {
	g8sClient versioned.Interface
	logger    micrologger.Logger

	provider string
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	if config.G8sClient == nil {
		return nil, microerror.Maskf(errors.InvalidConfigError, "%T.G8sClient must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(errors.InvalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Provider == "" {
		return nil, microerror.Maskf(errors.InvalidConfigError, "%T.Provider must not be empty", config)
	}
	if !key.IsKnownProvider(config.Provider) {
		return nil, microerror.Maskf(errors.InvalidConfigError, "%T.Provider must be one of %+v", config, key.Providers)
	}

	s := &Service{
		g8sClient: config.G8sClient,
		logger:    config.Logger,

		provider: config.Provider,
	}

	return s, nil
}

// searchAWSInfo searches for the cluster status in AWSConfigs resources.
func (s *Service) searchAWSStatus(ctx context.Context, clusterID string) (Response, error) {
	cr, err := s.g8sClient.ProviderV1alpha1().AWSConfigs(clusterNamespace).Get(clusterID, v1.GetOptions{})
	if k8serrors.IsNotFound(err) {
		return Response{}, microerror.Mask(errors.NotFoundError)
	} else if err != nil {
		return Response{}, microerror.Mask(err)
	}
	cluster := Response{
		Endpoint: cr.Spec.Cluster.Kubernetes.API.Domain,
		Status:   cr.Status.Cluster,
	}
	return cluster, nil
}

// searchAzureInfo searches for the cluster status in AzureConfig resources.
func (s *Service) searchAzureStatus(ctx context.Context, clusterID string) (Response, error) {
	cr, err := s.g8sClient.ProviderV1alpha1().AzureConfigs(clusterNamespace).Get(clusterID, v1.GetOptions{})
	if k8serrors.IsNotFound(err) {
		return Response{}, microerror.Mask(errors.NotFoundError)
	} else if err != nil {
		return Response{}, microerror.Mask(err)
	}
	cluster := Response{
		Endpoint: cr.Spec.Cluster.Kubernetes.API.Domain,
		Status:   cr.Status.Cluster,
	}
	return cluster, nil
}

// searchKVMInfo searches for the cluster status in KVMConfig resources.
func (s *Service) searchKVMStatus(ctx context.Context, clusterID string) (Response, error) {
	cr, err := s.g8sClient.ProviderV1alpha1().KVMConfigs(clusterNamespace).Get(clusterID, v1.GetOptions{})
	if k8serrors.IsNotFound(err) {
		return Response{}, microerror.Mask(errors.NotFoundError)
	} else if err != nil {
		return Response{}, microerror.Mask(err)
	}
	cluster := Response{
		Endpoint: cr.Spec.Cluster.Kubernetes.API.Domain,
		Status:   cr.Status.Cluster,
	}
	return cluster, nil
}

// SearchStatusCluster searches for the StatusCluster of a provider-specific cluster config CR.
func (s *Service) SearchStatusCluster(ctx context.Context, request Request) (*Response, error) {
	var err error
	var response Response

	switch s.provider {
	case "aws":
		response, err = s.searchAWSStatus(ctx, request.ClusterID)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	case "azure":
		response, err = s.searchAzureStatus(ctx, request.ClusterID)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	case "kvm":
		response, err = s.searchKVMStatus(ctx, request.ClusterID)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	return &response, nil
}
