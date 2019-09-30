package searcher

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	kitendpoint "github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/giantswarm/health-service/server/middleware"
	service "github.com/giantswarm/health-service/service"
	cluster "github.com/giantswarm/health-service/service/cluster/searcher"
	health "github.com/giantswarm/health-service/service/health/searcher"
	node "github.com/giantswarm/health-service/service/node/searcher"
)

const (
	// Method is the HTTP method this endpoint is registered for.
	Method = "GET"
	// Name identifies the endpoint. It is aligned to the package path.
	Name = "health"
	// Path is the HTTP request path this endpoint is registered for.
	Path = "/health/{id}"
)

type Config struct {
	Logger     micrologger.Logger
	Middleware *middleware.Middleware
	Service    *service.Service
}

type Endpoint struct {
	logger     micrologger.Logger
	middleware *middleware.Middleware
	service    *service.Service
}

func New(config Config) (*Endpoint, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "config.Logger must not be empty")
	}
	if config.Middleware == nil {
		return nil, microerror.Maskf(invalidConfigError, "config.Middleware must not be empty")
	}
	if config.Service == nil {
		return nil, microerror.Maskf(invalidConfigError, "config.Service must not be empty")
	}

	e := &Endpoint{
		logger:     config.Logger,
		middleware: config.Middleware,
		service:    config.Service,
	}

	return e, nil
}

func (e *Endpoint) Decoder() kithttp.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		vars := mux.Vars(r)
		id := vars["id"]

		return id, nil
	}
}

func (e *Endpoint) Encoder() kithttp.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		return json.NewEncoder(w).Encode(response)
	}
}

func (e *Endpoint) Endpoint() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		clusterID, ok := request.(string)
		if !ok {
			return nil, microerror.Mask(badRequestError)
		}

		clusterRequest := cluster.Request{
			ClusterID: clusterID,
		}
		clusterResponse, err := e.service.Cluster.Searcher.Search(ctx, clusterRequest)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		nodeRequest := node.Request{
			ClusterID: clusterID,
			Endpoint:  clusterResponse.Endpoint,
		}
		nodeResponse, err := e.service.Node.Searcher.Search(ctx, nodeRequest)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		healthRequest := health.Request{
			Cluster:   clusterResponse.Status,
			ClusterID: clusterID,
			Nodes:     nodeResponse.Nodes,
		}
		healthResponse, err := e.service.Health.Searcher.Search(ctx, healthRequest)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		return healthResponse, nil
	}
}

func (e *Endpoint) Method() string {
	return Method
}

// Middlewares returns a slice of the middlewares used in this endpoint.
func (e *Endpoint) Middlewares() []kitendpoint.Middleware {
	return []kitendpoint.Middleware{}
}

func (e *Endpoint) Name() string {
	return Name
}

func (e *Endpoint) Path() string {
	return Path
}
