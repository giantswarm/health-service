package service

import (
	"github.com/giantswarm/health-service/flag/service/kubernetes"
	"github.com/giantswarm/health-service/flag/service/provider"
)

type Service struct {
	Kubernetes kubernetes.Kubernetes
	Provider   provider.Provider
}
