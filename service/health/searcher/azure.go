package searcher

import (
	"context"
)

// searchAzureCR searches for the cluster config in AWSClusterConfigs resources.
func (s *Service) searchAzureCR(ctx context.Context, request Request) (*Response, error) {
	response := Response{
		ClusterHealth: "green",
	}

	return &response, nil
}
