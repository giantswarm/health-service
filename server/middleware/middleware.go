package middleware

import (
	"github.com/giantswarm/health-service/service"
	"github.com/giantswarm/micrologger"
)

// Config represents the configuration used to create a middleware.
type Config struct {
	// Dependencies.
	Logger  micrologger.Logger
	Service *service.Service
}

// New creates a new configured middleware.
func New(config Config) (*Middleware, error) {
	return &Middleware{}, nil
}

// Middleware is middleware collection.
type Middleware struct {
}
