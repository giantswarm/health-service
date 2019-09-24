package searcher

import (
	"context"

	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/health-service/service/health/key"
)

// Config represents the configuration used to create a new service object.
type Config struct {
	G8sClient versioned.Interface
	Logger    micrologger.Logger

	Provider string
}

type Service struct {
	g8sClient versioned.Interface
	logger    micrologger.Logger

	provider string
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	if config.G8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.G8sClient must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	if config.Provider == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.Provider must not be empty", config)
	}
	if !key.IsKnownProvider(config.Provider) {
		return nil, microerror.Maskf(invalidConfigError, "provider must be one of %+v", key.Providers)
	}

	s := &Service{
		g8sClient: config.G8sClient,
		logger:    config.Logger,

		provider: config.Provider,
	}

	return s, nil
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
