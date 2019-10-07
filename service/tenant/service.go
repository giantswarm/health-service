package tenant

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/tenantcluster"

	"github.com/giantswarm/health-service/service/tenant/nodes"
	"github.com/giantswarm/health-service/service/tenant/pods"
)

// Config represents the configuration used to create a new service.
type Config struct {
	Logger        micrologger.Logger
	TenantCluster tenantcluster.Interface
}

type Service struct {
	Nodes *nodes.Service
	Pods  *pods.Service
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

	var nodesService *nodes.Service
	{
		nodesConfig := nodes.Config{
			Logger:        config.Logger,
			TenantCluster: config.TenantCluster,
		}

		nodesService, err = nodes.New(nodesConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var podsService *pods.Service
	{
		podsConfig := pods.Config{
			Logger:        config.Logger,
			TenantCluster: config.TenantCluster,
		}

		podsService, err = pods.New(podsConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	s := &Service{
		Nodes: nodesService,
		Pods:  podsService,
	}

	return s, nil
}
