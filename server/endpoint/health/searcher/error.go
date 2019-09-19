package searcher

import (
	"github.com/giantswarm/microerror"
)

var invalidConfigError = &microerror.Error{
	Kind: "invalidConfigError",
}

var clusterNotFoundError = &microerror.Error{
	Kind: "clusterNotFoundError",
}

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}

// IsClusterNotFound asserts clusterNotFoundError.
func IsClusterNotFound(err error) bool {
	return microerror.Cause(err) == clusterNotFoundError
}
