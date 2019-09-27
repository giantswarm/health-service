package searcher

import (
	"github.com/giantswarm/microerror"
)

var invalidConfigError = &microerror.Error{
	Kind: "invalidConfigError",
}

var badRequestError = &microerror.Error{
	Kind: "badRequestError",
}

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}

// IsBadRequest asserts badRequestError.
func IsBadRequest(err error) bool {
	return microerror.Cause(err) == badRequestError
}
