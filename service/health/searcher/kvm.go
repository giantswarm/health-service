package searcher

import (
	"context"
)

// searchKVMCR searches for the cluster config in AWSClusterConfigs resources.
func (s *Service) searchKVMCR(ctx context.Context, request Request) (*Response, error) {
	response := Response{
		ClusterHealth: "green",
	}

	return &response, nil
}
