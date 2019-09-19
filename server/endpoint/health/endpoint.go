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

// DefaultConfig provides a default configuration to create a new info
// endpoint by best effort.
func DefaultConfig() Config {
	return Config{
		// Dependencies.
		Logger:     nil,
		Middleware: nil,
		Service:    nil,
	}
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
		searcherConfig := searcher.DefaultConfig()
		searcherConfig.Logger = config.Logger
		searcherConfig.Middleware = config.Middleware
		searcherConfig.Service = config.Service
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
