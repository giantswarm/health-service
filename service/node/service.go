// Package node provides cluster specific business logic.
package node

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/tenantcluster"

	"github.com/giantswarm/health-service/service/node/searcher"
)

// Config represents the configuration used to create a new service.
type Config struct {
	Logger        micrologger.Logger
	TenantCluster tenantcluster.Interface
}

type Service struct {
	Searcher *searcher.Service
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	if config.TenantCluster == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.TenantCluster must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var searcherService *searcher.Service
	{
		searcherConfig := searcher.Config{
			Logger:        config.Logger,
			TenantCluster: config.TenantCluster,
		}
		searcherService, err = searcher.New(searcherConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	newService := &Service{
		Searcher: searcherService,
	}

	return newService, nil
}
