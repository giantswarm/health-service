// Package cluster provides cluster specific business logic.
package health

import (
	"strings"

	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"

	"github.com/giantswarm/health-service/flag"
	"github.com/giantswarm/health-service/service/health/searcher"
)

// Config represents the configuration used to create a new service.
type Config struct {
	G8sClient versioned.Interface
	K8sClient kubernetes.Interface
	Logger    micrologger.Logger

	Flag  *flag.Flag
	Viper *viper.Viper
}

// DefaultConfig provides a default configuration to create a new service by
// best effort.
func DefaultConfig() Config {
	return Config{
		G8sClient: nil,
		K8sClient: nil,
		Logger:    nil,

		Flag:  nil,
		Viper: nil,
	}
}

type Service struct {
	Searcher *searcher.Service
}

// New creates a new configured service object.
func New(config Config) (*Service, error) {
	if config.Flag == nil {
		return nil, microerror.Maskf(invalidConfigError, "config.Flag must not be empty")
	}
	if config.K8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "config.K8sClient must not be empty")
	}
	if config.Viper == nil {
		return nil, microerror.Maskf(invalidConfigError, "config.Viper must not be empty")
	}

	var err error

	var searcherService *searcher.Service
	{
		searcherConfig := searcher.DefaultConfig()
		searcherConfig.G8sClient = config.G8sClient
		searcherConfig.Logger = config.Logger
		searcherConfig.Provider = strings.TrimSpace(config.Viper.GetString(config.Flag.Service.Provider.Kind))
		searcherService, err = searcher.New(searcherConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	newService := &Service{
		Searcher: searcherService,
	}

	return newService, nil
}
