// Package service implements business logic to create Kubernetes resources
// against the Kubernetes API.
package service

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/microendpoint/service/version"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/microstorage"
	"github.com/giantswarm/operatorkit/client/k8srestconfig"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/giantswarm/health-service/flag"
	"github.com/giantswarm/health-service/service/health"
)

const (
	// DefaultRetryCount is the number of times to retry a failed network call.
	DefaultRetryCount = 5
	// DefaultTimeout is the timeout for network calls.
	DefaultTimeout = 5 * time.Second
)

// Config represents the configuration used to create a new service.
type Config struct {
	// Dependencies.
	K8sClient kubernetes.Interface
	Logger    micrologger.Logger
	Storage   microstorage.Storage

	// Settings.
	Flag  *flag.Flag
	Viper *viper.Viper

	Description string
	GitCommit   string
	Name        string
	Source      string
	Version     string
}

type Service struct {
	Health  *health.Service
	Version *version.Service

	bootOnce sync.Once
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	if config.K8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "config.K8SClient must not be empty")
	}

	if config.Flag == nil {
		return nil, microerror.Maskf(invalidConfigError, "config.Flag must not be empty")
	}
	if config.Viper == nil {
		return nil, microerror.Maskf(invalidConfigError, "config.Viper must not be empty")
	}

	var err error

	var restConfig *rest.Config
	{
		c := k8srestconfig.Config{
			Logger: config.Logger,

			Address:    config.Viper.GetString(config.Flag.Service.Kubernetes.Address),
			InCluster:  config.Viper.GetBool(config.Flag.Service.Kubernetes.InCluster),
			KubeConfig: config.Viper.GetString(config.Flag.Service.Kubernetes.KubeConfig),
			TLS: k8srestconfig.ConfigTLS{
				CAFile:  config.Viper.GetString(config.Flag.Service.Kubernetes.TLS.CAFile),
				CrtFile: config.Viper.GetString(config.Flag.Service.Kubernetes.TLS.CrtFile),
				KeyFile: config.Viper.GetString(config.Flag.Service.Kubernetes.TLS.KeyFile),
			},
		}

		restConfig, err = k8srestconfig.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	g8sClient, err := versioned.NewForConfig(restConfig)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var healthService *health.Service
	{
		healthConfig := health.DefaultConfig()

		healthConfig.K8sClient = config.K8sClient
		healthConfig.Logger = config.Logger
		healthConfig.G8sClient = g8sClient
		healthConfig.Flag = config.Flag
		healthConfig.Viper = config.Viper

		healthService, err = health.New(healthConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var versionService *version.Service
	{
		versionConfig := version.Config{
			Description: config.Description,
			GitCommit:   config.GitCommit,
			Name:        config.Name,
			Source:      config.Source,
			Version:     config.Version,
		}

		versionService, err = version.New(versionConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	s := &Service{
		Health:  healthService,
		Version: versionService,

		bootOnce: sync.Once{},
	}

	return s, nil
}

func stringToInstanceTypes(s string) ([]string, error) {
	var instanceTypes []string
	err := json.Unmarshal([]byte(s), &instanceTypes)
	return instanceTypes, err
}
