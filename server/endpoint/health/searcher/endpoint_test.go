package searcher

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/health-service/service/health"
	"github.com/giantswarm/micrologger/microloggertest"
	"github.com/giantswarm/operatorkit/client/k8srestconfig"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/giantswarm/health-service/flag"
	"github.com/giantswarm/health-service/server/middleware"
	"github.com/giantswarm/health-service/service"
)

var testCases = []struct {
	clusterID      string
	k8sAPIResponse string
	expectedError  func(err error) bool
	expectedHealth string
}{
	{
		clusterID:      "abc",
		k8sAPIResponse: `{"items": []}`,
		expectedError:  nil,
		expectedHealth: "green",
	},
	{
		clusterID:      "def",
		k8sAPIResponse: `{"items": [{},{},{},{},{}]}`,
		expectedError:  nil,
		expectedHealth: "green",
	},
}

func TestEndpoint(t *testing.T) {
	for i, testCase := range testCases {
		var err error

		// Mock k8s api server handling `kubectl get nodes`.
		k8sAPIMockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(testCase.k8sAPIResponse))
		}))
		defer k8sAPIMockServer.Close()

		var restConfig *rest.Config
		{
			c := k8srestconfig.Config{
				Logger: microloggertest.New(),

				Address:   k8sAPIMockServer.URL,
				InCluster: false,
			}

			restConfig, err = k8srestconfig.New(c)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}
		}

		k8sClient, err := kubernetes.NewForConfig(restConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		g8sClient, err := versioned.NewForConfig(restConfig)
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		f := flag.New()
		v := viper.New()
		v.Set(f.Service.Provider.Kind, "aws")

		// Create health service.
		var healthService *health.Service
		{
			healthConfig := health.Config{
				Flag:      f,
				Viper:     v,
				K8sClient: k8sClient,
				G8sClient: g8sClient,
				Logger:    microloggertest.New(),
			}

			healthService, err = health.New(healthConfig)
			if err != nil {
				t.Fatal("Error creating health service: ", err)
			}
		}

		// Create health endpoint.
		cfg := Config{
			Logger:     microloggertest.New(),
			Middleware: &middleware.Middleware{},
			Service: &service.Service{
				Health: healthService,
			},
		}

		endpoint, err := New(cfg)
		if err != nil {
			t.Fatal("Error on endpoint creation: ", err)
		}

		req, err := http.NewRequest(Method, Path, nil)
		if err != nil {
			t.Fatal("Unexpected HTTP error:", err)
		}

		ctx := req.Context()
		type contextKey string
		ctx = context.WithValue(ctx, contextKey("Content-Type"), "application/json; charset=utf-8")

		// Call the endpoint and get its response.
		endpointFunc := endpoint.Endpoint()
		endpointResponse, err := endpointFunc(ctx, testCase.clusterID)
		if testCase.expectedError != nil {
			if !testCase.expectedError(err) {
				t.Fatal("Case", i, "- calling the endpoint didn't yield the expected error. got:", err)
			}
		} else {
			if err != nil {
				t.Fatal("Case", i, "- calling the endpoint yielded an error where we expected none:", err)
			}
		}

		if endpointResponse.(Response).ClusterHealth != testCase.expectedHealth {
			t.Fatal("Case", i, "- expected ", testCase.expectedHealth, ", got: ", endpointResponse.(Response).ClusterHealth)
		}
	}
}
