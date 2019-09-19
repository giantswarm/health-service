package kubernetes

type TLS struct {
	CAFile  string
	CrtFile string
	KeyFile string
}

type Kubernetes struct {
	Address                string
	GuestAPIEndpointFormat string
	InCluster              string
	TLS                    TLS
	KubeConfig             string
}
