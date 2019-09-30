package mock

// AWSHealthy is a JSON string representing a healthy AWS cluster
var AWSHealthy = `
{
	"apiVersion": "provider.giantswarm.io/v1alpha1",
	"kind": "AWSConfig",
	"metadata": {
		"creationTimestamp": "2019-08-05T10:41:07Z",
		"finalizers": [
			"operatorkit.giantswarm.io/aws-operator-drainer",
			"operatorkit.giantswarm.io/aws-operator"
		],
		"generation": 2,
		"labels": {
			"aws-operator.giantswarm.io/version": "5.2.0"
		},
		"name": "oby63",
		"namespace": "default",
		"resourceVersion": "22712556",
		"selfLink": "/apis/provider.giantswarm.io/v1alpha1/namespaces/default/awsconfigs/oby63",
		"uid": "8850cee7-b76d-11e9-bbd0-06322d0718ec"
	},
	"spec": {
		"aws": {
			"api": {
				"elb": {
					"idleTimeoutSeconds": 0
				},
				"hostedZones": ""
			},
			"availabilityZones": 2,
			"az": "eu-central-1a",
			"credentialSecret": {
				"name": "credential-default",
				"namespace": "giantswarm"
			},
			"etcd": {
				"elb": {
					"idleTimeoutSeconds": 0
				},
				"hostedZones": ""
			},
			"hostedZones": {
				"api": {
					"name": "gorilla.eu-central-1.aws.gigantic.io"
				},
				"etcd": {
					"name": "gorilla.eu-central-1.aws.gigantic.io"
				},
				"ingress": {
					"name": "gorilla.eu-central-1.aws.gigantic.io"
				}
			},
			"ingress": {
				"elb": {
					"idleTimeoutSeconds": 0
				},
				"hostedZones": ""
			},
			"masters": [
				{
					"dockerVolumeSizeGB": 0,
					"imageID": "ami-90c152ff",
					"instanceType": "m4.large"
				}
			],
			"region": "eu-central-1",
			"vpc": {
				"cidr": "",
				"peerId": "vpc-097f163f87959f9b4",
				"privateSubnetCidr": "",
				"publicSubnetCidr": "",
				"routeTableNames": [
					"gorilla_private_0",
					"gorilla_private_1",
					"gorilla_private_2"
				]
			},
			"workers": [
				{
					"dockerVolumeSizeGB": 0,
					"imageID": "ami-90c152ff",
					"instanceType": "m4.xlarge"
				},
				{
					"dockerVolumeSizeGB": 0,
					"imageID": "ami-90c152ff",
					"instanceType": "m4.xlarge"
				},
				{
					"dockerVolumeSizeGB": 0,
					"imageID": "ami-90c152ff",
					"instanceType": "m4.xlarge"
				}
			]
		},
		"cluster": {
			"calico": {
				"cidr": 16,
				"mtu": 1430,
				"subnet": "192.168.0.0"
			},
			"customer": {
				"id": "giantswarm-production"
			},
			"docker": {
				"daemon": {
					"cidr": "172.17.0.1/16"
				}
			},
			"etcd": {
				"altNames": "",
				"domain": "etcd.oby63.k8s.gorilla.eu-central-1.aws.gigantic.io",
				"port": 2379,
				"prefix": "giantswarm.io"
			},
			"id": "oby63",
			"kubernetes": {
				"api": {
					"clusterIPRange": "172.31.0.0/16",
					"domain": "api.oby63.k8s.gorilla.eu-central-1.aws.gigantic.io",
					"securePort": 443
				},
				"cloudProvider": "aws",
				"dns": {
					"ip": "172.31.0.10"
				},
				"domain": "cluster.local",
				"ingressController": {
					"docker": {
						"image": "quay.io/giantswarm/nginx-ingress-controller:0.9.0-beta.11"
					},
					"domain": "ingress.oby63.k8s.gorilla.eu-central-1.aws.gigantic.io",
					"insecurePort": 30010,
					"securePort": 30011,
					"wildcardDomain": "*.oby63.k8s.gorilla.eu-central-1.aws.gigantic.io"
				},
				"kubelet": {
					"altNames": "kubernetes,kubernetes.default,kubernetes.default.svc,kubernetes.default.svc.cluster.local",
					"domain": "worker.oby63.k8s.gorilla.eu-central-1.aws.gigantic.io",
					"labels": "aws-operator.giantswarm.io/version=5.2.0,giantswarm.io/provider=aws",
					"port": 10250
				},
				"networkSetup": {
					"docker": {
						"image": "quay.io/giantswarm/k8s-setup-network-environment:1f4ffc52095ac368847ce3428ea99b257003d9b9"
					}
				}
			},
			"masters": [
				{
					"id": "ggdp7"
				}
			],
			"scaling": {
				"max": 3,
				"min": 3
			},
			"version": "",
			"workers": [
				{
					"id": "j63ya"
				},
				{
					"id": "4p6sr"
				},
				{
					"id": "vck7j"
				}
			]
		},
		"versionBundle": {
			"version": "5.2.0"
		}
	},
	"status": {
		"aws": {
			"autoScalingGroup": {
				"name": ""
			},
			"availabilityZones": [
				{
					"name": "eu-central-1a",
					"subnet": {
						"private": {
							"cidr": "10.1.0.0/26"
						},
						"public": {
							"cidr": "10.1.0.64/26"
						}
					}
				},
				{
					"name": "eu-central-1b",
					"subnet": {
						"private": {
							"cidr": "10.1.0.128/26"
						},
						"public": {
							"cidr": "10.1.0.192/26"
						}
					}
				}
			]
		},
		"cluster": {
			"conditions": [
				{
					"lastTransitionTime": "2019-08-05T10:53:58.368466037Z",
					"status": "True",
					"type": "Created"
				}
			],
			"network": {
				"cidr": "10.1.0.0/24"
			},
			"nodes": [
				{
					"labels": {
						"aws-operator.giantswarm.io/version": "5.2.0",
						"beta.kubernetes.io/arch": "amd64",
						"beta.kubernetes.io/instance-type": "m4.xlarge",
						"beta.kubernetes.io/os": "linux",
						"failure-domain.beta.kubernetes.io/region": "eu-central-1",
						"failure-domain.beta.kubernetes.io/zone": "eu-central-1a",
						"giantswarm.io/provider": "aws",
						"ip": "10.1.0.12",
						"kubernetes.io/arch": "amd64",
						"kubernetes.io/hostname": "ip-10-1-0-12.eu-central-1.compute.internal",
						"kubernetes.io/os": "linux",
						"kubernetes.io/role": "worker",
						"node-role.kubernetes.io/worker": "",
						"node.kubernetes.io/worker": "",
						"role": "worker"
					},
					"lastTransitionTime": "2019-08-05T10:53:55.853266711Z",
					"name": "ip-10-1-0-12.eu-central-1.compute.internal",
					"version": "5.2.0"
				},
				{
					"labels": {
						"aws-operator.giantswarm.io/version": "5.2.0",
						"beta.kubernetes.io/arch": "amd64",
						"beta.kubernetes.io/instance-type": "m4.xlarge",
						"beta.kubernetes.io/os": "linux",
						"failure-domain.beta.kubernetes.io/region": "eu-central-1",
						"failure-domain.beta.kubernetes.io/zone": "eu-central-1b",
						"giantswarm.io/provider": "aws",
						"ip": "10.1.0.172",
						"kubernetes.io/arch": "amd64",
						"kubernetes.io/hostname": "ip-10-1-0-172.eu-central-1.compute.internal",
						"kubernetes.io/os": "linux",
						"kubernetes.io/role": "worker",
						"node-role.kubernetes.io/worker": "",
						"node.kubernetes.io/worker": "",
						"role": "worker"
					},
					"lastTransitionTime": "2019-08-05T10:53:55.853267564Z",
					"name": "ip-10-1-0-172.eu-central-1.compute.internal",
					"version": "5.2.0"
				},
				{
					"labels": {
						"aws-operator.giantswarm.io/version": "5.2.0",
						"beta.kubernetes.io/arch": "amd64",
						"beta.kubernetes.io/instance-type": "m4.xlarge",
						"beta.kubernetes.io/os": "linux",
						"failure-domain.beta.kubernetes.io/region": "eu-central-1",
						"failure-domain.beta.kubernetes.io/zone": "eu-central-1a",
						"giantswarm.io/provider": "aws",
						"ip": "10.1.0.37",
						"kubernetes.io/arch": "amd64",
						"kubernetes.io/hostname": "ip-10-1-0-37.eu-central-1.compute.internal",
						"kubernetes.io/os": "linux",
						"kubernetes.io/role": "worker",
						"node-role.kubernetes.io/worker": "",
						"node.kubernetes.io/worker": "",
						"role": "worker"
					},
					"lastTransitionTime": "2019-08-05T10:53:55.853268198Z",
					"name": "ip-10-1-0-37.eu-central-1.compute.internal",
					"version": "5.2.0"
				},
				{
					"labels": {
						"aws-operator.giantswarm.io/version": "5.2.0",
						"beta.kubernetes.io/arch": "amd64",
						"beta.kubernetes.io/instance-type": "m4.large",
						"beta.kubernetes.io/os": "linux",
						"failure-domain.beta.kubernetes.io/region": "eu-central-1",
						"failure-domain.beta.kubernetes.io/zone": "eu-central-1a",
						"giantswarm.io/provider": "aws",
						"ip": "10.1.0.48",
						"kubernetes.io/arch": "amd64",
						"kubernetes.io/hostname": "ip-10-1-0-48.eu-central-1.compute.internal",
						"kubernetes.io/os": "linux",
						"kubernetes.io/role": "master",
						"node-role.kubernetes.io/master": "",
						"node.kubernetes.io/master": "",
						"role": "master"
					},
					"lastTransitionTime": "2019-08-05T10:53:55.853268782Z",
					"name": "ip-10-1-0-48.eu-central-1.compute.internal",
					"version": "5.2.0"
				}
			],
			"resources": null,
			"scaling": {
				"desiredCapacity": 3
			},
			"versions": [
				{
					"date": "0001-01-01T00:00:00Z",
					"lastTransitionTime": "2019-08-05T10:53:59.033798538Z",
					"semver": "5.2.0"
				}
			]
		}
	}
}
`

