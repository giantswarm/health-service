package searcher

import (
	"context"
)

// searchAWSCR searches for the cluster config in AWSClusterConfigs resources.
func (s *Service) searchAWSCR(ctx context.Context, request Request) (*Response, error) {
	response := Response{
		ClusterHealth: "green",
	}

	return &response, nil
}
