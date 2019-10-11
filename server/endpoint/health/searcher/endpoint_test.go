package searcher

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strconv"
	"testing"
	"unicode"

	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/micrologger/microloggertest"
	"github.com/giantswarm/operatorkit/client/k8srestconfig"
	"github.com/giantswarm/tenantcluster/tenantclustertest"
	"github.com/google/go-cmp/cmp"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/giantswarm/health-service/pkg/errors"
	"github.com/giantswarm/health-service/server/middleware"
	"github.com/giantswarm/health-service/service"
	"github.com/giantswarm/health-service/service/health"
	"github.com/giantswarm/health-service/service/host"
	"github.com/giantswarm/health-service/service/tenant"
)

var update = flag.Bool("update", false, "update .golden expected response files")

// NormalizeFileName converts all non-digit, non-letter runes in input string to
// dash ('-'). Coalesces multiple dashes into one.
func normalizeFileName(s string) string {
	var result []rune
	for _, r := range []rune(s) {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			result = append(result, r)
		} else {
			l := len(result)
			if l > 0 && result[l-1] != '-' {
				result = append(result, rune('-'))
			}
		}
	}
	return string(result)
}

func Test_NotFound(t *testing.T) {
	testCases := []struct {
		name         string
		clusterID    string
		errorMatcher func(err error) bool
		provider     string
	}{
		{
			name:         "case 0: aws health for non-existent cluster returns error",
			clusterID:    "abc",
			errorMatcher: errors.IsNotFound,
			provider:     "aws",
		},
		{
			name:         "case 1: azure health for non-existent cluster returns error",
			clusterID:    "abc",
			errorMatcher: errors.IsNotFound,
			provider:     "azure",
		},
		{
			name:         "case 2: kvm health for non-existent cluster returns error",
			clusterID:    "abc",
			errorMatcher: errors.IsNotFound,
			provider:     "kvm",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			// Mock k8s api server handling `kubectl get aws/azure/kvmconfig <clusterID>`.
			k8sAPIMockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("{}"))
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

			g8sClient, err := versioned.NewForConfig(restConfig)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}

			// Create tenant service.
			var tenantService *tenant.Service
			{
				tenantConfig := tenant.Config{
					Logger:        microloggertest.New(),
					TenantCluster: tenantclustertest.New(tenantclustertest.Config{}),
				}

				tenantService, err = tenant.New(tenantConfig)
				if err != nil {
					t.Fatal("Error creating tenant service: ", err)
				}
			}

			// Create host service.
			var hostService *host.Service
			{
				hostConfig := host.Config{
					G8sClient: g8sClient,
					Logger:    microloggertest.New(),
					Provider:  tc.provider,
				}

				hostService, err = host.New(hostConfig)
				if err != nil {
					t.Fatal("Error creating host service: ", err)
				}
			}

			// Create health service.
			var healthService *health.Service
			{
				healthConfig := health.Config{
					Logger: microloggertest.New(),
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
						Host:   hostService,
						Tenant: tenantService,
					},
				}
				endpoint, err = New(c)
				if err != nil {
					t.Fatal("Error on endpoint creation: ", err)
				}
			}

			// Call the endpoint and get its response.
			endpointFunc := endpoint.Endpoint()
			_, err = endpointFunc(context.Background(), tc.clusterID)

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
		})
	}
}

