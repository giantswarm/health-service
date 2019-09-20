package health

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/health-service/server/endpoint/health/searcher"
	"github.com/giantswarm/health-service/server/middleware"
	"github.com/giantswarm/health-service/service"
)

// Config represents the configuration used to create a info endpoint.
type Config struct {
	// Dependencies.
	Logger     micrologger.Logger
	Middleware *middleware.Middleware
	Service    *service.Service
}

// Endpoint is the info endpoint collection.
type Endpoint struct {
	Searcher *searcher.Endpoint
}

// New creates a new configured info endpoint.
func New(config Config) (*Endpoint, error) {
	var err error

	var searcherEndpoint *searcher.Endpoint
	{
		searcherConfig := searcher.Config{
			Logger:     config.Logger,
			Middleware: config.Middleware,
			Service:    config.Service,
		}
		searcherEndpoint, err = searcher.New(searcherConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	newEndpoint := &Endpoint{
		Searcher: searcherEndpoint,
	}

	return newEndpoint, nil
}
