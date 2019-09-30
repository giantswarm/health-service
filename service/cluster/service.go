// Package cluster provides cluster specific business logic.
package cluster

import (
	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/health-service/service/cluster/searcher"
)

// Config represents the configuration used to create a new service.
type Config struct {
	G8sClient versioned.Interface
	Logger    micrologger.Logger
	Provider  string
}

type Service struct {
	Searcher *searcher.Service
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	if config.G8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.G8sClient must not be empty", config)
	}
	if config.Provider == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.Provider must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var searcherService *searcher.Service
	{
		searcherConfig := searcher.Config{
			G8sClient: config.G8sClient,
			Logger:    config.Logger,
			Provider:  config.Provider,
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
