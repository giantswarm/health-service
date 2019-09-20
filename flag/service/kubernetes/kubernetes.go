package kubernetes

import "github.com/giantswarm/health-service/flag/service/kubernetes/tls"

type Kubernetes struct {
	Address            string
	ApiserverExtension string
	InCluster          string
	KubeConfig         string
	TLS                tls.TLS
}
