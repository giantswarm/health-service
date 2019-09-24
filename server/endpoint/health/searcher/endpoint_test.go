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

var (
	awsMockResponse = `{}`
	azureMockResponse = `{}`
	kvmMockResponse = `{"apiVersion":"provider.giantswarm.io/v1alpha1","kind":"KVMConfig","metadata":{"creationTimestamp":"2019-09-23T11:53:38Z","name":"cxx2e","namespace":"default"},"spec":{"cluster":{"calico":{"cidr":16,"mtu":1430,"subnet":"192.168.0.0"},"customer":{"id":"giantswarm"},"docker":{"daemon":{"cidr":"172.17.0.1/16"}},"id":"cxx2e","masters":[{"id":"hhn5a"}],"scaling":{"max":3,"min":3},"version":"","workers":[{"id":"hi5ie"},{"id":"pq295"},{"id":"oh8xv"}]},"kvm":{"masters":[{"cpus":2,"disk":40,"dockerVolumeSizeGB":0,"memory":"8G"}],"network":{"flannel":{"vni":4}},"workers":[{"cpus":2,"disk":40,"dockerVolumeSizeGB":0,"memory":"3G"},{"cpus":2,"disk":40,"dockerVolumeSizeGB":0,"memory":"3G"},{"cpus":2,"disk":40,"dockerVolumeSizeGB":0,"memory":"3G"}]},"versionBundle":{"version":"3.8.0"}},"status":{"cluster":{"conditions":[{"lastTransitionTime":"2019-09-23T12:08:24.66025979Z","status":"True","type":"Updating"},{"lastTransitionTime":"2019-09-23T12:01:19.857790161Z","status":"True","type":"Created"}],"network":{"cidr":""},"nodes":[{"labels":{"beta.kubernetes.io/arch":"amd64","beta.kubernetes.io/os":"linux","giantswarm.io/provider":"kvm","ip":"172.23.0.206","kubernetes.io/arch":"amd64","kubernetes.io/hostname":"master-hhn5a-66b4789bbc-tjvb4","kubernetes.io/os":"linux","kubernetes.io/role":"master","kvm-operator.giantswarm.io/version":"3.8.0","node-role.kubernetes.io/master":"","node.kubernetes.io/master":"","role":"master"},"lastTransitionTime":"2019-09-23T17:55:06.434468696Z","name":"master-hhn5a-66b4789bbc-tjvb4","version":"3.8.0"},{"labels":{"beta.kubernetes.io/arch":"amd64","beta.kubernetes.io/os":"linux","giantswarm.io/provider":"kvm","ip":"172.23.0.154","kubernetes.io/arch":"amd64","kubernetes.io/hostname":"worker-hi5ie-5f79675f59-5wgss","kubernetes.io/os":"linux","kubernetes.io/role":"worker","kvm-operator.giantswarm.io/version":"3.8.0","node-role.kubernetes.io/worker":"","node.kubernetes.io/worker":"","role":"worker"},"lastTransitionTime":"2019-09-23T17:55:06.434469572Z","name":"worker-hi5ie-5f79675f59-5wgss","version":"3.8.0"},{"labels":{"beta.kubernetes.io/arch":"amd64","beta.kubernetes.io/os":"linux","giantswarm.io/provider":"kvm","ip":"172.23.0.242","kubernetes.io/arch":"amd64","kubernetes.io/hostname":"worker-oh8xv-c549d84f-24z2s","kubernetes.io/os":"linux","kubernetes.io/role":"worker","kvm-operator.giantswarm.io/version":"3.5.0","node-role.kubernetes.io/worker":"","node.kubernetes.io/worker":"","role":"worker"},"lastTransitionTime":"2019-09-23T17:55:06.434470124Z","name":"worker-oh8xv-c549d84f-24z2s","version":"3.5.0"},{"labels":{"beta.kubernetes.io/arch":"amd64","beta.kubernetes.io/os":"linux","giantswarm.io/provider":"kvm","ip":"172.23.0.86","kubernetes.io/arch":"amd64","kubernetes.io/hostname":"worker-pq295-6c4bbcccf7-kql2k","kubernetes.io/os":"linux","kubernetes.io/role":"worker","kvm-operator.giantswarm.io/version":"3.5.0","node-role.kubernetes.io/worker":"","node.kubernetes.io/worker":"","role":"worker"},"lastTransitionTime":"2019-09-23T17:55:06.434470788Z","name":"worker-pq295-6c4bbcccf7-kql2k","version":"3.5.0"}],"resources":null,"scaling":{"desiredCapacity":0},"versions":[{"date":"0001-01-01T00:00:00Z","lastTransitionTime":"2019-09-23T12:01:19.910404547Z","semver":"3.5.0"}]},"kvm":{"nodeIndexes":{"hhn5a":1,"hi5ie":2,"oh8xv":4,"pq295":3}}}}`
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
			k8sAPIResponse: awsMockResponse,
			expectedHealth: "green",
			provider:       "aws",
		},
		{
			name:           "case 1: azure health is returned successfully",
			inputObj:       "abc",
			errorMatcher:   nil,
			k8sAPIResponse: azureMockResponse,
			expectedHealth: "green",
			provider:       "azure",
		},
		{
			name:           "case 2: kvm health is returned successfully",
			inputObj:       "abc",
			errorMatcher:   nil,
			k8sAPIResponse: kvmMockResponse,
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
<<<<<<< HEAD
				t.Fatal("endpointResponse.(type) = %T, want %T", endpointResponse, endpointResponseTyped)
=======
				t.Fatalf("endpointResponse.(type) = %T, want %T", endpointResponse, endpointResponseTyped)
>>>>>>> master
			}

			if !cmp.Equal(endpointResponseTyped.ClusterHealth, tc.expectedHealth) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.expectedHealth, endpointResponseTyped.ClusterHealth))
			}
		})
	}
}
