package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/microkit/command"
	microserver "github.com/giantswarm/microkit/server"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/client/k8srestconfig"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/giantswarm/health-service/flag"
	"github.com/giantswarm/health-service/pkg/project"
	"github.com/giantswarm/health-service/server"
	"github.com/giantswarm/health-service/service"
)

var (
	f *flag.Flag = flag.New()
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	err := mainWithError()
	if err != nil {
		panic(fmt.Sprintf("%#v\n", microerror.Mask(err)))
	}
}

func mainWithError() error {
	var err error

	// Create a new logger which is used by all packages.
	var newLogger micrologger.Logger
	{
		newLogger, err = micrologger.New(micrologger.Config{})
		if err != nil {
			return microerror.Mask(err)
		}
	}

	// We define a server factory to create the custom server once all command
	// line flags are parsed and all microservice configuration is storted out.
	newServerFactory := func(v *viper.Viper) microserver.Server {
		var restConfig *rest.Config
		{
			c := k8srestconfig.Config{
				Logger: newLogger,

				Address:   v.GetString(f.Service.Kubernetes.Address),
				InCluster: v.GetBool(f.Service.Kubernetes.InCluster),
				TLS: k8srestconfig.ConfigTLS{
					CAFile:  v.GetString(f.Service.Kubernetes.TLS.CAFile),
					CrtFile: v.GetString(f.Service.Kubernetes.TLS.CrtFile),
					KeyFile: v.GetString(f.Service.Kubernetes.TLS.KeyFile),
				},
			}

			restConfig, err = k8srestconfig.New(c)
			if err != nil {
				panic(fmt.Sprintf("%#v", err))
			}
		}

		k8sClient, err := kubernetes.NewForConfig(restConfig)
		if err != nil {
			panic(fmt.Sprintf("%#v", err))
		}

		// Create a new custom service which implements business logic.
		var newService *service.Service
		{
			serviceConfig := service.Config{
				Flag:      f,
				K8sClient: k8sClient,
				Logger:    newLogger,
				Viper:     v,

				Description: project.Description(),
				GitCommit:   project.GitSHA(),
				Name:        project.Name(),
				Source:      project.Source(),
				Version:     project.Version(),
			}

			newService, err = service.New(serviceConfig)
			if err != nil {
				panic(fmt.Sprintf("%#v", err))
			}
		}

		// Create a new custom server which bundles our endpoints.
		var newServer microserver.Server
		{
			serverConfig := server.Config{
				Logger:  newLogger,
				Service: newService,
				Viper:   v,

				ProjectName: project.Name(),
			}

			newServer, err = server.New(serverConfig)
			if err != nil {
				panic(fmt.Sprintf("%#v", err))
			}
		}

		return newServer
	}

	// Create a new microkit command which manages our custom microservice.
	var newCommand command.Command
	{
		c := command.Config{
			Logger:        newLogger,
			ServerFactory: newServerFactory,

			Description: project.Description(),
			GitCommit:   project.GitSHA(),
			Name:        project.Name(),
			Source:      project.Source(),
			Version:     project.Version(),
		}

		newCommand, err = command.New(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	daemonCommand := newCommand.DaemonCommand().CobraCommand()

	daemonCommand.PersistentFlags().String(f.Service.Provider.Kind, "", "Kind of the underlying virtualization/technology provider the service is deployed in.")

	daemonCommand.PersistentFlags().String(f.Service.Kubernetes.GuestAPIEndpointFormat, "", "Format string used to create the API endpoint of the guest Kubernetes clusters.")
	daemonCommand.PersistentFlags().String(f.Service.Kubernetes.Address, "http://127.0.0.1:8080", "Address used to connect to Kubernetes. When empty in-cluster config is created.")
	daemonCommand.PersistentFlags().Bool(f.Service.Kubernetes.InCluster, true, "Whether to use the in-cluster config to authenticate with Kubernetes.")
	daemonCommand.PersistentFlags().String(f.Service.Kubernetes.KubeConfig, "", "KubeConfig used to connect to Kubernetes. When empty other settings are used.")
	daemonCommand.PersistentFlags().String(f.Service.Kubernetes.TLS.CAFile, "", "Certificate authority file path to use to authenticate with Kubernetes.")
	daemonCommand.PersistentFlags().String(f.Service.Kubernetes.TLS.CrtFile, "", "Certificate file path to use to authenticate with Kubernetes.")
	daemonCommand.PersistentFlags().String(f.Service.Kubernetes.TLS.KeyFile, "", "Key file path to use to authenticate with Kubernetes.")

	newCommand.CobraCommand().Execute()

	return nil
}
