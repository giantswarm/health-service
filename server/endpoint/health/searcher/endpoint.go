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

	"github.com/giantswarm/health-service/pkg/errors"
	"github.com/giantswarm/health-service/server/middleware"
	"github.com/giantswarm/health-service/service"
	"github.com/giantswarm/health-service/service/health"
	"github.com/giantswarm/health-service/service/host"
	"github.com/giantswarm/health-service/service/tenant/nodes"
	"github.com/giantswarm/health-service/service/tenant/pods"
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
		return nil, microerror.Maskf(errors.InvalidConfigError, "config.Logger must not be empty")
	}
	if config.Middleware == nil {
		return nil, microerror.Maskf(errors.InvalidConfigError, "config.Middleware must not be empty")
	}
	if config.Service == nil {
		return nil, microerror.Maskf(errors.InvalidConfigError, "config.Service must not be empty")
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
			return nil, microerror.Mask(errors.BadRequestError)
		}

		hostRequest := host.Request{
			ClusterID: clusterID,
		}
		hostResponse, err := e.service.Host.SearchStatusCluster(ctx, hostRequest)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		nodesRequest := nodes.Request{
			ClusterID: clusterID,
			Endpoint:  hostResponse.Endpoint,
		}
		nodesResponse, err := e.service.Tenant.Nodes.ListNodes(ctx, nodesRequest)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		podsRequest := pods.Request{
			ClusterID: clusterID,
			Endpoint:  hostResponse.Endpoint,
		}
		podsResponse, err := e.service.Tenant.Pods.ListPods(ctx, podsRequest)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		healthRequest := health.Request{
			Cluster:   hostResponse.Status,
			ClusterID: clusterID,
			Nodes:     nodesResponse.Nodes,
			Pods:      podsResponse.Pods,
		}
		healthResponse, err := e.service.Health.Search(ctx, healthRequest)
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
