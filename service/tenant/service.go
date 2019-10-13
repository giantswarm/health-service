package tenant

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/tenantcluster"

	"github.com/giantswarm/health-service/service/tenant/node"
	"github.com/giantswarm/health-service/service/tenant/pod"
)

// Config represents the configuration used to create a new service.
type Config struct {
	Logger        micrologger.Logger
	TenantCluster tenantcluster.Interface
}

type Service struct {
	Node *node.Service
	Pod  *pod.Service
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.TenantCluster == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.TenantCluster must not be empty", config)
	}

	var err error

	var nodeService *node.Service
	{
		nodeConfig := node.Config{
			Logger:        config.Logger,
			TenantCluster: config.TenantCluster,
		}

		nodeService, err = node.New(nodeConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var podService *pod.Service
	{
		podConfig := pod.Config{
			Logger:        config.Logger,
			TenantCluster: config.TenantCluster,
		}

		podService, err = pod.New(podConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	s := &Service{
		Node: nodeService,
		Pod:  podService,
	}

	return s, nil
}
