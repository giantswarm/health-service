package health

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

// Config represents the configuration used to create a new service object.
type Config struct {
	Logger micrologger.Logger
}

// Service is an object representing the health searcher service.
type Service struct {
	logger micrologger.Logger
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	s := &Service{
		logger: config.Logger,
	}

	return s, nil
}

// Search searches for the cluster information.
// It try to find cluster information in CR and fallback to storage service when nothing is found.
func (s *Service) Search(ctx context.Context, request Request) (*Response, error) {
	clusterHealth := NewClusterStatus(request.Cluster, request.Nodes)
	nodesHealth := NewNodesStatus(request.Nodes, request.Pods)
	nodesHealth = FillNodeVersions(nodesHealth, request.Cluster.Nodes)

	response := Response{
		Cluster: clusterHealth,
		Nodes:   nodesHealth,
	}

	return &response, nil
}
