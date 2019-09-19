package searcher

import (
	"context"

	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

// Config represents the configuration used to create a new service object.
type Config struct {
	// Dependencies.
	G8sClient versioned.Interface
	Logger    micrologger.Logger

	Provider string
}

// DefaultConfig provides a default configuration to create a new service object
// by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		Logger: nil,

		// Settings
		Provider: "",
	}

	return newConfig
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	// Dependencies.
	if config.G8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "g8sclient must not be empty")
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "logger must not be empty")
	}
	if config.Provider == "" {
		return nil, microerror.Maskf(invalidConfigError, "provider must not be empty")
	}

	newService := &Service{
		Config: config,
	}

	return newService, nil
}

type Service struct {
	Config
}

// Search searches for the cluster information.
// It try to find cluster information in CR and fallback to storage service when nothing is found.
func (s *Service) Search(ctx context.Context, request Request) (*Response, error) {
	response, err := s.searchCR(ctx, request)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return response, nil
}
