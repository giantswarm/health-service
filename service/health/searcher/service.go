package searcher

import (
	"context"

	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/health-service/service/health/key"
)

// Config represents the configuration used to create a new service object.
type Config struct {
	G8sClient versioned.Interface
	Logger    micrologger.Logger

	Provider string
}

type Service struct {
	g8sClient versioned.Interface
	logger    micrologger.Logger

	provider string
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	if config.G8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.G8sClient must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	if config.Provider == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.Provider must not be empty", config)
	}
	if !key.IsKnownProvider(config.Provider) {
		return nil, microerror.Maskf(invalidConfigError, "%T.Provider must be one of %+v", config, key.Providers)
	}

	s := &Service{
		g8sClient: config.G8sClient,
		logger:    config.Logger,

		provider: config.Provider,
	}

	return s, nil
}

func extractStatusCluster(status v1alpha1.StatusCluster) ClusterInfo {
	nodes := []NodeInfo{}
	for _, node := range status.Nodes {
		nodes = append(nodes, NodeInfo{
			labels: node.Labels,
		})
	}
	return ClusterInfo{
		nodes,
	}
}

// searchAWSCR searches for the cluster config in AWSClusterConfigs resources.
func (s *Service) searchAWSInfo(ctx context.Context, clusterID string) (*ClusterInfo, error) {
	cr, err := s.g8sClient.ProviderV1alpha1().AWSConfigs("default").Get(clusterID, v1.GetOptions{})
	if err != nil {
		return nil, microerror.Mask(err)
	}
	clusterInfo := extractStatusCluster(cr.Status.Cluster)
	return &clusterInfo, nil
}

// searchAzureCR searches for the cluster config in AWSClusterConfigs resources.
func (s *Service) searchAzureInfo(ctx context.Context, clusterID string) (*ClusterInfo, error) {
	cr, err := s.g8sClient.ProviderV1alpha1().AzureConfigs("default").Get(clusterID, v1.GetOptions{})
	if err != nil {
		return nil, microerror.Mask(err)
	}
	clusterInfo := extractStatusCluster(cr.Status.Cluster)
	return &clusterInfo, nil
}

// searchKVMCR abc
func (s *Service) searchKVMInfo(ctx context.Context, clusterID string) (*ClusterInfo, error) {
	cr, err := s.g8sClient.ProviderV1alpha1().KVMConfigs("default").Get(clusterID, v1.GetOptions{})
	if err != nil {
		return nil, microerror.Mask(err)
	}
	clusterInfo := extractStatusCluster(cr.Status.Cluster)
	return &clusterInfo, nil
}

func evaluateGeneralHealth(cluster *ClusterInfo) GeneralStatus {
	nodeCount := len(cluster.nodes)
	generalHealth := "red"
	if nodeCount > 3 {
		generalHealth = "green"
	} else if nodeCount > 1 {
		generalHealth = "yellow"
	}
	return GeneralStatus{
		Health: generalHealth,
	}
}

// Search searches for the cluster information.
// It try to find cluster information in CR and fallback to storage service when nothing is found.
func (s *Service) Search(ctx context.Context, request Request) (*Response, error) {
	var err error
	var cluster *ClusterInfo

	switch s.provider {
	case "aws":
		cluster, err = s.searchKVMInfo(ctx, request.ClusterID)
	case "azure":
		cluster, err = s.searchAzureInfo(ctx, request.ClusterID)
	case "kvm":
		cluster, err = s.searchKVMInfo(ctx, request.ClusterID)
	default:
		return nil, microerror.Maskf(invalidConfigError, "unsupported provider: %s", s.provider)
	}

	if err != nil {
		return nil, microerror.Mask(err)
	}

	return &Response{
		General: evaluateGeneralHealth(cluster),
		Nodes:   []NodeStatus{},
	}, nil
}
