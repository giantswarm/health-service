package health

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"testing"
	"unicode"

	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	"github.com/giantswarm/health-service/service/health/key"
	"github.com/giantswarm/micrologger/microloggertest"
	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func Test_Health_New(t *testing.T) {
	testCases := []struct {
		name         string
		inputObj     Config
		errorMatcher func(err error) bool
	}{
		{
			name: "case 0: a service is successfully created",
			inputObj: Config{
				Logger: microloggertest.New(),
			},
			errorMatcher: nil,
		},
		{
			name:         "case 1: invalidConfigError returned when logger is missing",
			inputObj:     Config{},
			errorMatcher: IsInvalidConfig,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			_, err := New(tc.inputObj)
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

var readyStatus = v1.NodeStatus{
	Conditions: []v1.NodeCondition{
		v1.NodeCondition{
			Type:   v1.NodeReady,
			Status: v1.ConditionTrue,
		},
	},
}

var notReadyStatus = v1.NodeStatus{
	Conditions: []v1.NodeCondition{
		v1.NodeCondition{
			Type:   v1.NodeReady,
			Status: v1.ConditionFalse,
		},
	},
}

var workerMeta = metav1.ObjectMeta{
	Labels: map[string]string{
		"kubernetes.io/role": "worker",
	},
}

var masterMeta = metav1.ObjectMeta{
	Labels: map[string]string{
		"kubernetes.io/role": "master",
	},
}

var readyWorkerNode = v1.Node{
	ObjectMeta: workerMeta,
	Status:     readyStatus,
}

var readyMasterNode = v1.Node{
	ObjectMeta: masterMeta,
	Status:     readyStatus,
}

var notReadyWorkerNode = v1.Node{
	ObjectMeta: workerMeta,
	Status:     notReadyStatus,
}

var notReadyMasterNode = v1.Node{
	ObjectMeta: masterMeta,
	Status:     notReadyStatus,
}

var healthyNodes = []v1.Node{
	readyMasterNode,
	readyWorkerNode,
	readyWorkerNode,
	readyWorkerNode,
}

func Test_ClusterHealth(t *testing.T) {
	testCases := []struct {
		name           string
		clusterStatus  v1alpha1.StatusCluster
		nodes          []v1.Node
		expectedHealth key.Health
	}{
		{
			name:           "case 0: healthy cluster should be green",
			expectedHealth: key.HealthGreen,
			nodes:          healthyNodes,
		},
		{
			name:           "case 1: creating cluster should be yellow",
			expectedHealth: key.HealthYellow,
			clusterStatus: v1alpha1.StatusCluster{
				Conditions: []v1alpha1.StatusClusterCondition{
					v1alpha1.StatusClusterCondition{
						Type:   v1alpha1.StatusClusterTypeCreating,
						Status: v1alpha1.StatusClusterStatusTrue,
					},
				},
			},
			nodes: healthyNodes,
		},
		{
			name:           "case 2: updating cluster should be green",
			expectedHealth: key.HealthGreen,
			clusterStatus: v1alpha1.StatusCluster{
				Conditions: []v1alpha1.StatusClusterCondition{
					v1alpha1.StatusClusterCondition{
						Type:   v1alpha1.StatusClusterTypeUpdating,
						Status: v1alpha1.StatusClusterStatusTrue,
					},
				},
			},
			nodes: healthyNodes,
		},
		{
			name:           "case 3: deleting cluster should be yellow",
			expectedHealth: key.HealthYellow,
			clusterStatus: v1alpha1.StatusCluster{
				Conditions: []v1alpha1.StatusClusterCondition{
					v1alpha1.StatusClusterCondition{
						Type:   v1alpha1.StatusClusterTypeDeleting,
						Status: v1alpha1.StatusClusterStatusTrue,
					},
				},
			},
			nodes: healthyNodes,
		},
		{
			name:           "case 4: not ready master should be red",
			expectedHealth: key.HealthRed,
			nodes: []v1.Node{
				notReadyMasterNode,
				readyWorkerNode,
				readyWorkerNode,
				readyWorkerNode,
			},
		},
		{
			name:           "case 5: not ready master during update should be yellow",
			expectedHealth: key.HealthYellow,
			clusterStatus: v1alpha1.StatusCluster{
				Conditions: []v1alpha1.StatusClusterCondition{
					v1alpha1.StatusClusterCondition{
						Type:   v1alpha1.StatusClusterTypeUpdating,
						Status: v1alpha1.StatusClusterStatusTrue,
					},
				},
			},
			nodes: []v1.Node{
				notReadyMasterNode,
				readyWorkerNode,
				readyWorkerNode,
				readyWorkerNode,
			},
		},
		{
			name:           "case 6: >=3 but <75% ready workers should be yellow",
			expectedHealth: key.HealthYellow,
			nodes: []v1.Node{
				readyMasterNode,
				notReadyWorkerNode,
				notReadyWorkerNode,
				readyWorkerNode,
				readyWorkerNode,
				readyWorkerNode,
			},
		},
		{
			name:           "case 7: <3 ready workers should be red",
			expectedHealth: key.HealthRed,
			nodes: []v1.Node{
				readyMasterNode,
				readyWorkerNode,
				readyWorkerNode,
			},
		},
		{
			name:           "case 8: >20 workers and <50% ready should be red",
			expectedHealth: key.HealthRed,
			nodes: []v1.Node{
				readyMasterNode,
				readyWorkerNode,
				readyWorkerNode,
				readyWorkerNode,
				readyWorkerNode,
				readyWorkerNode,
				readyWorkerNode,
				readyWorkerNode,
				readyWorkerNode,
				notReadyWorkerNode,
				notReadyWorkerNode,
				notReadyWorkerNode,
				notReadyWorkerNode,
				notReadyWorkerNode,
				notReadyWorkerNode,
				notReadyWorkerNode,
				notReadyWorkerNode,
				notReadyWorkerNode,
				notReadyWorkerNode,
				notReadyWorkerNode,
				notReadyWorkerNode,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			// Create health service.
			var service *Service
			{
				config := Config{
					Logger: microloggertest.New(),
				}

				service, err = New(config)
				if err != nil {
					t.Fatal("Error creating health service: ", err)
				}
			}

			request := Request{
				Cluster: Cluster{
					Status: tc.clusterStatus,
					Nodes:  tc.nodes,
				},
			}
			response, err := service.Search(context.Background(), request)

			if !cmp.Equal(tc.expectedHealth, response.Cluster.Health) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.expectedHealth, response.Cluster.Health))
			}
		})
	}
}

