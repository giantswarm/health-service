package searcher

import "github.com/giantswarm/microerror"

var clusterNotFoundError = &microerror.Error{
	Kind: "clusterNotFoundError",
}

// IsClusterNotFound asserts clusterNotFoundError.
func IsClusterNotFound(err error) bool {
	return microerror.Cause(err) == clusterNotFoundError
}

var invalidConfigError = &microerror.Error{
	Kind: "invalidConfigError",
}

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}
