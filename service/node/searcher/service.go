package searcher

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/tenantcluster"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const clusterNamespace = "default"

// Config represents the configuration used to create a new service object.
type Config struct {
	Logger        micrologger.Logger
	TenantCluster tenantcluster.Interface
}

// Service is an object representing the health searcher service.
type Service struct {
	logger        micrologger.Logger
	tenantCluster tenantcluster.Interface
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.TenantCluster == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.TenantCluster must not be empty", config)
	}

	s := &Service{
		logger:        config.Logger,
		tenantCluster: config.TenantCluster,
	}

	return s, nil
}

// Search searches a tenant cluster for information about nodes.
func (s *Service) Search(ctx context.Context, request Request) (*Response, error) {
	var err error
	var response Response

	var k8sClient kubernetes.Interface
	{
		// TODO: cache this?
		k8sClient, err = s.tenantCluster.NewK8sClient(ctx, request.ClusterID, request.Endpoint)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	{
		nodes, err := k8sClient.CoreV1().Nodes().List(v1.ListOptions{})
		if err != nil {
			return nil, microerror.Mask(err)
		}
		response.Nodes = nodes.Items
	}

	return &response, nil
}
