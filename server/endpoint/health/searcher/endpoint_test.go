package searcher

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/micrologger/microloggertest"
	"github.com/giantswarm/operatorkit/client/k8srestconfig"
	"github.com/giantswarm/tenantcluster/tenantclustertest"
	"github.com/google/go-cmp/cmp"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/giantswarm/health-service/mock"
	"github.com/giantswarm/health-service/pkg/errors"
	"github.com/giantswarm/health-service/server/middleware"
	"github.com/giantswarm/health-service/service"
	"github.com/giantswarm/health-service/service/health"
	"github.com/giantswarm/health-service/service/health/key"
	"github.com/giantswarm/health-service/service/host"
	"github.com/giantswarm/health-service/service/tenant"
)

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
		name             string
		clusterID        string
		k8sCPAPIResponse string
		k8sTCAPIResponse string
		errorMatcher     func(err error) bool
		expectedResponse health.Response
		provider         string
	}{
		{
			name:             "case 0: aws health is returned successfully",
			clusterID:        "oby63",
			errorMatcher:     nil,
			k8sCPAPIResponse: mock.AWSHealthy,
			k8sTCAPIResponse: mock.AWSHealthyTC,
			expectedResponse: health.Response{
				Cluster: health.ClusterStatus{
					Health:    key.Green,
					State:     key.Normal,
					NodeCount: 4,
				},
				Nodes: []health.NodeStatus{
					health.NodeStatus{
						Name:                   "ip-10-1-0-172.eu-central-1.compute.internal",
						Ready:                  true,
						IP:                     "10.1.1.104",
						Hostname:               "ip-10-1-1-104.eu-central-1.compute.internal",
						InstanceType:           "m4.xlarge",
						KubeletVersion:         "v1.14.5",
						OperatorVersion:        "5.2.0",
						AvailabilityRegion:     "eu-central-1",
						AvailabilityZone:       "eu-central-1c",
						CPUCount:               4,
						MemoryCapacityBytes:    16817954816,
						MemoryAllocatableBytes: 16608239616,
						EphemeralStorageCap:    107321753600,
						EphemeralStorageAvail:  107321753600,
					},
					health.NodeStatus{
						Name:                   "ip-10-1-1-125.eu-central-1.compute.internal",
						Ready:                  true,
						IP:                     "10.1.1.125",
						Hostname:               "ip-10-1-1-125.eu-central-1.compute.internal",
						InstanceType:           "m4.xlarge",
						KubeletVersion:         "v1.14.5",
						AvailabilityRegion:     "eu-central-1",
						AvailabilityZone:       "eu-central-1c",
						CPUCount:               4,
						MemoryCapacityBytes:    16817954816,
						MemoryAllocatableBytes: 16608239616,
						EphemeralStorageCap:    107321753600,
						EphemeralStorageAvail:  107321753600,
					},
					health.NodeStatus{
						Name:                   "ip-10-1-1-57.eu-central-1.compute.internal",
						Ready:                  true,
						IP:                     "10.1.1.57",
						Hostname:               "ip-10-1-1-57.eu-central-1.compute.internal",
						InstanceType:           "m4.xlarge",
						KubeletVersion:         "v1.14.5",
						AvailabilityRegion:     "eu-central-1",
						AvailabilityZone:       "eu-central-1c",
						CPUCount:               4,
						MemoryCapacityBytes:    16817954816,
						MemoryAllocatableBytes: 16608239616,
						EphemeralStorageCap:    5843333120,
						EphemeralStorageAvail:  5843333120,
					},
					health.NodeStatus{
						Name:                   "ip-10-1-1-85.eu-central-1.compute.internal",
						Ready:                  true,
						IP:                     "10.1.1.85",
						Hostname:               "ip-10-1-1-85.eu-central-1.compute.internal",
						InstanceType:           "m4.xlarge",
						KubeletVersion:         "v1.14.5",
						AvailabilityRegion:     "eu-central-1",
						AvailabilityZone:       "eu-central-1c",
						CPUCount:               4,
						MemoryCapacityBytes:    16817954816,
						MemoryAllocatableBytes: 16608239616,
						EphemeralStorageCap:    107321753600,
						EphemeralStorageAvail:  107321753600,
					},
				},
			},
			provider: "aws",
		},
		{
			name:             "case 1: azure health is returned successfully",
			clusterID:        "0cu4f",
			errorMatcher:     nil,
			k8sCPAPIResponse: mock.AzureHealthy,
			k8sTCAPIResponse: mock.AzureHealthyTC,
			expectedResponse: health.Response{
				Cluster: health.ClusterStatus{
					Health:    "green",
					State:     key.Normal,
					NodeCount: 4,
				},
				Nodes: []health.NodeStatus{
					health.NodeStatus{
						Name:                   "6iec4-master-000000",
						Ready:                  true,
						IP:                     "10.15.0.5",
						Hostname:               "6iec4-master-000000",
						InstanceType:           "Standard_D2s_v3",
						KubeletVersion:         "v1.14.3",
						AvailabilityRegion:     "westeurope",
						AvailabilityZone:       "0",
						CPUCount:               2,
						MemoryCapacityBytes:    8340721664,
						MemoryAllocatableBytes: 8131006464,
						EphemeralStorageCap:    29137096704,
						EphemeralStorageAvail:  29137096704,
					},
					health.NodeStatus{
						Name:                   "6iec4-worker-000003",
						Ready:                  true,
						IP:                     "10.15.1.7",
						Hostname:               "6iec4-worker-000003",
						InstanceType:           "Standard_A2_v2",
						KubeletVersion:         "v1.14.3",
						AvailabilityRegion:     "westeurope",
						AvailabilityZone:       "3",
						CPUCount:               2,
						MemoryCapacityBytes:    4112863232,
						MemoryAllocatableBytes: 3903148032,
						EphemeralStorageCap:    29137096704,
						EphemeralStorageAvail:  29137096704,
					},
					health.NodeStatus{
						Name:                   "6iec4-worker-000004",
						Ready:                  true,
						IP:                     "10.15.1.4",
						Hostname:               "6iec4-worker-000004",
						InstanceType:           "Standard_A2_v2",
						KubeletVersion:         "v1.14.3",
						AvailabilityRegion:     "westeurope",
						AvailabilityZone:       "0",
						CPUCount:               2,
						MemoryCapacityBytes:    4112863232,
						MemoryAllocatableBytes: 3903148032,
						EphemeralStorageCap:    29137096704,
						EphemeralStorageAvail:  29137096704,
					},
					health.NodeStatus{
						Name:                   "6iec4-worker-000009",
						Ready:                  true,
						IP:                     "10.15.1.10",
						Hostname:               "6iec4-worker-000009",
						InstanceType:           "Standard_A2_v2",
						KubeletVersion:         "v1.14.3",
						AvailabilityRegion:     "westeurope",
						AvailabilityZone:       "4",
						CPUCount:               2,
						MemoryCapacityBytes:    4112855040,
						MemoryAllocatableBytes: 3903139840,
						EphemeralStorageCap:    29137096704,
						EphemeralStorageAvail:  29137096704,
					},
					health.NodeStatus{
						Name:                   "6iec4-worker-00000j",
						Ready:                  true,
						IP:                     "10.15.1.8",
						Hostname:               "6iec4-worker-00000j",
						InstanceType:           "Standard_A2_v2",
						KubeletVersion:         "v1.14.3",
						AvailabilityRegion:     "westeurope",
						AvailabilityZone:       "0",
						CPUCount:               2,
						MemoryCapacityBytes:    4112863232,
						MemoryAllocatableBytes: 3903148032,
						EphemeralStorageCap:    29137096704,
						EphemeralStorageAvail:  29137096704,
					},
					health.NodeStatus{
						Name:                   "6iec4-worker-00000k",
						Ready:                  true,
						IP:                     "10.15.1.9",
						Hostname:               "6iec4-worker-00000k",
						InstanceType:           "Standard_A2_v2",
						KubeletVersion:         "v1.14.3",
						AvailabilityRegion:     "westeurope",
						AvailabilityZone:       "1",
						CPUCount:               2,
						MemoryCapacityBytes:    4112863232,
						MemoryAllocatableBytes: 3903148032,
						EphemeralStorageCap:    29137096704,
						EphemeralStorageAvail:  29137096704,
					},
					health.NodeStatus{
						Name:                   "6iec4-worker-00000m",
						Ready:                  true,
						IP:                     "10.15.1.11",
						Hostname:               "6iec4-worker-00000m",
						InstanceType:           "Standard_A2_v2",
						KubeletVersion:         "v1.14.3",
						AvailabilityRegion:     "westeurope",
						AvailabilityZone:       "2",
						CPUCount:               2,
						MemoryCapacityBytes:    4112863232,
						MemoryAllocatableBytes: 3903148032,
						EphemeralStorageCap:    29137096704,
						EphemeralStorageAvail:  29137096704,
					},
				},
			},
			provider: "azure",
		},
		{
			name:             "case 2: kvm health is returned successfully",
			clusterID:        "cxx2e",
			errorMatcher:     nil,
			k8sCPAPIResponse: mock.KVMHealthy,
			k8sTCAPIResponse: mock.KVMHealthyTC,
			expectedResponse: health.Response{
				Cluster: health.ClusterStatus{
					Health:    "green",
					State:     key.Normal,
					NodeCount: 4,
				},
				Nodes: []health.NodeStatus{
					health.NodeStatus{
						Name:                   "master-8w5xy-58494dd955-7x4px",
						Ready:                  true,
						IP:                     "172.23.2.210",
						Hostname:               "master-8w5xy-58494dd955-7x4px",
						KubeletVersion:         "v1.14.6",
						CPUCount:               2,
						MemoryCapacityBytes:    8364433408,
						MemoryAllocatableBytes: 6409887744,
						EphemeralStorageCap:    5358223360,
						EphemeralStorageAvail:  4284481536,
					},
					health.NodeStatus{
						Name:                   "worker-3505b-57c995659f-55shv",
						Ready:                  true,
						IP:                     "172.23.2.22",
						Hostname:               "worker-3505b-57c995659f-55shv",
						KubeletVersion:         "v1.14.6",
						CPUCount:               2,
						MemoryCapacityBytes:    3079614464,
						MemoryAllocatableBytes: 1661939712,
						EphemeralStorageCap:    10726932480,
						EphemeralStorageAvail:  9653190656,
					},
					health.NodeStatus{
						Name:                   "worker-fqu7z-6b4d5d9cf5-vkrxt",
						Ready:                  true,
						IP:                     "172.23.2.138",
						Hostname:               "worker-fqu7z-6b4d5d9cf5-vkrxt",
						KubeletVersion:         "v1.14.6",
						CPUCount:               2,
						MemoryCapacityBytes:    3079598080,
						MemoryAllocatableBytes: 1661923328,
						EphemeralStorageCap:    10726932480,
						EphemeralStorageAvail:  9653190656,
					},
				},
			},
			provider: "kvm",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			// Mock k8s api server handling `kubectl get aws/azure/kvmconfig <clusterID>`.
			k8sCPAPIMockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tc.k8sCPAPIResponse))
			}))
			defer k8sCPAPIMockServer.Close()

			// Mock TC k8s api server handling `kubectl get nodes`.
			k8sTCAPIMockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("TCR", r.URL)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tc.k8sTCAPIResponse))
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

			if !cmp.Equal(tc.expectedResponse, *endpointResponseTyped) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.expectedResponse, *endpointResponseTyped))
			}
		})
	}
}
