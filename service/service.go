// Package service implements business logic to create Kubernetes resources
// against the Kubernetes API.
package service

import (
	"sync"
	"time"

	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/microendpoint/service/version"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/microstorage"
	"github.com/giantswarm/tenantcluster"

	"github.com/giantswarm/health-service/service/health"
)

const (
	// DefaultRetryCount is the number of times to retry a failed network call.
	DefaultRetryCount = 5
	// DefaultTimeout is the timeout for network calls.
	DefaultTimeout = 5 * time.Second
)

// Config represents the configuration used to create a new service.
type Config struct {
	G8sClient     versioned.Interface
	Logger        micrologger.Logger
	Storage       microstorage.Storage
	TenantCluster tenantcluster.Interface
	Provider      string

	Description string
	GitCommit   string
	Name        string
	Source      string
	Version     string
}

type Service struct {
	Health  *health.Service
	Version *version.Service

	bootOnce sync.Once
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	if config.G8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.G8sClient must not be empty", config)
	}
	if config.Provider == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.Provider must not be empty", config)
	}
	if config.TenantCluster == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.TenantCluster must not be empty", config)
	}

	var err error
	var healthService *health.Service
	{
		healthConfig := health.Config{
			G8sClient:     config.G8sClient,
			Logger:        config.Logger,
			TenantCluster: config.TenantCluster,
			Provider:      config.Provider,
		}

		healthService, err = health.New(healthConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var versionService *version.Service
	{
		versionConfig := version.Config{
			Description: config.Description,
			GitCommit:   config.GitCommit,
			Name:        config.Name,
			Source:      config.Source,
			Version:     config.Version,
		}

		versionService, err = version.New(versionConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	s := &Service{
		Health:  healthService,
		Version: versionService,

		bootOnce: sync.Once{},
	}

	return s, nil
}
