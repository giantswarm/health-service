package searcher

import (
	"context"

	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/health-service/service/cluster/key"
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

// searchAWSInfo searches for the cluster status in AWSConfigs resources.
func (s *Service) searchAWSInfo(ctx context.Context, clusterID string) (Response, error) {
	cr, err := s.g8sClient.ProviderV1alpha1().AWSConfigs(clusterNamespace).Get(clusterID, v1.GetOptions{})
	if errors.IsNotFound(err) {
		return Response{}, microerror.Mask(clusterNotFoundError)
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
func (s *Service) searchAzureInfo(ctx context.Context, clusterID string) (Response, error) {
	cr, err := s.g8sClient.ProviderV1alpha1().AzureConfigs(clusterNamespace).Get(clusterID, v1.GetOptions{})
	if errors.IsNotFound(err) {
		return Response{}, microerror.Mask(clusterNotFoundError)
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
func (s *Service) searchKVMInfo(ctx context.Context, clusterID string) (Response, error) {
	cr, err := s.g8sClient.ProviderV1alpha1().KVMConfigs(clusterNamespace).Get(clusterID, v1.GetOptions{})
	if errors.IsNotFound(err) {
		return Response{}, microerror.Mask(clusterNotFoundError)
	} else if err != nil {
		return Response{}, microerror.Mask(err)
	}
	cluster := Response{
		Endpoint: cr.Spec.Cluster.Kubernetes.API.Domain,
		Status:   cr.Status.Cluster,
	}
	return cluster, nil
}

// Search searches for the cluster information.
// It try to find cluster information in CR and fallback to storage service when nothing is found.
func (s *Service) Search(ctx context.Context, request Request) (*Response, error) {
	var err error
	var response Response

	switch s.provider {
	case "aws":
		response, err = s.searchAWSInfo(ctx, request.ClusterID)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	case "azure":
		response, err = s.searchAzureInfo(ctx, request.ClusterID)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	case "kvm":
		response, err = s.searchKVMInfo(ctx, request.ClusterID)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	return &response, nil
}
