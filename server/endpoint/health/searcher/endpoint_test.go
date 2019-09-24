package searcher

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/micrologger/microloggertest"
	"github.com/giantswarm/operatorkit/client/k8srestconfig"
	"github.com/google/go-cmp/cmp"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/giantswarm/health-service/flag"
	"github.com/giantswarm/health-service/server/middleware"
	"github.com/giantswarm/health-service/service"
	"github.com/giantswarm/health-service/service/health"
)

func Test_Health_Endpoint(t *testing.T) {
	testCases := []struct {
		name           string
		inputObj       string
		k8sAPIResponse string
		errorMatcher   func(err error) bool
		expectedHealth string
		provider       string
	}{
		{
			name:           "case 0: aws health is returned successfully",
			inputObj:       "abc",
			errorMatcher:   nil,
			k8sAPIResponse: `{"items": []}`,
			expectedHealth: "green",
			provider:       "aws",
		},
		{
			name:           "case 1: azure health is returned successfully",
			inputObj:       "abc",
			errorMatcher:   nil,
			k8sAPIResponse: `{"items": []}`,
			expectedHealth: "green",
			provider:       "azure",
		},
		{
			name:           "case 2: kvm health is returned successfully",
			inputObj:       "abc",
			errorMatcher:   nil,
			k8sAPIResponse: `{"items": []}`,
			expectedHealth: "green",
			provider:       "kvm",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			// Mock k8s api server handling `kubectl get aws/azure/kvmconfig <clusterID>`.
			k8sAPIMockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tc.k8sAPIResponse))
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
			v.Set(f.Service.Provider.Kind, tc.provider)

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
			var endpoint *Endpoint
			{
				c := Config{
					Logger:     microloggertest.New(),
					Middleware: &middleware.Middleware{},
					Service: &service.Service{
						Health: healthService,
					},
				}
				endpoint, err = New(c)
				if err != nil {
					t.Fatal("Error on endpoint creation: ", err)
				}
			}

			// Call the endpoint and get its response.
			endpointFunc := endpoint.Endpoint()
			endpointResponse, err := endpointFunc(context.Background(), tc.inputObj)

			switch {
			case err == nil && tc.errorMatcher == nil:
				// correct; carry on
			case err != nil && tc.errorMatcher == nil:
				t.Fatalf("error == %#v, want nil", err)
			case err == nil && tc.errorMatcher != nil:
				t.Fatalf("error == nil, want non-nil")
			case !tc.errorMatcher(err):
				t.Fatalf("error == %#v, want matching", err)
			}

			endpointResponseTyped, ok := endpointResponse.(Response)
			if !ok {
				t.Fatalf("endpointResponse.(type) = %T, want %T", endpointResponse, endpointResponseTyped)
			}

			if !cmp.Equal(endpointResponseTyped.ClusterHealth, tc.expectedHealth) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.expectedHealth, endpointResponseTyped.ClusterHealth))
			}
		})
	}
}