// AWSHealthyTC is a JSON string representing a healthy AWS tenant cluster's node list.
var AWSHealthyTC = `
{
  "kind": "NodeList",
  "apiVersion": "v1",
  "metadata": {
    "selfLink": "/api/v1/nodes",
    "resourceVersion": "50349"
  },
  "items": [
    {
      "metadata": {
        "name": "ip-10-1-1-104.eu-central-1.compute.internal",
        "selfLink": "/api/v1/nodes/ip-10-1-1-104.eu-central-1.compute.internal",
        "uid": "35c333ba-e10e-11e9-b8a6-0ab414f5eb72",
        "resourceVersion": "50304",
        "creationTimestamp": "2019-09-27T10:04:35Z",
        "labels": {
          "aws-operator.giantswarm.io/version": "5.2.1",
          "beta.kubernetes.io/arch": "amd64",
          "beta.kubernetes.io/instance-type": "m4.xlarge",
          "beta.kubernetes.io/os": "linux",
          "failure-domain.beta.kubernetes.io/region": "eu-central-1",
          "failure-domain.beta.kubernetes.io/zone": "eu-central-1c",
          "giantswarm.io/provider": "aws",
          "ip": "10.1.1.104",
          "kubernetes.io/arch": "amd64",
          "kubernetes.io/hostname": "ip-10-1-1-104.eu-central-1.compute.internal",
          "kubernetes.io/os": "linux",
          "kubernetes.io/role": "worker",
          "node-role.kubernetes.io/worker": "",
          "node.kubernetes.io/worker": "",
          "role": "worker"
        },
        "annotations": {
          "node.alpha.kubernetes.io/ttl": "0",
          "volumes.kubernetes.io/controller-managed-attach-detach": "true"
        }
      },
      "spec": {
        "providerID": "aws:///eu-central-1c/i-09367d050fdda6534"
      },
      "status": {
        "capacity": {
          "attachable-volumes-aws-ebs": "39",
          "cpu": "4",
          "ephemeral-storage": "102350Mi",
          "hugepages-1Gi": "0",
          "hugepages-2Mi": "0",
          "memory": "16423784Ki",
          "pods": "110"
        },
        "allocatable": {
          "attachable-volumes-aws-ebs": "39",
          "cpu": "4",
          "ephemeral-storage": "102350Mi",
          "hugepages-1Gi": "0",
          "hugepages-2Mi": "0",
          "memory": "16218984Ki",
          "pods": "110"
        },
        "conditions": [
          {
            "type": "NetworkUnavailable",
            "status": "False",
            "lastHeartbeatTime": "2019-09-27T10:07:59Z",
            "lastTransitionTime": "2019-09-27T10:07:59Z",
            "reason": "CalicoIsUp",
            "message": "Calico is running on this node"
          },
          {
            "type": "MemoryPressure",
            "status": "False",
            "lastHeartbeatTime": "2019-09-27T15:56:20Z",
            "lastTransitionTime": "2019-09-27T10:04:35Z",
            "reason": "KubeletHasSufficientMemory",
            "message": "kubelet has sufficient memory available"
          },
          {
            "type": "DiskPressure",
            "status": "False",
            "lastHeartbeatTime": "2019-09-27T15:56:20Z",
            "lastTransitionTime": "2019-09-27T10:04:35Z",
            "reason": "KubeletHasNoDiskPressure",
            "message": "kubelet has no disk pressure"
          },
          {
            "type": "PIDPressure",
            "status": "False",
            "lastHeartbeatTime": "2019-09-27T15:56:20Z",
            "lastTransitionTime": "2019-09-27T10:04:35Z",
            "reason": "KubeletHasSufficientPID",
            "message": "kubelet has sufficient PID available"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastHeartbeatTime": "2019-09-27T15:56:20Z",
            "lastTransitionTime": "2019-09-27T10:05:45Z",
            "reason": "KubeletReady",
            "message": "kubelet is posting ready status"
          }
        ],
        "addresses": [
          {
            "type": "InternalIP",
            "address": "10.1.1.104"
          },
          {
            "type": "InternalDNS",
            "address": "ip-10-1-1-104.eu-central-1.compute.internal"
          },
          {
            "type": "Hostname",
            "address": "ip-10-1-1-104.eu-central-1.compute.internal"
          }
        ],
        "daemonEndpoints": {
          "kubeletEndpoint": {
            "Port": 10250
          }
        },
        "nodeInfo": {
          "machineID": "44a4d2157ac64a7db47d907162c4b6d4",
          "systemUUID": "ec242200-f543-30af-ac10-d9eecad03ff6",
          "bootID": "7cacf390-441a-48dc-ba94-41a256b490c9",
          "kernelVersion": "4.19.50-coreos-r1",
          "osImage": "Container Linux by CoreOS 2135.4.0 (Rhyolite)",
          "containerRuntimeVersion": "docker://18.6.3",
          "kubeletVersion": "v1.14.5",
          "kubeProxyVersion": "v1.14.5",
          "operatingSystem": "linux",
          "architecture": "amd64"
        },
        "images": [
          {
            "names": [
              "quay.io/giantswarm/hyperkube@sha256:36b72fd47d1c9e4fab2cf2375eb91b791a161b8cdbd39152c23efab389df7c78",
              "quay.io/giantswarm/hyperkube:v1.14.5"
            ],
            "sizeBytes": 603981854
          },
          {
            "names": [
              "quay.io/giantswarm/nginx-ingress-controller@sha256:8d95766156584273b74d1b315530428eadae795c3bdbff05e6775023e7b380eb",
              "quay.io/giantswarm/nginx-ingress-controller:0.25.1"
            ],
            "sizeBytes": 510687563
          },
          {
            "names": [
              "quay.io/giantswarm/node@sha256:5077abf4b7757f3d6c50bbacb57877a6d2d980e3889150c6df23c39666072e80",
              "quay.io/giantswarm/node:v3.7.2"
            ],
            "sizeBytes": 156259143
          },
          {
            "names": [
              "quay.io/giantswarm/cni@sha256:bbe61ad6aae9de719d70c23d8195704a4d978924b7956b22cc86eeab2d517d65",
              "quay.io/giantswarm/cni:v3.7.2"
            ],
            "sizeBytes": 135366007
          },
          {
            "names": [
              "quay.io/giantswarm/metrics-server-amd64@sha256:7e70cf062bafd653e2ce2305edc9fcb3369d097911006afff1e16aa1413297ae",
              "quay.io/giantswarm/metrics-server-amd64:v0.3.3"
            ],
            "sizeBytes": 39933796
          },
          {
            "names": [
              "quay.io/giantswarm/addon-resizer@sha256:9cf17b7c7412d9356dd0fef3c8401622060612d58ca8d6651e121eccc12f810c",
              "quay.io/giantswarm/addon-resizer:1.8.4"
            ],
            "sizeBytes": 38349857
          },
          {
            "names": [
              "quay.io/giantswarm/namespace-labeler@sha256:e451eacf65af560bc103ed8ac317032016952b6bed95b8fabb7884b5d9f041ce",
              "quay.io/giantswarm/namespace-labeler:latest"
            ],
            "sizeBytes": 37594320
          },
          {
            "names": [
              "quay.io/giantswarm/kube-state-metrics@sha256:c6c5033bba9af1e2e78346011392406f1a3904df296ca75de751ffe0f738220c",
              "quay.io/giantswarm/kube-state-metrics:v1.6.0"
            ],
            "sizeBytes": 35535767
          },
          {
            "names": [
              "quay.io/giantswarm/net-exporter@sha256:e6c1932cb38b62d789e21defcd5b4c8f21b555890517f9a5f7e4511aa3b7edee",
              "quay.io/giantswarm/net-exporter:9476dceada4f4577dd765fb8026d7d9a214800ec"
            ],
            "sizeBytes": 32227567
          },
          {
            "names": [
              "quay.io/giantswarm/node-exporter@sha256:904ce5ad72875696a5a084e8dba10c96b94dc269abda9720729756898b74d881",
              "quay.io/giantswarm/node-exporter:v0.18.0-giantswarm"
            ],
            "sizeBytes": 22891793
          },
          {
            "names": [
              "quay.io/giantswarm/cert-exporter@sha256:2222ed4ffb5e7404c7eb8f02f25ae1ffda66abb1955b649f747929f9f50d2194",
              "quay.io/giantswarm/cert-exporter:228632041f419d45c1aadf7a713a4b3ab697abcb"
            ],
            "sizeBytes": 15307207
          },
          {
            "names": [
              "quay.io/giantswarm/k8s-setup-network-environment@sha256:e337d03e569e53b246f4dea91359efbabe7b3ddc78878e1875d0c7aaf0e17fd5",
              "quay.io/giantswarm/k8s-setup-network-environment:1f4ffc52095ac368847ce3428ea99b257003d9b9"
            ],
            "sizeBytes": 7882969
          },
          {
            "names": [
              "k8s.gcr.io/pause@sha256:f78411e19d84a252e53bff71a4407a5686c46983a2c2eeed83929b888179acea",
              "k8s.gcr.io/pause:3.1"
            ],
            "sizeBytes": 742472
          }
        ]
      }
    },
    {
      "metadata": {
        "name": "ip-10-1-1-125.eu-central-1.compute.internal",
        "selfLink": "/api/v1/nodes/ip-10-1-1-125.eu-central-1.compute.internal",
        "uid": "350b4df4-e10e-11e9-b8a6-0ab414f5eb72",
        "resourceVersion": "50297",
        "creationTimestamp": "2019-09-27T10:04:34Z",
        "labels": {
          "aws-operator.giantswarm.io/version": "5.2.1",
          "beta.kubernetes.io/arch": "amd64",
          "beta.kubernetes.io/instance-type": "m4.xlarge",
          "beta.kubernetes.io/os": "linux",
          "failure-domain.beta.kubernetes.io/region": "eu-central-1",
          "failure-domain.beta.kubernetes.io/zone": "eu-central-1c",
          "giantswarm.io/provider": "aws",
          "ip": "10.1.1.125",
          "kubernetes.io/arch": "amd64",
          "kubernetes.io/hostname": "ip-10-1-1-125.eu-central-1.compute.internal",
          "kubernetes.io/os": "linux",
          "kubernetes.io/role": "worker",
          "node-role.kubernetes.io/worker": "",
          "node.kubernetes.io/worker": "",
          "role": "worker"
        },
        "annotations": {
          "node.alpha.kubernetes.io/ttl": "0",
          "volumes.kubernetes.io/controller-managed-attach-detach": "true"
        }
      },
      "spec": {
        "providerID": "aws:///eu-central-1c/i-0ef4543312ea00a9b"
      },
      "status": {
        "capacity": {
          "attachable-volumes-aws-ebs": "39",
          "cpu": "4",
          "ephemeral-storage": "102350Mi",
          "hugepages-1Gi": "0",
          "hugepages-2Mi": "0",
          "memory": "16423784Ki",
          "pods": "110"
        },
        "allocatable": {
          "attachable-volumes-aws-ebs": "39",
          "cpu": "4",
          "ephemeral-storage": "102350Mi",
          "hugepages-1Gi": "0",
          "hugepages-2Mi": "0",
          "memory": "16218984Ki",
          "pods": "110"
        },
        "conditions": [
          {
            "type": "NetworkUnavailable",
            "status": "False",
            "lastHeartbeatTime": "2019-09-27T10:08:40Z",
            "lastTransitionTime": "2019-09-27T10:08:40Z",
            "reason": "CalicoIsUp",
            "message": "Calico is running on this node"
          },
          {
            "type": "MemoryPressure",
            "status": "False",
            "lastHeartbeatTime": "2019-09-27T15:56:18Z",
            "lastTransitionTime": "2019-09-27T10:04:34Z",
            "reason": "KubeletHasSufficientMemory",
            "message": "kubelet has sufficient memory available"
          },
          {
            "type": "DiskPressure",
            "status": "False",
            "lastHeartbeatTime": "2019-09-27T15:56:18Z",
            "lastTransitionTime": "2019-09-27T10:04:34Z",
            "reason": "KubeletHasNoDiskPressure",
            "message": "kubelet has no disk pressure"
          },
          {
            "type": "PIDPressure",
            "status": "False",
            "lastHeartbeatTime": "2019-09-27T15:56:18Z",
            "lastTransitionTime": "2019-09-27T10:04:34Z",
            "reason": "KubeletHasSufficientPID",
            "message": "kubelet has sufficient PID available"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastHeartbeatTime": "2019-09-27T15:56:18Z",
            "lastTransitionTime": "2019-09-27T10:05:44Z",
            "reason": "KubeletReady",
            "message": "kubelet is posting ready status"
          }
        ],
        "addresses": [
          {
            "type": "InternalIP",
            "address": "10.1.1.125"
          },
          {
            "type": "InternalDNS",
            "address": "ip-10-1-1-125.eu-central-1.compute.internal"
          },
          {
            "type": "Hostname",
            "address": "ip-10-1-1-125.eu-central-1.compute.internal"
          }
        ],
        "daemonEndpoints": {
          "kubeletEndpoint": {
            "Port": 10250
          }
        },
        "nodeInfo": {
          "machineID": "1732f81da0b842b0a97e6e4369e0fd20",
          "systemUUID": "ec25543a-104e-23a2-c3a4-f62cfefadb09",
          "bootID": "05d7e166-f141-4959-84a9-56c1208e6593",
          "kernelVersion": "4.19.50-coreos-r1",
          "osImage": "Container Linux by CoreOS 2135.4.0 (Rhyolite)",
          "containerRuntimeVersion": "docker://18.6.3",
          "kubeletVersion": "v1.14.5",
          "kubeProxyVersion": "v1.14.5",
          "operatingSystem": "linux",
          "architecture": "amd64"
        },
        "images": [
          {
            "names": [
              "quay.io/giantswarm/hyperkube@sha256:36b72fd47d1c9e4fab2cf2375eb91b791a161b8cdbd39152c23efab389df7c78",
              "quay.io/giantswarm/hyperkube:v1.14.5"
            ],
            "sizeBytes": 603981854
          },
          {
            "names": [
              "quay.io/giantswarm/nginx-ingress-controller@sha256:8d95766156584273b74d1b315530428eadae795c3bdbff05e6775023e7b380eb",
              "quay.io/giantswarm/nginx-ingress-controller:0.25.1"
            ],
            "sizeBytes": 510687563
          },
          {
            "names": [
              "quay.io/giantswarm/node@sha256:5077abf4b7757f3d6c50bbacb57877a6d2d980e3889150c6df23c39666072e80",
              "quay.io/giantswarm/node:v3.7.2"
            ],
            "sizeBytes": 156259143
          },
          {
            "names": [
              "quay.io/giantswarm/cni@sha256:bbe61ad6aae9de719d70c23d8195704a4d978924b7956b22cc86eeab2d517d65",
              "quay.io/giantswarm/cni:v3.7.2"
            ],
            "sizeBytes": 135366007
          },
          {
            "names": [
              "quay.io/giantswarm/coredns@sha256:a0818efc411db0a1e53bc63756d558d3bc7d75f2f10bb19d28837b7e9badd071",
              "quay.io/giantswarm/coredns:1.5.1"
            ],
            "sizeBytes": 40822376
          },
          {
            "names": [
              "quay.io/giantswarm/addon-resizer@sha256:9cf17b7c7412d9356dd0fef3c8401622060612d58ca8d6651e121eccc12f810c",
              "quay.io/giantswarm/addon-resizer:1.8.4"
            ],
            "sizeBytes": 38349857
          },
          {
            "names": [
              "quay.io/giantswarm/namespace-labeler@sha256:e451eacf65af560bc103ed8ac317032016952b6bed95b8fabb7884b5d9f041ce",
              "quay.io/giantswarm/namespace-labeler:latest"
            ],
            "sizeBytes": 37594320
          },
          {
            "names": [
              "quay.io/giantswarm/kube-state-metrics@sha256:c6c5033bba9af1e2e78346011392406f1a3904df296ca75de751ffe0f738220c",
              "quay.io/giantswarm/kube-state-metrics:v1.6.0"
            ],
            "sizeBytes": 35535767
          },
          {
            "names": [
              "quay.io/giantswarm/net-exporter@sha256:e6c1932cb38b62d789e21defcd5b4c8f21b555890517f9a5f7e4511aa3b7edee",
              "quay.io/giantswarm/net-exporter:9476dceada4f4577dd765fb8026d7d9a214800ec"
            ],
            "sizeBytes": 32227567
          },
          {
            "names": [
              "quay.io/giantswarm/node-exporter@sha256:904ce5ad72875696a5a084e8dba10c96b94dc269abda9720729756898b74d881",
              "quay.io/giantswarm/node-exporter:v0.18.0-giantswarm"
            ],
            "sizeBytes": 22891793
          },
          {
            "names": [
              "quay.io/giantswarm/cert-exporter@sha256:2222ed4ffb5e7404c7eb8f02f25ae1ffda66abb1955b649f747929f9f50d2194",
              "quay.io/giantswarm/cert-exporter:228632041f419d45c1aadf7a713a4b3ab697abcb"
            ],
            "sizeBytes": 15307207
          },
          {
            "names": [
              "quay.io/giantswarm/k8s-setup-network-environment@sha256:e337d03e569e53b246f4dea91359efbabe7b3ddc78878e1875d0c7aaf0e17fd5",
              "quay.io/giantswarm/k8s-setup-network-environment:1f4ffc52095ac368847ce3428ea99b257003d9b9"
            ],
            "sizeBytes": 7882969
          },
          {
            "names": [
              "k8s.gcr.io/pause@sha256:f78411e19d84a252e53bff71a4407a5686c46983a2c2eeed83929b888179acea",
              "k8s.gcr.io/pause:3.1"
            ],
            "sizeBytes": 742472
          }
        ]
      }
    },
    {
      "metadata": {
        "name": "ip-10-1-1-57.eu-central-1.compute.internal",
        "selfLink": "/api/v1/nodes/ip-10-1-1-57.eu-central-1.compute.internal",
        "uid": "317f139b-e10e-11e9-b8a6-0ab414f5eb72",
        "resourceVersion": "50247",
        "creationTimestamp": "2019-09-27T10:04:28Z",
        "labels": {
          "aws-operator.giantswarm.io/version": "5.2.1",
          "beta.kubernetes.io/arch": "amd64",
          "beta.kubernetes.io/instance-type": "m4.xlarge",
          "beta.kubernetes.io/os": "linux",
          "failure-domain.beta.kubernetes.io/region": "eu-central-1",
          "failure-domain.beta.kubernetes.io/zone": "eu-central-1c",
          "giantswarm.io/provider": "aws",
          "ip": "10.1.1.57",
          "kubernetes.io/arch": "amd64",
          "kubernetes.io/hostname": "ip-10-1-1-57.eu-central-1.compute.internal",
          "kubernetes.io/os": "linux",
          "kubernetes.io/role": "master",
          "node-role.kubernetes.io/master": "",
          "node.kubernetes.io/master": "",
          "role": "master"
        },
        "annotations": {
          "node.alpha.kubernetes.io/ttl": "0",
          "volumes.kubernetes.io/controller-managed-attach-detach": "true"
        }
      },
      "spec": {
        "providerID": "aws:///eu-central-1c/i-0d655433a6c4c81cd",
        "taints": [
          {
            "key": "node-role.kubernetes.io/master",
            "effect": "NoSchedule"
          }
        ]
      },
      "status": {
        "capacity": {
          "attachable-volumes-aws-ebs": "39",
          "cpu": "4",
          "ephemeral-storage": "5706380Ki",
          "hugepages-1Gi": "0",
          "hugepages-2Mi": "0",
          "memory": "16423784Ki",
          "pods": "110"
        },
        "allocatable": {
          "attachable-volumes-aws-ebs": "39",
          "cpu": "4",
          "ephemeral-storage": "5706380Ki",
          "hugepages-1Gi": "0",
          "hugepages-2Mi": "0",
          "memory": "16218984Ki",
          "pods": "110"
        },
        "conditions": [
          {
            "type": "NetworkUnavailable",
            "status": "False",
            "lastHeartbeatTime": "2019-09-27T10:08:20Z",
            "lastTransitionTime": "2019-09-27T10:08:20Z",
            "reason": "CalicoIsUp",
            "message": "Calico is running on this node"
          },
          {
            "type": "MemoryPressure",
            "status": "False",
            "lastHeartbeatTime": "2019-09-27T15:55:56Z",
            "lastTransitionTime": "2019-09-27T10:04:28Z",
            "reason": "KubeletHasSufficientMemory",
            "message": "kubelet has sufficient memory available"
          },
          {
            "type": "DiskPressure",
            "status": "False",
            "lastHeartbeatTime": "2019-09-27T15:55:56Z",
            "lastTransitionTime": "2019-09-27T10:04:28Z",
            "reason": "KubeletHasNoDiskPressure",
            "message": "kubelet has no disk pressure"
          },
          {
            "type": "PIDPressure",
            "status": "False",
            "lastHeartbeatTime": "2019-09-27T15:55:56Z",
            "lastTransitionTime": "2019-09-27T10:04:28Z",
            "reason": "KubeletHasSufficientPID",
            "message": "kubelet has sufficient PID available"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastHeartbeatTime": "2019-09-27T15:55:56Z",
            "lastTransitionTime": "2019-09-27T10:06:17Z",
            "reason": "KubeletReady",
            "message": "kubelet is posting ready status"
          }
        ],
        "addresses": [
          {
            "type": "InternalIP",
            "address": "10.1.1.57"
          },
          {
            "type": "InternalDNS",
            "address": "ip-10-1-1-57.eu-central-1.compute.internal"
          },
          {
            "type": "Hostname",
            "address": "ip-10-1-1-57.eu-central-1.compute.internal"
          }
        ],
        "daemonEndpoints": {
          "kubeletEndpoint": {
            "Port": 10250
          }
        },
        "nodeInfo": {
          "machineID": "7da794b2687c4e1ca9ade9c7e2c466cd",
          "systemUUID": "ec253546-8f3c-509f-fe2f-a08573f388a1",
          "bootID": "6a322db2-b4d1-4b01-9af0-cc4b176a595a",
          "kernelVersion": "4.19.50-coreos-r1",
          "osImage": "Container Linux by CoreOS 2135.4.0 (Rhyolite)",
          "containerRuntimeVersion": "docker://18.6.3",
          "kubeletVersion": "v1.14.5",
          "kubeProxyVersion": "v1.14.5",
          "operatingSystem": "linux",
          "architecture": "amd64"
        },
        "images": [
          {
            "names": [
              "quay.io/giantswarm/hyperkube@sha256:36b72fd47d1c9e4fab2cf2375eb91b791a161b8cdbd39152c23efab389df7c78",
              "quay.io/giantswarm/hyperkube:v1.14.5"
            ],
            "sizeBytes": 603981854
          },
          {
            "names": [
              "quay.io/giantswarm/node@sha256:5077abf4b7757f3d6c50bbacb57877a6d2d980e3889150c6df23c39666072e80",
              "quay.io/giantswarm/node:v3.7.2"
            ],
            "sizeBytes": 156259143
          },
          {
            "names": [
              "quay.io/giantswarm/cluster-autoscaler@sha256:7a2879059e883417736c659fe985ce36cd8525cd776c96d77931e7a6e07074c2",
              "quay.io/giantswarm/cluster-autoscaler:v1.14.3"
            ],
            "sizeBytes": 142108093
          },
          {
            "names": [
              "quay.io/giantswarm/cni@sha256:bbe61ad6aae9de719d70c23d8195704a4d978924b7956b22cc86eeab2d517d65",
              "quay.io/giantswarm/cni:v3.7.2"
            ],
            "sizeBytes": 135366007
          },
          {
            "names": [
              "quay.io/giantswarm/docker-kubectl@sha256:93404a39c4b83568aef2971504fecc39efc6a9b949fd84cc9814742a76033ed3",
              "quay.io/giantswarm/docker-kubectl:f5cae44c480bd797dc770dd5f62d40b74063c0d7"
            ],
            "sizeBytes": 64559926
          },
          {
            "names": [
              "quay.io/giantswarm/chart-operator@sha256:f210c32dc89ba6bd6b6d3429da53b25f37e6a743827cf1d24c907e1b12140855",
              "quay.io/giantswarm/chart-operator:ea0d7e1eadc6f62ee2c48d239c64519aa57c0864"
            ],
            "sizeBytes": 49392950
          },
          {
            "names": [
              "quay.io/giantswarm/kube-controllers@sha256:0bd19e09717c61858f0d89b9817fd31bfd0e5a33c2e94c6d639f532f3e2c695e",
              "quay.io/giantswarm/kube-controllers:v3.7.2"
            ],
            "sizeBytes": 46774220
          },
          {
            "names": [
              "quay.io/giantswarm/namespace-labeler@sha256:e451eacf65af560bc103ed8ac317032016952b6bed95b8fabb7884b5d9f041ce",
              "quay.io/giantswarm/namespace-labeler:latest"
            ],
            "sizeBytes": 37594320
          },
          {
            "names": [
              "quay.io/giantswarm/etcd@sha256:6d71e368c3f98fe6745a7edf8e7828ffeacff199cf2fb9db645a1af825d4284c",
              "quay.io/giantswarm/etcd:v3.3.13"
            ],
            "sizeBytes": 35959206
          },
          {
            "names": [
              "quay.io/giantswarm/net-exporter@sha256:e6c1932cb38b62d789e21defcd5b4c8f21b555890517f9a5f7e4511aa3b7edee",
              "quay.io/giantswarm/net-exporter:9476dceada4f4577dd765fb8026d7d9a214800ec"
            ],
            "sizeBytes": 32227567
          },
          {
            "names": [
              "quay.io/giantswarm/node-exporter@sha256:904ce5ad72875696a5a084e8dba10c96b94dc269abda9720729756898b74d881",
              "quay.io/giantswarm/node-exporter:v0.18.0-giantswarm"
            ],
            "sizeBytes": 22891793
          },
          {
            "names": [
              "quay.io/giantswarm/cert-exporter@sha256:2222ed4ffb5e7404c7eb8f02f25ae1ffda66abb1955b649f747929f9f50d2194",
              "quay.io/giantswarm/cert-exporter:228632041f419d45c1aadf7a713a4b3ab697abcb"
            ],
            "sizeBytes": 15307207
          },
          {
            "names": [
              "quay.io/giantswarm/k8s-setup-network-environment@sha256:e337d03e569e53b246f4dea91359efbabe7b3ddc78878e1875d0c7aaf0e17fd5",
              "quay.io/giantswarm/k8s-setup-network-environment:1f4ffc52095ac368847ce3428ea99b257003d9b9"
            ],
            "sizeBytes": 7882969
          },
          {
            "names": [
              "k8s.gcr.io/pause@sha256:f78411e19d84a252e53bff71a4407a5686c46983a2c2eeed83929b888179acea",
              "k8s.gcr.io/pause:3.1"
            ],
            "sizeBytes": 742472
          }
        ]
      }
    },
    {
      "metadata": {
        "name": "ip-10-1-1-85.eu-central-1.compute.internal",
        "selfLink": "/api/v1/nodes/ip-10-1-1-85.eu-central-1.compute.internal",
        "uid": "352b71b7-e10e-11e9-b8a6-0ab414f5eb72",
        "resourceVersion": "50296",
        "creationTimestamp": "2019-09-27T10:04:34Z",
        "labels": {
          "aws-operator.giantswarm.io/version": "5.2.1",
          "beta.kubernetes.io/arch": "amd64",
          "beta.kubernetes.io/instance-type": "m4.xlarge",
          "beta.kubernetes.io/os": "linux",
          "failure-domain.beta.kubernetes.io/region": "eu-central-1",
          "failure-domain.beta.kubernetes.io/zone": "eu-central-1c",
          "giantswarm.io/provider": "aws",
          "ip": "10.1.1.85",
          "kubernetes.io/arch": "amd64",
          "kubernetes.io/hostname": "ip-10-1-1-85.eu-central-1.compute.internal",
          "kubernetes.io/os": "linux",
          "kubernetes.io/role": "worker",
          "node-role.kubernetes.io/worker": "",
          "node.kubernetes.io/worker": "",
          "role": "worker"
        },
        "annotations": {
          "node.alpha.kubernetes.io/ttl": "0",
          "volumes.kubernetes.io/controller-managed-attach-detach": "true"
        }
      },
      "spec": {
        "providerID": "aws:///eu-central-1c/i-0eecd28fd1bbef822"
      },
      "status": {
        "capacity": {
          "attachable-volumes-aws-ebs": "39",
          "cpu": "4",
          "ephemeral-storage": "102350Mi",
          "hugepages-1Gi": "0",
          "hugepages-2Mi": "0",
          "memory": "16423784Ki",
          "pods": "110"
        },
        "allocatable": {
          "attachable-volumes-aws-ebs": "39",
          "cpu": "4",
          "ephemeral-storage": "102350Mi",
          "hugepages-1Gi": "0",
          "hugepages-2Mi": "0",
          "memory": "16218984Ki",
          "pods": "110"
        },
        "conditions": [
          {
            "type": "NetworkUnavailable",
            "status": "False",
            "lastHeartbeatTime": "2019-09-27T10:07:37Z",
            "lastTransitionTime": "2019-09-27T10:07:37Z",
            "reason": "CalicoIsUp",
            "message": "Calico is running on this node"
          },
          {
            "type": "MemoryPressure",
            "status": "False",
            "lastHeartbeatTime": "2019-09-27T15:56:18Z",
            "lastTransitionTime": "2019-09-27T10:04:34Z",
            "reason": "KubeletHasSufficientMemory",
            "message": "kubelet has sufficient memory available"
          },
          {
            "type": "DiskPressure",
            "status": "False",
            "lastHeartbeatTime": "2019-09-27T15:56:18Z",
            "lastTransitionTime": "2019-09-27T10:04:34Z",
            "reason": "KubeletHasNoDiskPressure",
            "message": "kubelet has no disk pressure"
          },
          {
            "type": "PIDPressure",
            "status": "False",
            "lastHeartbeatTime": "2019-09-27T15:56:18Z",
            "lastTransitionTime": "2019-09-27T10:04:34Z",
            "reason": "KubeletHasSufficientPID",
            "message": "kubelet has sufficient PID available"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastHeartbeatTime": "2019-09-27T15:56:18Z",
            "lastTransitionTime": "2019-09-27T10:05:44Z",
            "reason": "KubeletReady",
            "message": "kubelet is posting ready status"
          }
        ],
        "addresses": [
          {
            "type": "InternalIP",
            "address": "10.1.1.85"
          },
          {
            "type": "InternalDNS",
            "address": "ip-10-1-1-85.eu-central-1.compute.internal"
          },
          {
            "type": "Hostname",
            "address": "ip-10-1-1-85.eu-central-1.compute.internal"
          }
        ],
        "daemonEndpoints": {
          "kubeletEndpoint": {
            "Port": 10250
          }
        },
        "nodeInfo": {
          "machineID": "002aefbdcd8942f6994d3973dfd6ec17",
          "systemUUID": "ec2a52ed-ee67-fc72-e429-e69c3b46543d",
          "bootID": "d1c3fddc-a81b-4aa1-bb89-ed28b7feed1e",
          "kernelVersion": "4.19.50-coreos-r1",
          "osImage": "Container Linux by CoreOS 2135.4.0 (Rhyolite)",
          "containerRuntimeVersion": "docker://18.6.3",
          "kubeletVersion": "v1.14.5",
          "kubeProxyVersion": "v1.14.5",
          "operatingSystem": "linux",
          "architecture": "amd64"
        },
        "images": [
          {
            "names": [
              "quay.io/giantswarm/hyperkube@sha256:36b72fd47d1c9e4fab2cf2375eb91b791a161b8cdbd39152c23efab389df7c78",
              "quay.io/giantswarm/hyperkube:v1.14.5"
            ],
            "sizeBytes": 603981854
          },
          {
            "names": [
              "quay.io/giantswarm/nginx-ingress-controller@sha256:8d95766156584273b74d1b315530428eadae795c3bdbff05e6775023e7b380eb",
              "quay.io/giantswarm/nginx-ingress-controller:0.25.1"
            ],
            "sizeBytes": 510687563
          },
          {
            "names": [
              "quay.io/giantswarm/node@sha256:5077abf4b7757f3d6c50bbacb57877a6d2d980e3889150c6df23c39666072e80",
              "quay.io/giantswarm/node:v3.7.2"
            ],
            "sizeBytes": 156259143
          },
          {
            "names": [
              "quay.io/giantswarm/cni@sha256:bbe61ad6aae9de719d70c23d8195704a4d978924b7956b22cc86eeab2d517d65",
              "quay.io/giantswarm/cni:v3.7.2"
            ],
            "sizeBytes": 135366007
          },
          {
            "names": [
              "quay.io/giantswarm/tiller@sha256:3d632052b15d7b39204ea952b085a28acd74f8cd96817e63b1d7221268aac6e6",
              "quay.io/giantswarm/tiller:v2.12.0"
            ],
            "sizeBytes": 81345536
          },
          {
            "names": [
              "quay.io/giantswarm/coredns@sha256:a0818efc411db0a1e53bc63756d558d3bc7d75f2f10bb19d28837b7e9badd071",
              "quay.io/giantswarm/coredns:1.5.1"
            ],
            "sizeBytes": 40822376
          },
          {
            "names": [
              "quay.io/giantswarm/namespace-labeler@sha256:e451eacf65af560bc103ed8ac317032016952b6bed95b8fabb7884b5d9f041ce",
              "quay.io/giantswarm/namespace-labeler:latest"
            ],
            "sizeBytes": 37594320
          },
          {
            "names": [
              "quay.io/giantswarm/net-exporter@sha256:e6c1932cb38b62d789e21defcd5b4c8f21b555890517f9a5f7e4511aa3b7edee",
              "quay.io/giantswarm/net-exporter:9476dceada4f4577dd765fb8026d7d9a214800ec"
            ],
            "sizeBytes": 32227567
          },
          {
            "names": [
              "quay.io/giantswarm/node-exporter@sha256:904ce5ad72875696a5a084e8dba10c96b94dc269abda9720729756898b74d881",
              "quay.io/giantswarm/node-exporter:v0.18.0-giantswarm"
            ],
            "sizeBytes": 22891793
          },
          {
            "names": [
              "quay.io/giantswarm/cert-exporter@sha256:2222ed4ffb5e7404c7eb8f02f25ae1ffda66abb1955b649f747929f9f50d2194",
              "quay.io/giantswarm/cert-exporter:228632041f419d45c1aadf7a713a4b3ab697abcb"
            ],
            "sizeBytes": 15307207
          },
          {
            "names": [
              "quay.io/giantswarm/k8s-setup-network-environment@sha256:e337d03e569e53b246f4dea91359efbabe7b3ddc78878e1875d0c7aaf0e17fd5",
              "quay.io/giantswarm/k8s-setup-network-environment:1f4ffc52095ac368847ce3428ea99b257003d9b9"
            ],
            "sizeBytes": 7882969
          },
          {
            "names": [
              "k8s.gcr.io/pause@sha256:f78411e19d84a252e53bff71a4407a5686c46983a2c2eeed83929b888179acea",
              "k8s.gcr.io/pause:3.1"
            ],
            "sizeBytes": 742472
          }
        ]
      }
    }
  ]
}
`
