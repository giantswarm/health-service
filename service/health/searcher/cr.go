package searcher

import (
	"context"

	"github.com/giantswarm/microerror"
)

// searchCR searches for the cluster config in CR resources.
func (s *Service) searchCR(ctx context.Context, request Request) (response *Response, err error) {
	switch s.provider {
	case "aws":
		response, err = s.searchAWSCR(ctx, request)
	case "azure":
		response, err = s.searchAzureCR(ctx, request)
	case "kvm":
		response, err = s.searchKVMCR(ctx, request)
	default:
		return nil, microerror.Maskf(invalidConfigError, "unsupported provider: %s", s.provider)
	}

	if err != nil {
		return nil, microerror.Mask(err)
	}

	return response, nil
}