func Test_Health_Endpoint(t *testing.T) {
	testCases := []struct {
		name         string
		clusterID    string
		errorMatcher func(err error) bool
		provider     string
	}{
		{
			name:         "case 0: aws",
			clusterID:    "oby63",
			errorMatcher: nil,
			provider:     "aws",
		},
		{
			name:         "case 1: azure",
			clusterID:    "0cu4f",
			errorMatcher: nil,
			provider:     "azure",
		},
		{
			name:         "case 2: kvm",
			clusterID:    "cxx2e",
			errorMatcher: nil,
			provider:     "kvm",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error
			testCaseNormalizedName := normalizeFileName(tc.name)

			// Mock k8s api server handling `kubectl get aws/azure/kvmconfig <clusterID>`.
			k8sCPAPIMockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				responsePath := filepath.Join("testdata", testCaseNormalizedName+"-"+tc.provider+"config.json")
				responseJSON, err := ioutil.ReadFile(responsePath)
				if err != nil {
					t.Fatal(err)
				}
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusOK)
				w.Write(responseJSON)
			}))
			defer k8sCPAPIMockServer.Close()

			// Mock TC k8s api server handling node list and pod list.
			// Node list can be generated using:
			//   kubectl get --raw='/api/v1/nodes' | jq .
			// Pod list (filtered) can be generated using:
			//   kubectl get --raw='/api/v1/pods?fieldSelector=status.phase%21%3DFailed%2Cstatus.phase%21%3DSucceeded' | python3 -c "import sys; import json; print(json.dumps({key : [{'spec': {'nodeName': pod['spec']['nodeName'], 'containers': [{'resources': c['resources']} for c in pod['spec']['containers']]}} for pod in value] if key == 'items' else value for key, value in json.loads(sys.stdin.read()).items()}))" | jq .
			k8sTCAPIMockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var responseFilename string
				switch r.RequestURI {
				case "/api/v1/nodes":
					responseFilename = filepath.Join("testdata", testCaseNormalizedName+"-nodes.json")
				case "/api/v1/pods":
					responseFilename = filepath.Join("testdata", testCaseNormalizedName+"-pods.json")
				default:
					t.Fatal("Unknown mock request URI", r.RequestURI)
				}
				responseJSON, err := ioutil.ReadFile(responseFilename)
				if err != nil {
					t.Fatal(err)
				}
				fmt.Println(err, tc.name, responseFilename, responseJSON)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusOK)
				w.Write(responseJSON)
			}))
			defer k8sTCAPIMockServer.Close()

			var restConfig *rest.Config
			{
				c := k8srestconfig.Config{
					Logger: microloggertest.New(),

					Address:   k8sCPAPIMockServer.URL,
					InCluster: false,
				}

				restConfig, err = k8srestconfig.New(c)
				if err != nil {
					t.Fatal("expected", nil, "got", err)
				}
			}

			g8sClient, err := versioned.NewForConfig(restConfig)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}

			var restConfigTC *rest.Config
			{
				c := k8srestconfig.Config{
					Logger: microloggertest.New(),

					Address:   k8sTCAPIMockServer.URL,
					InCluster: false,
				}

				restConfigTC, err = k8srestconfig.New(c)
				if err != nil {
					t.Fatal("expected", nil, "got", err)
				}
			}

			k8sClientTC, err := kubernetes.NewForConfig(restConfigTC)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}

			// Create tenant service.
			var tenantService *tenant.Service
			{
				tenantConfig := tenant.Config{
					Logger: microloggertest.New(),
					TenantCluster: tenantclustertest.New(tenantclustertest.Config{
						K8sClient: k8sClientTC,
					}),
				}

				tenantService, err = tenant.New(tenantConfig)
				if err != nil {
					t.Fatal("Error creating tenant service: ", err)
				}
			}

			// Create host service.
			var hostService *host.Service
			{
				hostConfig := host.Config{
					G8sClient: g8sClient,
					Logger:    microloggertest.New(),
					Provider:  tc.provider,
				}

				hostService, err = host.New(hostConfig)
				if err != nil {
					t.Fatal("Error creating host service: ", err)
				}
			}

			// Create health service.
			var healthService *health.Service
			{
				healthConfig := health.Config{
					Logger: microloggertest.New(),
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
						Host:   hostService,
						Tenant: tenantService,
					},
				}
				endpoint, err = New(c)
				if err != nil {
					t.Fatal("Error on endpoint creation: ", err)
				}
			}

			// Call the endpoint and get its response.
			endpointFunc := endpoint.Endpoint()
			endpointResponse, err := endpointFunc(context.Background(), tc.clusterID)

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

			endpointResponseTyped, ok := endpointResponse.(*health.Response)
			if !ok {
				t.Fatalf("endpointResponse.(type) = %T, want %T", endpointResponse, endpointResponseTyped)
			}

			p := filepath.Join("testdata", normalizeFileName(tc.name)+".golden")

			if *update {
				json, err := json.MarshalIndent(endpointResponseTyped, "", "    ")
				if err != nil {
					t.Fatal(err)
				}
				err = ioutil.WriteFile(p, json, 0644)
				if err != nil {
					t.Fatal(err)
				}
			}
			goldenFile, err := ioutil.ReadFile(p)
			if err != nil {
				t.Fatal(err)
			}
			var goldenFileUnmarshalled health.Response
			json.Unmarshal(goldenFile, &goldenFileUnmarshalled)

			if !cmp.Equal(goldenFileUnmarshalled, *endpointResponseTyped) {
				t.Fatalf("\n\n%s\n", cmp.Diff(goldenFileUnmarshalled, *endpointResponseTyped))
			}
		})
	}
}