func Test_NodeParsing(t *testing.T) {
	testCases := []struct {
		name     string
		provider string
	}{
		{
			name:     "case 0: aws node",
			provider: "aws",
		},
		{
			name:     "case 1: azure node",
			provider: "azure",
		},
		{
			name:     "case 2: kvm node",
			provider: "kvm",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			// Create health service.
			var service *Service
			{
				config := Config{
					Logger: microloggertest.New(),
				}

				service, err = New(config)
				if err != nil {
					t.Fatal("Error creating health service: ", err)
				}
			}

			testCaseNormalizedName := normalizeFileName(tc.name)
			nodeDataPath := filepath.Join("testdata", testCaseNormalizedName+".json")
			nodeDataJSON, err := ioutil.ReadFile(nodeDataPath)
			var node v1.Node
			err = json.Unmarshal(nodeDataJSON, &node)
			if err != nil {
				t.Fatal(err)
			}
			request := Request{
				Cluster: Cluster{
					Nodes: []v1.Node{
						node,
					},
				},
			}
			response, err := service.Search(context.Background(), request)
			if err != nil {
				t.Fatal(err)
			}
			if len(response.Nodes) == 0 {
				t.Fatalf("no node statuses returned")
			}
			actualNodeStatus := response.Nodes[0]

			p := filepath.Join("testdata", testCaseNormalizedName+".golden")

			if *update {
				json, err := json.MarshalIndent(actualNodeStatus, "", "    ")
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
			var expectedNodeStatus NodeStatus
			err = json.Unmarshal(goldenFile, &expectedNodeStatus)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(expectedNodeStatus, actualNodeStatus) {
				t.Fatalf("\n\n%s\n", cmp.Diff(expectedNodeStatus, actualNodeStatus))
			}
		})
	}
}
