package endpoint

import (
	"github.com/giantswarm/microendpoint/endpoint/healthz"
	"github.com/giantswarm/microendpoint/endpoint/version"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/health-service/server/endpoint/health"
	"github.com/giantswarm/health-service/server/middleware"
	"github.com/giantswarm/health-service/service"
)

// Config represents the configuration used to create a endpoint.
type Config struct {
	// Dependencies.
	Logger     micrologger.Logger
	Middleware *middleware.Middleware
	Service    *service.Service
}

type Endpoint struct {
	Health  *health.Endpoint
	Healthz *healthz.Endpoint
	Version *version.Endpoint
}

func New(config Config) (*Endpoint, error) {
	var err error

	var healthEndpoint *health.Endpoint
	{
		healthConfig := health.Config{
			Logger:     config.Logger,
			Middleware: config.Middleware,
			Service:    config.Service,
		}
		healthEndpoint, err = health.New(healthConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var healthzEndpoint *healthz.Endpoint
	{
		c := healthz.Config{
			Logger: config.Logger,
		}

		healthzEndpoint, err = healthz.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var versionEndpoint *version.Endpoint
	{
		versionConfig := version.Config{
			Logger:  config.Logger,
			Service: config.Service.Version,
		}

		versionEndpoint, err = version.New(versionConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	e := &Endpoint{
		Health:  healthEndpoint,
		Healthz: healthzEndpoint,
		Version: versionEndpoint,
	}

	return e, nil
}
