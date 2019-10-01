package mock

// AzureHealthy is a JSON string representing a healthy Azure cluster
var AzureHealthy = `
{
  "apiVersion": "provider.giantswarm.io/v1alpha1",
  "kind": "AzureConfig",
  "metadata": {
    "creationTimestamp": "2019-09-25T08:54:44Z",
    "name": "0cu4f",
    "namespace": "default"
  },
  "spec": {
    "azure": {
      "masters": [
        {
          "dockerVolumeSizeGB": 0,
          "vmSize": "Standard_D2s_v3"
        }
      ],
      "virtualNetwork": {
        "calicoSubnetCIDR": "10.5.128.0/17",
        "cidr": "10.5.0.0/16",
        "masterSubnetCIDR": "10.5.0.0/24",
        "workerSubnetCIDR": "10.5.1.0/24"
      },
      "workers": [
        {
          "dockerVolumeSizeGB": 0,
          "vmSize": "Standard_D2s_v3"
        },
        {
          "dockerVolumeSizeGB": 0,
          "vmSize": "Standard_D2s_v3"
        },
        {
          "dockerVolumeSizeGB": 0,
          "vmSize": "Standard_D2s_v3"
        }
      ]
    },
    "cluster": {
      "calico": {
        "cidr": 0,
        "mtu": 1430,
        "subnet": ""
      },
      "customer": {
        "id": "giantswarm"
      },
      "docker": {
        "daemon": {
          "cidr": "172.17.0.1/16"
        }
      },
      "id": "0cu4f",
      "masters": [
        {
          "id": "cu4pk"
        }
      ],
      "scaling": {
        "max": 3,
        "min": 3
      },
      "version": "",
      "workers": [
        {
          "id": "b65zp"
        },
        {
          "id": "6n50u"
        },
        {
          "id": "py4ji"
        }
      ]
    },
    "versionBundle": {
      "version": "2.7.0"
    }
  },
  "status": {
    "cluster": {
      "conditions": [
        {
          "lastTransitionTime": "2019-09-25T09:11:09.628703797Z",
          "status": "True",
          "type": "Created"
        }
      ],
      "network": {
        "cidr": ""
      },
      "nodes": [
        {
          "labels": {
            "azure-operator.giantswarm.io/version": "2.7.0",
            "beta.kubernetes.io/arch": "amd64",
            "beta.kubernetes.io/instance-type": "Standard_D2s_v3",
            "beta.kubernetes.io/os": "linux",
            "failure-domain.beta.kubernetes.io/region": "westeurope",
            "failure-domain.beta.kubernetes.io/zone": "0",
            "giantswarm.io/provider": "azure",
            "ip": "10.5.0.5",
            "kubernetes.io/arch": "amd64",
            "kubernetes.io/hostname": "0cu4f-master-000000",
            "kubernetes.io/os": "linux",
            "kubernetes.io/role": "master",
            "node-role.kubernetes.io/master": "",
            "node.kubernetes.io/master": "",
            "role": "master"
          },
          "lastTransitionTime": "2019-09-25T09:07:50.224012229Z",
          "name": "0cu4f-master-000000",
          "version": "2.7.0"
        },
        {
          "labels": {
            "azure-operator.giantswarm.io/version": "2.7.0",
            "beta.kubernetes.io/arch": "amd64",
            "beta.kubernetes.io/instance-type": "Standard_D2s_v3",
            "beta.kubernetes.io/os": "linux",
            "failure-domain.beta.kubernetes.io/region": "westeurope",
            "failure-domain.beta.kubernetes.io/zone": "0",
            "giantswarm.io/provider": "azure",
            "ip": "10.5.1.4",
            "kubernetes.io/arch": "amd64",
            "kubernetes.io/hostname": "0cu4f-worker-000000",
            "kubernetes.io/os": "linux",
            "kubernetes.io/role": "worker",
            "node-role.kubernetes.io/worker": "",
            "node.kubernetes.io/worker": "",
            "role": "worker"
          },
          "lastTransitionTime": "2019-09-25T09:07:50.224013129Z",
          "name": "0cu4f-worker-000000",
          "version": "2.7.0"
        },
        {
          "labels": {
            "azure-operator.giantswarm.io/version": "2.7.0",
            "beta.kubernetes.io/arch": "amd64",
            "beta.kubernetes.io/instance-type": "Standard_D2s_v3",
            "beta.kubernetes.io/os": "linux",
            "failure-domain.beta.kubernetes.io/region": "westeurope",
            "failure-domain.beta.kubernetes.io/zone": "4",
            "giantswarm.io/provider": "azure",
            "ip": "10.5.1.8",
            "kubernetes.io/arch": "amd64",
            "kubernetes.io/hostname": "0cu4f-worker-000004",
            "kubernetes.io/os": "linux",
            "kubernetes.io/role": "worker",
            "node-role.kubernetes.io/worker": "",
            "node.kubernetes.io/worker": "",
            "role": "worker"
          },
          "lastTransitionTime": "2019-09-25T09:07:50.224013929Z",
          "name": "0cu4f-worker-000004",
          "version": "2.7.0"
        },
        {
          "labels": {
            "azure-operator.giantswarm.io/version": "2.7.0",
            "beta.kubernetes.io/arch": "amd64",
            "beta.kubernetes.io/instance-type": "Standard_D2s_v3",
            "beta.kubernetes.io/os": "linux",
            "failure-domain.beta.kubernetes.io/region": "westeurope",
            "failure-domain.beta.kubernetes.io/zone": "1",
            "giantswarm.io/provider": "azure",
            "ip": "10.5.1.10",
            "kubernetes.io/arch": "amd64",
            "kubernetes.io/hostname": "0cu4f-worker-000006",
            "kubernetes.io/os": "linux",
            "kubernetes.io/role": "worker",
            "node-role.kubernetes.io/worker": "",
            "node.kubernetes.io/worker": "",
            "role": "worker"
          },
          "lastTransitionTime": "2019-09-25T09:07:50.224014429Z",
          "name": "0cu4f-worker-000006",
          "version": "2.7.0"
        }
      ],
      "resources": [
        {
          "conditions": [
            {
              "lastTransitionTime": "0001-01-01T00:00:00Z",
              "status": "DeploymentInitialized",
              "type": "Stage"
            }
          ],
          "name": "instancev11"
        }
      ],
      "scaling": {
        "desiredCapacity": 0
      },
      "versions": [
        {
          "date": "0001-01-01T00:00:00Z",
          "lastTransitionTime": "2019-09-25T09:14:26.454983138Z",
          "semver": "2.7.0"
        }
      ]
    }
  }
}
`

// AzureHealthyTC is a JSON string representing a healthy Azure tenant cluster's node list.
var AzureHealthyTC = `
{
    "apiVersion": "v1",
    "items": [
        {
            "apiVersion": "v1",
            "kind": "Node",
            "metadata": {
                "annotations": {
                    "node.alpha.kubernetes.io/ttl": "0",
                    "volumes.kubernetes.io/controller-managed-attach-detach": "true"
                },
                "creationTimestamp": "2019-07-15T08:38:57Z",
                "labels": {
                    "azure-operator.giantswarm.io/version": "2.4.0",
                    "beta.kubernetes.io/arch": "amd64",
                    "beta.kubernetes.io/instance-type": "Standard_D2s_v3",
                    "beta.kubernetes.io/os": "linux",
                    "failure-domain.beta.kubernetes.io/region": "westeurope",
                    "failure-domain.beta.kubernetes.io/zone": "0",
                    "giantswarm.io/provider": "azure",
                    "ip": "10.15.0.5",
                    "kubernetes.io/arch": "amd64",
                    "kubernetes.io/hostname": "6iec4-master-000000",
                    "kubernetes.io/os": "linux",
                    "kubernetes.io/role": "master",
                    "node-role.kubernetes.io/master": "",
                    "node.kubernetes.io/master": "",
                    "role": "master"
                },
                "name": "6iec4-master-000000",
                "resourceVersion": "87461875",
                "selfLink": "/api/v1/nodes/6iec4-master-000000",
                "uid": "fcaf3f98-a6db-11e9-aef7-000d3a43f679"
            },
            "spec": {
                "podCIDR": "10.15.128.0/24",
                "providerID": "azure:///subscriptions/1be3b2e6-497b-45b9-915f-eb35cae23c6a/resourceGroups/6iec4/providers/Microsoft.Compute/virtualMachineScaleSets/6iec4-master/virtualMachines/0",
                "taints": [
                    {
                        "effect": "NoSchedule",
                        "key": "node-role.kubernetes.io/master"
                    }
                ]
            },
            "status": {
                "addresses": [
                    {
                        "address": "10.15.0.5",
                        "type": "InternalIP"
                    },
                    {
                        "address": "6iec4-master-000000",
                        "type": "Hostname"
                    }
                ],
                "allocatable": {
                    "attachable-volumes-azure-disk": "4",
                    "cpu": "2",
                    "ephemeral-storage": "28454196Ki",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "7940436Ki",
                    "pods": "110"
                },
                "capacity": {
                    "attachable-volumes-azure-disk": "4",
                    "cpu": "2",
                    "ephemeral-storage": "28454196Ki",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "8145236Ki",
                    "pods": "110"
                },
                "conditions": [
                    {
                        "lastHeartbeatTime": "2019-07-15T08:41:06Z",
                        "lastTransitionTime": "2019-07-15T08:41:06Z",
                        "message": "RouteController created a route",
                        "reason": "RouteCreated",
                        "status": "False",
                        "type": "NetworkUnavailable"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:47Z",
                        "lastTransitionTime": "2019-07-15T08:38:42Z",
                        "message": "kubelet has sufficient memory available",
                        "reason": "KubeletHasSufficientMemory",
                        "status": "False",
                        "type": "MemoryPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:47Z",
                        "lastTransitionTime": "2019-07-15T08:38:42Z",
                        "message": "kubelet has no disk pressure",
                        "reason": "KubeletHasNoDiskPressure",
                        "status": "False",
                        "type": "DiskPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:47Z",
                        "lastTransitionTime": "2019-07-15T08:38:42Z",
                        "message": "kubelet has sufficient PID available",
                        "reason": "KubeletHasSufficientPID",
                        "status": "False",
                        "type": "PIDPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:47Z",
                        "lastTransitionTime": "2019-07-15T08:39:17Z",
                        "message": "kubelet is posting ready status",
                        "reason": "KubeletReady",
                        "status": "True",
                        "type": "Ready"
                    }
                ],
                "daemonEndpoints": {
                    "kubeletEndpoint": {
                        "Port": 10250
                    }
                },
                "images": [
                    {
                        "names": [
                            "quay.io/giantswarm/hyperkube@sha256:b6cfcecaa8275fd0921bc9c7d839775830d7ab20fd675e2f0fe90c9cec297dcb",
                            "quay.io/giantswarm/hyperkube:v1.14.3"
                        ],
                        "sizeBytes": 603973662
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/node@sha256:4735634d37bb1b15611096617e0d23d0bba6ee187f8ea2a2324ae8901e402716",
                            "quay.io/giantswarm/node:v3.5.1"
                        ],
                        "sizeBytes": 75919845
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cni@sha256:7dd62f4d47fbb2920a0cb54bfc45dc56920c7aaf8150b9096dacd91f66b39657",
                            "quay.io/giantswarm/cni:v3.5.1"
                        ],
                        "sizeBytes": 75434816
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
                            "quay.io/giantswarm/chart-operator@sha256:6819734c898e2d3358c43f15c805f723014c98a64467274d946aa7f19dfeb62b",
                            "quay.io/giantswarm/chart-operator:57d2bdf045a21fd2d279b77acc1cc20ccc870bd7"
                        ],
                        "sizeBytes": 46745395
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
                            "quay.io/giantswarm/net-exporter@sha256:4a3ea165bae3d1c91c36cdb30821ca459b8e3d8c255091bb460af8c21a7a264f",
                            "quay.io/giantswarm/net-exporter:7f178ed2127c0397c9d443bde38beabe435f603a"
                        ],
                        "sizeBytes": 34907651
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:fd333ca6c6b0268bebfca03550a07b959c9fafed2bddf8c559d9ebf050bc855b",
                            "quay.io/giantswarm/net-exporter:1ad313bfd2d835e4b50a606bc510d697976daa93"
                        ],
                        "sizeBytes": 34907614
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:41d00544b36ca1da0a62e326524ebda2e7d94937c81e10872b94306990e55d26",
                            "quay.io/giantswarm/net-exporter:753e1afd06b01c8add545d0673a7dbc5d55880a5"
                        ],
                        "sizeBytes": 32702797
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:32839177290abb6d2f455c09b7d37d9400e29bd30a5a55d8f021195743ce5879",
                            "quay.io/giantswarm/net-exporter:b2fb626131b1ebc941d3fc7f0eac89cd4abfbf41"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:80461ff6f6af959cb1a4de14799da04263787b44e9784182a6a35bb7650462bc",
                            "quay.io/giantswarm/net-exporter:7f2d2252afe1856292c2c2748d02a34b7149b2cb"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:600d067e2a257bc036719ae5c0387f00ca64f5ed65ff5bdad8e65e42b2483a2f",
                            "quay.io/giantswarm/net-exporter:dc80ec9f5e2603bc01626fe8ffa7e5f2c2c6f005"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:8bc26239349b8ab0f93e7b00683b79616a5b3424d595a8553086404c13eda46b",
                            "quay.io/giantswarm/net-exporter:90df7e1dce71e7cf6fcb924e93dcdea2306fc6a2"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:5421106383aa8be7d7482e3c1da6defab40b0f904c18942652520dba84ec95cb",
                            "quay.io/giantswarm/net-exporter:441558237c0023a6a54d11fc4f1d1b2deec3fca9"
                        ],
                        "sizeBytes": 32559983
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:94e1a7e56048e7889a340c953a64ee2f0963cbd52d289515b6c8e291d27a4894",
                            "quay.io/giantswarm/net-exporter:9609ec286a4b55bdaa86fdd8867164e8fe9273d8"
                        ],
                        "sizeBytes": 32559983
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:88d962ca443a338309228922b8fc851da6569056c7144efb389701811098a6fe",
                            "quay.io/giantswarm/net-exporter:3a18405a2445d217c423ba4bca915cb1047d20d0"
                        ],
                        "sizeBytes": 32294939
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:a6f03bb5dfb4b8fc6a40f54299b49602529f9495e7c9e42f22c6d92b948abee7",
                            "quay.io/giantswarm/net-exporter:812f6c6a537582ea0a1011236f76da15bd4f8146"
                        ],
                        "sizeBytes": 32294939
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
                            "quay.io/giantswarm/cert-exporter@sha256:cbcb2ce76df7bdcc61ed6cc447a8e32a00dd8ad1aab7fea781c0861abfb27bcc",
                            "quay.io/giantswarm/cert-exporter:afbc779ea4835d5fa6e7d3c3450846a3a701b835"
                        ],
                        "sizeBytes": 14489266
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:6ca6e09972957c61964d904ca371fe778a2c710f2cdfca6fc98b29b19f3d92b8",
                            "quay.io/giantswarm/cert-exporter:84f8029d53a89549301fa4c1148571ca9aac8175"
                        ],
                        "sizeBytes": 14489266
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:8ff36bdfbc5df23835d5c9f1e0c8f5e34e9192e6bb0d04ddec3313fa26eebad4",
                            "quay.io/giantswarm/cert-exporter:93996b25368a5e72c4e8d08670deef3cb1afaa19"
                        ],
                        "sizeBytes": 14471305
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:434c272bfc2f17d40b8c583b03b3b02f3d7d7a6c4e7f81d931938a815622f5cc",
                            "quay.io/giantswarm/cert-exporter:e9062664cc0b2d89c80e31c69453b0e5d73a35f5"
                        ],
                        "sizeBytes": 14471305
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:8dd95249482ccbf77014b47d48d309c083bf9c69ae71caa0ca8c3bc35ae7d93b",
                            "quay.io/giantswarm/cert-exporter:8ade19af74402d643bb90162b23a175728d4d8e9"
                        ],
                        "sizeBytes": 14471305
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:82ffa782d890234291bbf4ae1d6ab89ade761656f33e17db8882363103d0e1d8",
                            "quay.io/giantswarm/cert-exporter:a84beb354d6db93301b6957d1a1daf35003bf45a"
                        ],
                        "sizeBytes": 14471305
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
                ],
                "nodeInfo": {
                    "architecture": "amd64",
                    "bootID": "d127b0c9-b09a-47cb-868e-386b43b6d0d1",
                    "containerRuntimeVersion": "docker://18.6.3",
                    "kernelVersion": "4.19.50-coreos-r1",
                    "kubeProxyVersion": "v1.14.3",
                    "kubeletVersion": "v1.14.3",
                    "machineID": "5a28fed7d6b6440fa30f780ca655c144",
                    "operatingSystem": "linux",
                    "osImage": "Container Linux by CoreOS 2135.4.0 (Rhyolite)",
                    "systemUUID": "5f1191f6-564e-8043-90e9-ce171335f8c7"
                }
            }
        },
        {
            "apiVersion": "v1",
            "kind": "Node",
            "metadata": {
                "annotations": {
                    "node.alpha.kubernetes.io/ttl": "0",
                    "volumes.kubernetes.io/controller-managed-attach-detach": "true"
                },
                "creationTimestamp": "2019-07-15T08:48:54Z",
                "labels": {
                    "azure-operator.giantswarm.io/version": "2.4.0",
                    "beta.kubernetes.io/arch": "amd64",
                    "beta.kubernetes.io/instance-type": "Standard_A2_v2",
                    "beta.kubernetes.io/os": "linux",
                    "failure-domain.beta.kubernetes.io/region": "westeurope",
                    "failure-domain.beta.kubernetes.io/zone": "3",
                    "giantswarm.io/provider": "azure",
                    "ip": "10.15.1.7",
                    "kubernetes.io/arch": "amd64",
                    "kubernetes.io/hostname": "6iec4-worker-000003",
                    "kubernetes.io/os": "linux",
                    "kubernetes.io/role": "worker",
                    "node-role.kubernetes.io/worker": "",
                    "node.kubernetes.io/worker": "",
                    "role": "worker"
                },
                "name": "6iec4-worker-000003",
                "resourceVersion": "87461874",
                "selfLink": "/api/v1/nodes/6iec4-worker-000003",
                "uid": "605a0b14-a6dd-11e9-aef7-000d3a43f679"
            },
            "spec": {
                "podCIDR": "10.15.129.0/24",
                "providerID": "azure:///subscriptions/1be3b2e6-497b-45b9-915f-eb35cae23c6a/resourceGroups/6iec4/providers/Microsoft.Compute/virtualMachineScaleSets/6iec4-worker/virtualMachines/3"
            },
            "status": {
                "addresses": [
                    {
                        "address": "10.15.1.7",
                        "type": "InternalIP"
                    },
                    {
                        "address": "6iec4-worker-000003",
                        "type": "Hostname"
                    }
                ],
                "allocatable": {
                    "attachable-volumes-azure-disk": "4",
                    "cpu": "2",
                    "ephemeral-storage": "28454196Ki",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "3811668Ki",
                    "pods": "110"
                },
                "capacity": {
                    "attachable-volumes-azure-disk": "4",
                    "cpu": "2",
                    "ephemeral-storage": "28454196Ki",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "4016468Ki",
                    "pods": "110"
                },
                "conditions": [
                    {
                        "lastHeartbeatTime": "2019-07-15T08:49:06Z",
                        "lastTransitionTime": "2019-07-15T08:49:06Z",
                        "message": "RouteController created a route",
                        "reason": "RouteCreated",
                        "status": "False",
                        "type": "NetworkUnavailable"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:47Z",
                        "lastTransitionTime": "2019-07-15T08:48:54Z",
                        "message": "kubelet has sufficient memory available",
                        "reason": "KubeletHasSufficientMemory",
                        "status": "False",
                        "type": "MemoryPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:47Z",
                        "lastTransitionTime": "2019-07-15T08:48:54Z",
                        "message": "kubelet has no disk pressure",
                        "reason": "KubeletHasNoDiskPressure",
                        "status": "False",
                        "type": "DiskPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:47Z",
                        "lastTransitionTime": "2019-07-15T08:48:54Z",
                        "message": "kubelet has sufficient PID available",
                        "reason": "KubeletHasSufficientPID",
                        "status": "False",
                        "type": "PIDPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:47Z",
                        "lastTransitionTime": "2019-07-15T08:49:24Z",
                        "message": "kubelet is posting ready status",
                        "reason": "KubeletReady",
                        "status": "True",
                        "type": "Ready"
                    }
                ],
                "daemonEndpoints": {
                    "kubeletEndpoint": {
                        "Port": 10250
                    }
                },
                "images": [
                    {
                        "names": [
                            "quay.io/giantswarm/nginx-ingress-controller@sha256:9086f624cbddb4265a4d8213a0aaeab6795e274004797cc45932354606853406",
                            "quay.io/giantswarm/nginx-ingress-controller:0.24.1"
                        ],
                        "sizeBytes": 631358200
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/hyperkube@sha256:b6cfcecaa8275fd0921bc9c7d839775830d7ab20fd675e2f0fe90c9cec297dcb",
                            "quay.io/giantswarm/hyperkube:v1.14.3"
                        ],
                        "sizeBytes": 603973662
                    },
                    {
                        "names": [
                            "teemow/debug@sha256:ea9f32fc704ebf1f3fa91cf7581d871107a57122aa38b77872a2130ad2fbd599",
                            "teemow/debug:latest"
                        ],
                        "sizeBytes": 451631394
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:a5f2238a37df9aa9819b131a305b953793cb15886f9ebd89984a268344e21cd7",
                            "quay.io/giantswarm/giantswarmio-webapp:ecac78ce5ce988211be5524a9f53e664bb12a218"
                        ],
                        "sizeBytes": 153717267
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:762bbdf749300afa80699dae65b7381f5bcec31d0d1f2ce356b1636f2479d6db",
                            "quay.io/giantswarm/giantswarmio-webapp:0d2d96a89d89036b4e1d42b33e314ee6570fd195"
                        ],
                        "sizeBytes": 153632248
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:13e08c633812c8a377c13136ab8a1be8fd0af567d6b5dd0b3d3a10ac481369cb",
                            "quay.io/giantswarm/giantswarmio-webapp:649894e9a518865a4f69c670c7a7354bade38576"
                        ],
                        "sizeBytes": 147224012
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:14c3c780fbb5081cd8c45447242b25374bc7adfddfc26b0026d3ced8ca3db3db",
                            "quay.io/giantswarm/giantswarmio-webapp:81090449165423c36387c48f5cd639d2e24d6274"
                        ],
                        "sizeBytes": 147214080
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:053a22c29a4463af5491b13c0798ce500c9959e12af68b56b3a2560ba43dc3c9",
                            "quay.io/giantswarm/giantswarmio-webapp:0539c69b5a0e28d52d0a9b6c627ea65ad73a4f3b"
                        ],
                        "sizeBytes": 147136389
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:5f1805ed757a261be995e0d70971c595484fc0ede743f35d29236912a94923ae",
                            "quay.io/giantswarm/giantswarmio-webapp:960ab74ae0e379b69188e9bbd455ff85656af9ab"
                        ],
                        "sizeBytes": 147135789
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:333d96c3dcc2cf3fd62df953edd4c796914c16b1df7bf208682e56781444b1e5",
                            "quay.io/giantswarm/giantswarmio-webapp:f61c1bb9f8e061d683473b49a2409782e9490e18"
                        ],
                        "sizeBytes": 147120531
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:c91557f5e7e09ec14cba05b142de249a7103f988fa44b66d75e25ef8c6b21134",
                            "quay.io/giantswarm/giantswarmio-webapp:6c0bac466231e3831fa1b952b4c653be67d6f30a"
                        ],
                        "sizeBytes": 147111540
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:997b800f6afb2dc6614a331d7a9af7cc5072dc90cfe591cb274e17201785b3fc",
                            "quay.io/giantswarm/giantswarmio-webapp:4d93451a194c490ed43974c661743266c5116bf8"
                        ],
                        "sizeBytes": 132564126
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs-indexer@sha256:38cddbe7f78d581a699da0d0d6ad64be0e1754f60a1df3a98ea221d5488eff69",
                            "quay.io/giantswarm/docs-indexer:8ba9d3f3f9cbb6d713e478add6c0f2a55d73242c"
                        ],
                        "sizeBytes": 115198415
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/draughtsman@sha256:d3237163bbe45702a8753201276a85dc9c8d1430276aa09049d58d00f77e84d1",
                            "quay.io/giantswarm/draughtsman:f5db76c0526b1a5394d607ab0ec46ca294936ce8"
                        ],
                        "sizeBytes": 104926639
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/node@sha256:4735634d37bb1b15611096617e0d23d0bba6ee187f8ea2a2324ae8901e402716",
                            "quay.io/giantswarm/node:v3.5.1"
                        ],
                        "sizeBytes": 75919845
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cni@sha256:7dd62f4d47fbb2920a0cb54bfc45dc56920c7aaf8150b9096dacd91f66b39657",
                            "quay.io/giantswarm/cni:v3.5.1"
                        ],
                        "sizeBytes": 75434816
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:fb7eaba9a6cb8daedadec8cd6fa6504c5cfde93be7cf947df8c77272b74a2d12",
                            "quay.io/giantswarm/docs:eed835d6dcc366b4ce9b3abb8fa0b42bca3d4b65"
                        ],
                        "sizeBytes": 69751927
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:179d9ef76c74ab394b77b1e20cf170ce0c077d4cf69ef2c61f4231410061840a",
                            "quay.io/giantswarm/docs:45071e90b3a0b4fba8ea0b848695e1cc958bea7a"
                        ],
                        "sizeBytes": 69751371
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:00035fd9ab10a93fabd68aed0d3658b86b463bb17987f756c2f14b6197ad149a",
                            "quay.io/giantswarm/docs:6357e17a68bcc2722b0292fb7f380207d8f99a7e"
                        ],
                        "sizeBytes": 69749049
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:12d6697067c6fbdc8da84b8c2cd5ca1d691234fa6bc111ede9170241a2dd6d65",
                            "quay.io/giantswarm/docs:b47e6ca9762ff7df9b10abee2ff4819fda8678c9"
                        ],
                        "sizeBytes": 69624200
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:8a3e9485f888146d4060cf5dc822219cf8c71d529ec9df6423f3d347a2c48614",
                            "quay.io/giantswarm/docs:80c5cf9b07f8b15ba81954f00ea09ee505aa4e7f"
                        ],
                        "sizeBytes": 69623008
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:8b365eeaab9ccbb542449869b0d5f334cf794561daf9d62450be91a309484b18",
                            "quay.io/giantswarm/docs:8b09158ed582ad5326cd7661452e4cb36f3ac4c0"
                        ],
                        "sizeBytes": 69606280
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:1deac5cb09a7be41c70a6d7ec9c9e6f1a1fcb8510ab31093178293ba59bb0df8",
                            "quay.io/giantswarm/docs:56adfdb23b183848a8f38835bd2d84054a09a16a"
                        ],
                        "sizeBytes": 69605018
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:e7d9b1a83300ecb5a333b963c6cf0f187b9457b8a5ab559f17996e351f5fde58",
                            "quay.io/giantswarm/docs:edaadd4b291e41a92867c3426a88c68a157701d4"
                        ],
                        "sizeBytes": 69121282
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:c350892aaf5b52a36a271c576d76c62a60ff5d1b4851c3d79211cbc6da3801bd",
                            "quay.io/giantswarm/docs:163db17d4de66d941fa3af5135d4f436d54154f4"
                        ],
                        "sizeBytes": 68787076
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
                            "quay.io/giantswarm/net-exporter@sha256:4a3ea165bae3d1c91c36cdb30821ca459b8e3d8c255091bb460af8c21a7a264f",
                            "quay.io/giantswarm/net-exporter:7f178ed2127c0397c9d443bde38beabe435f603a"
                        ],
                        "sizeBytes": 34907651
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:fd333ca6c6b0268bebfca03550a07b959c9fafed2bddf8c559d9ebf050bc855b",
                            "quay.io/giantswarm/net-exporter:1ad313bfd2d835e4b50a606bc510d697976daa93"
                        ],
                        "sizeBytes": 34907614
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:41d00544b36ca1da0a62e326524ebda2e7d94937c81e10872b94306990e55d26",
                            "quay.io/giantswarm/net-exporter:753e1afd06b01c8add545d0673a7dbc5d55880a5"
                        ],
                        "sizeBytes": 32702797
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:80461ff6f6af959cb1a4de14799da04263787b44e9784182a6a35bb7650462bc",
                            "quay.io/giantswarm/net-exporter:7f2d2252afe1856292c2c2748d02a34b7149b2cb"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:600d067e2a257bc036719ae5c0387f00ca64f5ed65ff5bdad8e65e42b2483a2f",
                            "quay.io/giantswarm/net-exporter:dc80ec9f5e2603bc01626fe8ffa7e5f2c2c6f005"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:8bc26239349b8ab0f93e7b00683b79616a5b3424d595a8553086404c13eda46b",
                            "quay.io/giantswarm/net-exporter:90df7e1dce71e7cf6fcb924e93dcdea2306fc6a2"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:32839177290abb6d2f455c09b7d37d9400e29bd30a5a55d8f021195743ce5879",
                            "quay.io/giantswarm/net-exporter:b2fb626131b1ebc941d3fc7f0eac89cd4abfbf41"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:94e1a7e56048e7889a340c953a64ee2f0963cbd52d289515b6c8e291d27a4894",
                            "quay.io/giantswarm/net-exporter:9609ec286a4b55bdaa86fdd8867164e8fe9273d8"
                        ],
                        "sizeBytes": 32559983
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:5421106383aa8be7d7482e3c1da6defab40b0f904c18942652520dba84ec95cb",
                            "quay.io/giantswarm/net-exporter:441558237c0023a6a54d11fc4f1d1b2deec3fca9"
                        ],
                        "sizeBytes": 32559983
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:88d962ca443a338309228922b8fc851da6569056c7144efb389701811098a6fe",
                            "quay.io/giantswarm/net-exporter:3a18405a2445d217c423ba4bca915cb1047d20d0"
                        ],
                        "sizeBytes": 32294939
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:a6f03bb5dfb4b8fc6a40f54299b49602529f9495e7c9e42f22c6d92b948abee7",
                            "quay.io/giantswarm/net-exporter:812f6c6a537582ea0a1011236f76da15bd4f8146"
                        ],
                        "sizeBytes": 32294939
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
                            "quay.io/giantswarm/api-spec@sha256:50410e077f1c0b07276dad07432f0e48ecd06ddfb940b9ed5faca2d072268374",
                            "quay.io/giantswarm/api-spec:c481aed170dad9189b3b870a1bd0fb723d351c61"
                        ],
                        "sizeBytes": 22729759
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/api-spec@sha256:fd6fb3608f55768c10d8da831aef24347d905c798b5d1ac299f64b84d585e55f",
                            "quay.io/giantswarm/api-spec:9535cb7cddcdda07b24f3d4a2ff88113637ab94b"
                        ],
                        "sizeBytes": 22726204
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/redis@sha256:604e75bf72bb6c80023f500e8d36135d02553e94cdf77eb4fffaa7a549ca10d7",
                            "quay.io/giantswarm/redis:3.2.11-alpine"
                        ],
                        "sizeBytes": 20658687
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-nginx@sha256:7723c8d619630ae8584ab5a559a33bc3897e4b0fb1f9522900a1ecd5fc431b82",
                            "quay.io/giantswarm/giantswarmio-nginx:14481bec984a6da07d4a96ef2d08ec9e21f9104b"
                        ],
                        "sizeBytes": 18481231
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs-proxy@sha256:a1a1d9a8a70b4d10b6956a34c297fb2ec81d0cd0eeeecc26b0213fe5c953a0b5",
                            "quay.io/giantswarm/docs-proxy:f68e0e1b306452505ebcd901863756f6bacee3dd"
                        ],
                        "sizeBytes": 17408306
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
                            "quay.io/giantswarm/cert-exporter@sha256:cbcb2ce76df7bdcc61ed6cc447a8e32a00dd8ad1aab7fea781c0861abfb27bcc",
                            "quay.io/giantswarm/cert-exporter:afbc779ea4835d5fa6e7d3c3450846a3a701b835"
                        ],
                        "sizeBytes": 14489266
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:6ca6e09972957c61964d904ca371fe778a2c710f2cdfca6fc98b29b19f3d92b8",
                            "quay.io/giantswarm/cert-exporter:84f8029d53a89549301fa4c1148571ca9aac8175"
                        ],
                        "sizeBytes": 14489266
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:434c272bfc2f17d40b8c583b03b3b02f3d7d7a6c4e7f81d931938a815622f5cc",
                            "quay.io/giantswarm/cert-exporter:e9062664cc0b2d89c80e31c69453b0e5d73a35f5"
                        ],
                        "sizeBytes": 14471305
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:8dd95249482ccbf77014b47d48d309c083bf9c69ae71caa0ca8c3bc35ae7d93b",
                            "quay.io/giantswarm/cert-exporter:8ade19af74402d643bb90162b23a175728d4d8e9"
                        ],
                        "sizeBytes": 14471305
                    }
                ],
                "nodeInfo": {
                    "architecture": "amd64",
                    "bootID": "d2c2f7a9-5284-4ab6-b4c9-e06aaf98ff69",
                    "containerRuntimeVersion": "docker://18.6.3",
                    "kernelVersion": "4.19.50-coreos-r1",
                    "kubeProxyVersion": "v1.14.3",
                    "kubeletVersion": "v1.14.3",
                    "machineID": "dde62c02a4444f919bafc765b43d4dcd",
                    "operatingSystem": "linux",
                    "osImage": "Container Linux by CoreOS 2135.4.0 (Rhyolite)",
                    "systemUUID": "cf52146f-5351-d34c-a657-5e3e3aee3067"
                }
            }
        },
        {
            "apiVersion": "v1",
            "kind": "Node",
            "metadata": {
                "annotations": {
                    "node.alpha.kubernetes.io/ttl": "0",
                    "volumes.kubernetes.io/controller-managed-attach-detach": "true"
                },
                "creationTimestamp": "2019-07-15T09:03:28Z",
                "labels": {
                    "azure-operator.giantswarm.io/version": "2.4.0",
                    "beta.kubernetes.io/arch": "amd64",
                    "beta.kubernetes.io/instance-type": "Standard_A2_v2",
                    "beta.kubernetes.io/os": "linux",
                    "failure-domain.beta.kubernetes.io/region": "westeurope",
                    "failure-domain.beta.kubernetes.io/zone": "0",
                    "giantswarm.io/provider": "azure",
                    "ip": "10.15.1.4",
                    "kubernetes.io/arch": "amd64",
                    "kubernetes.io/hostname": "6iec4-worker-000004",
                    "kubernetes.io/os": "linux",
                    "kubernetes.io/role": "worker",
                    "node-role.kubernetes.io/worker": "",
                    "node.kubernetes.io/worker": "",
                    "role": "worker"
                },
                "name": "6iec4-worker-000004",
                "resourceVersion": "87461866",
                "selfLink": "/api/v1/nodes/6iec4-worker-000004",
                "uid": "6978bf91-a6df-11e9-aef7-000d3a43f679"
            },
            "spec": {
                "podCIDR": "10.15.130.0/24",
                "providerID": "azure:///subscriptions/1be3b2e6-497b-45b9-915f-eb35cae23c6a/resourceGroups/6iec4/providers/Microsoft.Compute/virtualMachineScaleSets/6iec4-worker/virtualMachines/4"
            },
            "status": {
                "addresses": [
                    {
                        "address": "10.15.1.4",
                        "type": "InternalIP"
                    },
                    {
                        "address": "6iec4-worker-000004",
                        "type": "Hostname"
                    }
                ],
                "allocatable": {
                    "attachable-volumes-azure-disk": "4",
                    "cpu": "2",
                    "ephemeral-storage": "28454196Ki",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "3811668Ki",
                    "pods": "110"
                },
                "capacity": {
                    "attachable-volumes-azure-disk": "4",
                    "cpu": "2",
                    "ephemeral-storage": "28454196Ki",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "4016468Ki",
                    "pods": "110"
                },
                "conditions": [
                    {
                        "lastHeartbeatTime": "2019-07-15T09:05:07Z",
                        "lastTransitionTime": "2019-07-15T09:05:07Z",
                        "message": "RouteController created a route",
                        "reason": "RouteCreated",
                        "status": "False",
                        "type": "NetworkUnavailable"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:44Z",
                        "lastTransitionTime": "2019-07-15T09:03:28Z",
                        "message": "kubelet has sufficient memory available",
                        "reason": "KubeletHasSufficientMemory",
                        "status": "False",
                        "type": "MemoryPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:44Z",
                        "lastTransitionTime": "2019-07-15T09:03:28Z",
                        "message": "kubelet has no disk pressure",
                        "reason": "KubeletHasNoDiskPressure",
                        "status": "False",
                        "type": "DiskPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:44Z",
                        "lastTransitionTime": "2019-07-15T09:03:28Z",
                        "message": "kubelet has sufficient PID available",
                        "reason": "KubeletHasSufficientPID",
                        "status": "False",
                        "type": "PIDPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:44Z",
                        "lastTransitionTime": "2019-07-15T09:03:49Z",
                        "message": "kubelet is posting ready status",
                        "reason": "KubeletReady",
                        "status": "True",
                        "type": "Ready"
                    }
                ],
                "daemonEndpoints": {
                    "kubeletEndpoint": {
                        "Port": 10250
                    }
                },
                "images": [
                    {
                        "names": [
                            "quay.io/giantswarm/nginx-ingress-controller@sha256:9086f624cbddb4265a4d8213a0aaeab6795e274004797cc45932354606853406",
                            "quay.io/giantswarm/nginx-ingress-controller:0.24.1"
                        ],
                        "sizeBytes": 631358200
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/hyperkube@sha256:b6cfcecaa8275fd0921bc9c7d839775830d7ab20fd675e2f0fe90c9cec297dcb",
                            "quay.io/giantswarm/hyperkube:v1.14.3"
                        ],
                        "sizeBytes": 603973662
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:69a83fa3a5a40ec68c41ec1b43d46e007eed1a07017b79ba58247111ca3857ff",
                            "quay.io/giantswarm/giantswarmio-webapp:03c853e56743aa585a7ec19e4cdac833e89d49f2"
                        ],
                        "sizeBytes": 147132682
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:dfcb79b116e07bdf5b18ba6c83f16b5123ce43e364da98853d0ceb6d91899ff6",
                            "quay.io/giantswarm/giantswarmio-webapp:ac0f3a19cfc4ab2aa66be0506ac9268c97102775"
                        ],
                        "sizeBytes": 147132137
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:ec6ccc8abdaf1ec31983247f34ca9fa827f8fc58d1d65c458dc556d1e0635dac",
                            "quay.io/giantswarm/giantswarmio-webapp:8c7a9b1e31d687af7fa1e6b5d75334a3a6200218"
                        ],
                        "sizeBytes": 147111609
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:997b800f6afb2dc6614a331d7a9af7cc5072dc90cfe591cb274e17201785b3fc",
                            "quay.io/giantswarm/giantswarmio-webapp:4d93451a194c490ed43974c661743266c5116bf8"
                        ],
                        "sizeBytes": 132564126
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/prometheus@sha256:790832143cb8969b6978d943da1bd9047e1cd69344405826955af82482f1ec37",
                            "quay.io/giantswarm/prometheus:v2.9.1"
                        ],
                        "sizeBytes": 121187253
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
                            "quay.io/giantswarm/node@sha256:4735634d37bb1b15611096617e0d23d0bba6ee187f8ea2a2324ae8901e402716",
                            "quay.io/giantswarm/node:v3.5.1"
                        ],
                        "sizeBytes": 75919845
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cni@sha256:7dd62f4d47fbb2920a0cb54bfc45dc56920c7aaf8150b9096dacd91f66b39657",
                            "quay.io/giantswarm/cni:v3.5.1"
                        ],
                        "sizeBytes": 75434816
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:c944c511e8392298f9a7a529daacd80bf7ebf731c93322f2482090f92d100a2c",
                            "quay.io/giantswarm/docs:766d83031cb60787290adf065ede7126e08bd964"
                        ],
                        "sizeBytes": 69750504
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:8a3e9485f888146d4060cf5dc822219cf8c71d529ec9df6423f3d347a2c48614",
                            "quay.io/giantswarm/docs:80c5cf9b07f8b15ba81954f00ea09ee505aa4e7f"
                        ],
                        "sizeBytes": 69623008
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:b21deb4691862992e7de5752619cccadfc5db5ae4016e7565a1d2192dca03458",
                            "quay.io/giantswarm/docs:5ec7929c0c3c401ea0c5a9fab850d26022bcc45a"
                        ],
                        "sizeBytes": 69605018
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:e7d9b1a83300ecb5a333b963c6cf0f187b9457b8a5ab559f17996e351f5fde58",
                            "quay.io/giantswarm/docs:edaadd4b291e41a92867c3426a88c68a157701d4"
                        ],
                        "sizeBytes": 69121282
                    },
                    {
                        "names": [
                            "quay.io/jetstack/cert-manager-controller@sha256:619be127a7fc9d5b1c3d5b066485ba3fedf82fc8994f6f5da02b78b84e654a66",
                            "quay.io/jetstack/cert-manager-controller:v0.4.1"
                        ],
                        "sizeBytes": 51851986
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
                            "quay.io/giantswarm/net-exporter@sha256:4a3ea165bae3d1c91c36cdb30821ca459b8e3d8c255091bb460af8c21a7a264f",
                            "quay.io/giantswarm/net-exporter:7f178ed2127c0397c9d443bde38beabe435f603a"
                        ],
                        "sizeBytes": 34907651
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:fd333ca6c6b0268bebfca03550a07b959c9fafed2bddf8c559d9ebf050bc855b",
                            "quay.io/giantswarm/net-exporter:1ad313bfd2d835e4b50a606bc510d697976daa93"
                        ],
                        "sizeBytes": 34907614
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:41d00544b36ca1da0a62e326524ebda2e7d94937c81e10872b94306990e55d26",
                            "quay.io/giantswarm/net-exporter:753e1afd06b01c8add545d0673a7dbc5d55880a5"
                        ],
                        "sizeBytes": 32702797
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:32839177290abb6d2f455c09b7d37d9400e29bd30a5a55d8f021195743ce5879",
                            "quay.io/giantswarm/net-exporter:b2fb626131b1ebc941d3fc7f0eac89cd4abfbf41"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:8bc26239349b8ab0f93e7b00683b79616a5b3424d595a8553086404c13eda46b",
                            "quay.io/giantswarm/net-exporter:90df7e1dce71e7cf6fcb924e93dcdea2306fc6a2"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:80461ff6f6af959cb1a4de14799da04263787b44e9784182a6a35bb7650462bc",
                            "quay.io/giantswarm/net-exporter:7f2d2252afe1856292c2c2748d02a34b7149b2cb"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:600d067e2a257bc036719ae5c0387f00ca64f5ed65ff5bdad8e65e42b2483a2f",
                            "quay.io/giantswarm/net-exporter:dc80ec9f5e2603bc01626fe8ffa7e5f2c2c6f005"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:94e1a7e56048e7889a340c953a64ee2f0963cbd52d289515b6c8e291d27a4894",
                            "quay.io/giantswarm/net-exporter:9609ec286a4b55bdaa86fdd8867164e8fe9273d8"
                        ],
                        "sizeBytes": 32559983
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:5421106383aa8be7d7482e3c1da6defab40b0f904c18942652520dba84ec95cb",
                            "quay.io/giantswarm/net-exporter:441558237c0023a6a54d11fc4f1d1b2deec3fca9"
                        ],
                        "sizeBytes": 32559983
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:a6f03bb5dfb4b8fc6a40f54299b49602529f9495e7c9e42f22c6d92b948abee7",
                            "quay.io/giantswarm/net-exporter:812f6c6a537582ea0a1011236f76da15bd4f8146"
                        ],
                        "sizeBytes": 32294939
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:88d962ca443a338309228922b8fc851da6569056c7144efb389701811098a6fe",
                            "quay.io/giantswarm/net-exporter:3a18405a2445d217c423ba4bca915cb1047d20d0"
                        ],
                        "sizeBytes": 32294939
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
                            "quay.io/giantswarm/giantswarmio-nginx@sha256:023261b50767a7b9b8b8fd835246570d368e2c0d2a119eb0e453c2671a47b322",
                            "quay.io/giantswarm/giantswarmio-nginx:0ee7977c8e10b41b513c5edbd391a44ba11bb119"
                        ],
                        "sizeBytes": 21970970
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/redis@sha256:604e75bf72bb6c80023f500e8d36135d02553e94cdf77eb4fffaa7a549ca10d7",
                            "quay.io/giantswarm/redis:3.2.11-alpine"
                        ],
                        "sizeBytes": 20658687
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-nginx@sha256:7723c8d619630ae8584ab5a559a33bc3897e4b0fb1f9522900a1ecd5fc431b82",
                            "quay.io/giantswarm/giantswarmio-nginx:14481bec984a6da07d4a96ef2d08ec9e21f9104b"
                        ],
                        "sizeBytes": 18481231
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
                            "quay.io/giantswarm/cert-exporter@sha256:cbcb2ce76df7bdcc61ed6cc447a8e32a00dd8ad1aab7fea781c0861abfb27bcc",
                            "quay.io/giantswarm/cert-exporter:afbc779ea4835d5fa6e7d3c3450846a3a701b835"
                        ],
                        "sizeBytes": 14489266
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:6ca6e09972957c61964d904ca371fe778a2c710f2cdfca6fc98b29b19f3d92b8",
                            "quay.io/giantswarm/cert-exporter:84f8029d53a89549301fa4c1148571ca9aac8175"
                        ],
                        "sizeBytes": 14489266
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:434c272bfc2f17d40b8c583b03b3b02f3d7d7a6c4e7f81d931938a815622f5cc",
                            "quay.io/giantswarm/cert-exporter:e9062664cc0b2d89c80e31c69453b0e5d73a35f5"
                        ],
                        "sizeBytes": 14471305
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:82ffa782d890234291bbf4ae1d6ab89ade761656f33e17db8882363103d0e1d8",
                            "quay.io/giantswarm/cert-exporter:a84beb354d6db93301b6957d1a1daf35003bf45a"
                        ],
                        "sizeBytes": 14471305
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:8dd95249482ccbf77014b47d48d309c083bf9c69ae71caa0ca8c3bc35ae7d93b",
                            "quay.io/giantswarm/cert-exporter:8ade19af74402d643bb90162b23a175728d4d8e9"
                        ],
                        "sizeBytes": 14471305
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:8ff36bdfbc5df23835d5c9f1e0c8f5e34e9192e6bb0d04ddec3313fa26eebad4",
                            "quay.io/giantswarm/cert-exporter:93996b25368a5e72c4e8d08670deef3cb1afaa19"
                        ],
                        "sizeBytes": 14471305
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
                ],
                "nodeInfo": {
                    "architecture": "amd64",
                    "bootID": "890be693-cc30-45a7-84b9-596d37ab67e2",
                    "containerRuntimeVersion": "docker://18.6.3",
                    "kernelVersion": "4.19.50-coreos-r1",
                    "kubeProxyVersion": "v1.14.3",
                    "kubeletVersion": "v1.14.3",
                    "machineID": "44c8140d28b143c6922607325da224cb",
                    "operatingSystem": "linux",
                    "osImage": "Container Linux by CoreOS 2135.4.0 (Rhyolite)",
                    "systemUUID": "d3242c82-08d9-5042-8305-123f0a1c9d21"
                },
                "volumesAttached": [
                    {
                        "devicePath": "-1",
                        "name": "kubernetes.io/azure-disk//subscriptions/1be3b2e6-497b-45b9-915f-eb35cae23c6a/resourceGroups/6iec4/providers/Microsoft.Compute/disks/kubernetes-dynamic-pvc-d16e45f2-6810-11e9-870a-000d3a43f679"
                    }
                ],
                "volumesInUse": [
                    "kubernetes.io/azure-disk//subscriptions/1be3b2e6-497b-45b9-915f-eb35cae23c6a/resourceGroups/6iec4/providers/Microsoft.Compute/disks/kubernetes-dynamic-pvc-d16e45f2-6810-11e9-870a-000d3a43f679"
                ]
            }
        },
        {
            "apiVersion": "v1",
            "kind": "Node",
            "metadata": {
                "annotations": {
                    "node.alpha.kubernetes.io/ttl": "0",
                    "volumes.kubernetes.io/controller-managed-attach-detach": "true"
                },
                "creationTimestamp": "2019-07-15T09:47:41Z",
                "labels": {
                    "azure-operator.giantswarm.io/version": "2.4.0",
                    "beta.kubernetes.io/arch": "amd64",
                    "beta.kubernetes.io/instance-type": "Standard_A2_v2",
                    "beta.kubernetes.io/os": "linux",
                    "failure-domain.beta.kubernetes.io/region": "westeurope",
                    "failure-domain.beta.kubernetes.io/zone": "4",
                    "giantswarm.io/provider": "azure",
                    "ip": "10.15.1.10",
                    "kubernetes.io/arch": "amd64",
                    "kubernetes.io/hostname": "6iec4-worker-000009",
                    "kubernetes.io/os": "linux",
                    "kubernetes.io/role": "worker",
                    "node-role.kubernetes.io/worker": "",
                    "node.kubernetes.io/worker": "",
                    "role": "worker"
                },
                "name": "6iec4-worker-000009",
                "resourceVersion": "87461890",
                "selfLink": "/api/v1/nodes/6iec4-worker-000009",
                "uid": "96661340-a6e5-11e9-aef7-000d3a43f679"
            },
            "spec": {
                "podCIDR": "10.15.131.0/24",
                "providerID": "azure:///subscriptions/1be3b2e6-497b-45b9-915f-eb35cae23c6a/resourceGroups/6iec4/providers/Microsoft.Compute/virtualMachineScaleSets/6iec4-worker/virtualMachines/9"
            },
            "status": {
                "addresses": [
                    {
                        "address": "10.15.1.10",
                        "type": "InternalIP"
                    },
                    {
                        "address": "6iec4-worker-000009",
                        "type": "Hostname"
                    }
                ],
                "allocatable": {
                    "attachable-volumes-azure-disk": "4",
                    "cpu": "2",
                    "ephemeral-storage": "28454196Ki",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "3811660Ki",
                    "pods": "110"
                },
                "capacity": {
                    "attachable-volumes-azure-disk": "4",
                    "cpu": "2",
                    "ephemeral-storage": "28454196Ki",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "4016460Ki",
                    "pods": "110"
                },
                "conditions": [
                    {
                        "lastHeartbeatTime": "2019-07-15T09:49:07Z",
                        "lastTransitionTime": "2019-07-15T09:49:07Z",
                        "message": "RouteController created a route",
                        "reason": "RouteCreated",
                        "status": "False",
                        "type": "NetworkUnavailable"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:52Z",
                        "lastTransitionTime": "2019-07-15T09:47:41Z",
                        "message": "kubelet has sufficient memory available",
                        "reason": "KubeletHasSufficientMemory",
                        "status": "False",
                        "type": "MemoryPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:52Z",
                        "lastTransitionTime": "2019-07-15T09:47:41Z",
                        "message": "kubelet has no disk pressure",
                        "reason": "KubeletHasNoDiskPressure",
                        "status": "False",
                        "type": "DiskPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:52Z",
                        "lastTransitionTime": "2019-07-15T09:47:41Z",
                        "message": "kubelet has sufficient PID available",
                        "reason": "KubeletHasSufficientPID",
                        "status": "False",
                        "type": "PIDPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:52Z",
                        "lastTransitionTime": "2019-07-15T09:48:01Z",
                        "message": "kubelet is posting ready status",
                        "reason": "KubeletReady",
                        "status": "True",
                        "type": "Ready"
                    }
                ],
                "daemonEndpoints": {
                    "kubeletEndpoint": {
                        "Port": 10250
                    }
                },
                "images": [
                    {
                        "names": [
                            "quay.io/giantswarm/nginx-ingress-controller@sha256:9086f624cbddb4265a4d8213a0aaeab6795e274004797cc45932354606853406",
                            "quay.io/giantswarm/nginx-ingress-controller:0.24.1"
                        ],
                        "sizeBytes": 631358200
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/hyperkube@sha256:b6cfcecaa8275fd0921bc9c7d839775830d7ab20fd675e2f0fe90c9cec297dcb",
                            "quay.io/giantswarm/hyperkube:v1.14.3"
                        ],
                        "sizeBytes": 603973662
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/hound@sha256:778b9d70bf6a388a6261959fac93fc804bbb056a687f3ea1786a8568a98e57bb",
                            "quay.io/giantswarm/hound:0.1.3"
                        ],
                        "sizeBytes": 192715802
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:cf17afb6c6b97d87570d17e90464186d8fbfc722192d40db7eb03993a7cf5d46",
                            "quay.io/giantswarm/giantswarmio-webapp:540e23d2208d41e6c5bc8b0861b3da60729e20ec"
                        ],
                        "sizeBytes": 153717050
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:762bbdf749300afa80699dae65b7381f5bcec31d0d1f2ce356b1636f2479d6db",
                            "quay.io/giantswarm/giantswarmio-webapp:0d2d96a89d89036b4e1d42b33e314ee6570fd195"
                        ],
                        "sizeBytes": 153632248
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:097ce3ee8be6f7557e546a8bdbbf993b66c9f69d84869afe3872b00ef5459c81",
                            "quay.io/giantswarm/giantswarmio-webapp:4a112738beef5ba5fe95458c12b96f291078fb28"
                        ],
                        "sizeBytes": 153632225
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:df0d937111e3b4f87fe468f3c420288da04928460fb14a879270658a5d855127",
                            "quay.io/giantswarm/giantswarmio-webapp:f7bbf78c49b83e728e644a09e857ebebb2366d93"
                        ],
                        "sizeBytes": 147145121
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:9c479b566994ebfd7b899a6800522b183869ce3d2110160c54c00b51950e757b",
                            "quay.io/giantswarm/giantswarmio-webapp:a81b18d438212d8497f54ae146add487a52d1c66"
                        ],
                        "sizeBytes": 147128122
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:3a97718451fea83b498b7b6dec4e7f3c75fc5e57caf881de8fe6267e8172b17b",
                            "quay.io/giantswarm/giantswarmio-webapp:c2d0e1616565bfec161378aae0d143c87af78046"
                        ],
                        "sizeBytes": 147102708
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs-indexer@sha256:38cddbe7f78d581a699da0d0d6ad64be0e1754f60a1df3a98ea221d5488eff69",
                            "quay.io/giantswarm/docs-indexer:8ba9d3f3f9cbb6d713e478add6c0f2a55d73242c"
                        ],
                        "sizeBytes": 115198415
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/draughtsman@sha256:8c1cd20b6a2c0349f935f1459be1b6369f191a99689f4816c1468da1a678511e",
                            "quay.io/giantswarm/draughtsman:4f401dddabf6b43ebd7eba8ecb3d5db875748b86"
                        ],
                        "sizeBytes": 104928330
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
                            "quay.io/giantswarm/node@sha256:4735634d37bb1b15611096617e0d23d0bba6ee187f8ea2a2324ae8901e402716",
                            "quay.io/giantswarm/node:v3.5.1"
                        ],
                        "sizeBytes": 75919845
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cni@sha256:7dd62f4d47fbb2920a0cb54bfc45dc56920c7aaf8150b9096dacd91f66b39657",
                            "quay.io/giantswarm/cni:v3.5.1"
                        ],
                        "sizeBytes": 75434816
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:4a20140aaf748eea1b8035ba4c47e7ac3cd4ccb4f70d38d2eb0d37cf764fa6c5",
                            "quay.io/giantswarm/docs:4224533c62ffd39d709f9901fd11309289bbdc95"
                        ],
                        "sizeBytes": 69749459
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:c7ecc7c809a0cbfa2f7533dcdc434b05e9f17a2dfec0059f4acf312528ff1982",
                            "quay.io/giantswarm/docs:6ff58ccfb6541e89fa4e541792d1eef2d16cd574"
                        ],
                        "sizeBytes": 69749376
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:ebb5f2a40a0e9f7f4008280ace9381cc00c5c2ed30a63368a5628c7fe5a546b2",
                            "quay.io/giantswarm/docs:a319fdd9880a9c9620ddc36e0a602fad48fbaffc"
                        ],
                        "sizeBytes": 69220401
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
                            "quay.io/giantswarm/net-exporter@sha256:4a3ea165bae3d1c91c36cdb30821ca459b8e3d8c255091bb460af8c21a7a264f",
                            "quay.io/giantswarm/net-exporter:7f178ed2127c0397c9d443bde38beabe435f603a"
                        ],
                        "sizeBytes": 34907651
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:fd333ca6c6b0268bebfca03550a07b959c9fafed2bddf8c559d9ebf050bc855b",
                            "quay.io/giantswarm/net-exporter:1ad313bfd2d835e4b50a606bc510d697976daa93"
                        ],
                        "sizeBytes": 34907614
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:41d00544b36ca1da0a62e326524ebda2e7d94937c81e10872b94306990e55d26",
                            "quay.io/giantswarm/net-exporter:753e1afd06b01c8add545d0673a7dbc5d55880a5"
                        ],
                        "sizeBytes": 32702797
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:8bc26239349b8ab0f93e7b00683b79616a5b3424d595a8553086404c13eda46b",
                            "quay.io/giantswarm/net-exporter:90df7e1dce71e7cf6fcb924e93dcdea2306fc6a2"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:80461ff6f6af959cb1a4de14799da04263787b44e9784182a6a35bb7650462bc",
                            "quay.io/giantswarm/net-exporter:7f2d2252afe1856292c2c2748d02a34b7149b2cb"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:600d067e2a257bc036719ae5c0387f00ca64f5ed65ff5bdad8e65e42b2483a2f",
                            "quay.io/giantswarm/net-exporter:dc80ec9f5e2603bc01626fe8ffa7e5f2c2c6f005"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:32839177290abb6d2f455c09b7d37d9400e29bd30a5a55d8f021195743ce5879",
                            "quay.io/giantswarm/net-exporter:b2fb626131b1ebc941d3fc7f0eac89cd4abfbf41"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:5421106383aa8be7d7482e3c1da6defab40b0f904c18942652520dba84ec95cb",
                            "quay.io/giantswarm/net-exporter:441558237c0023a6a54d11fc4f1d1b2deec3fca9"
                        ],
                        "sizeBytes": 32559983
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:94e1a7e56048e7889a340c953a64ee2f0963cbd52d289515b6c8e291d27a4894",
                            "quay.io/giantswarm/net-exporter:9609ec286a4b55bdaa86fdd8867164e8fe9273d8"
                        ],
                        "sizeBytes": 32559983
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:88d962ca443a338309228922b8fc851da6569056c7144efb389701811098a6fe",
                            "quay.io/giantswarm/net-exporter:3a18405a2445d217c423ba4bca915cb1047d20d0"
                        ],
                        "sizeBytes": 32294939
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:a6f03bb5dfb4b8fc6a40f54299b49602529f9495e7c9e42f22c6d92b948abee7",
                            "quay.io/giantswarm/net-exporter:812f6c6a537582ea0a1011236f76da15bd4f8146"
                        ],
                        "sizeBytes": 32294939
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
                            "quay.io/giantswarm/api-spec@sha256:fea4a288f39b4c6a32894aba36ae42d52c183b1f4f99fef7303574498b1897ed",
                            "quay.io/giantswarm/api-spec:d485fbcc933b8ce5c3bbff3c5bef185e1b8dca51"
                        ],
                        "sizeBytes": 23539782
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/api-spec@sha256:ea4fc9a60f411a427a879e414cab0d15de69c0de278b221b59a822b2a7d0cc37",
                            "quay.io/giantswarm/api-spec:a3d9968a6d82afaa061856e11d55b5f894bebad5"
                        ],
                        "sizeBytes": 23536484
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/api-spec@sha256:312c9da3cb762618d3760742b036ec1a82ba91e3e16938bc64c18f4ec2fc0369",
                            "quay.io/giantswarm/api-spec:67949a50b8bd55f956c56494bbb3443bc70bab55"
                        ],
                        "sizeBytes": 23536020
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/api-spec@sha256:7032f467f4f762958fb9626ee22713e2d216994d405e6e4f63d9b53500f179e8",
                            "quay.io/giantswarm/api-spec:85ad5086cee3b02960715a555dc0d56bfbb45f63"
                        ],
                        "sizeBytes": 23527771
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
                            "quay.io/giantswarm/api-spec@sha256:50410e077f1c0b07276dad07432f0e48ecd06ddfb940b9ed5faca2d072268374",
                            "quay.io/giantswarm/api-spec:c481aed170dad9189b3b870a1bd0fb723d351c61"
                        ],
                        "sizeBytes": 22729759
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/api-spec@sha256:fd6fb3608f55768c10d8da831aef24347d905c798b5d1ac299f64b84d585e55f",
                            "quay.io/giantswarm/api-spec:9535cb7cddcdda07b24f3d4a2ff88113637ab94b"
                        ],
                        "sizeBytes": 22726204
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-nginx@sha256:47142466cffa86dfa99cbd6e3510e79acb0d5e23d48fcdd7ffc3a430bbac960d",
                            "quay.io/giantswarm/giantswarmio-nginx:1a5c26c638caa7e1c5f6d3c1acf9f8b2620b4b75"
                        ],
                        "sizeBytes": 21970970
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/redis@sha256:604e75bf72bb6c80023f500e8d36135d02553e94cdf77eb4fffaa7a549ca10d7",
                            "quay.io/giantswarm/redis:3.2.11-alpine"
                        ],
                        "sizeBytes": 20658687
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/auto-oncall@sha256:2f083bcefe4ed729563ba92ac8fbf1d0487107b541efebd1905a2a1936483917",
                            "quay.io/giantswarm/auto-oncall:dbef151a94d3ce79bdedab06fcf32488056d33f5"
                        ],
                        "sizeBytes": 17593873
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/auto-oncall@sha256:f480a0632d535eb3d6581cd153e6a38955a1cdeb62b3f269c3bf52ce744a3a87",
                            "quay.io/giantswarm/auto-oncall:fa1f92eeba8d1ed258d67188cf481af2820039df"
                        ],
                        "sizeBytes": 17583891
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/oncall-scheduler@sha256:c3c4612d6646f624830d9dfd1cdc0452fdad90178f10083d08a8a16f8d372639",
                            "quay.io/giantswarm/oncall-scheduler:latest"
                        ],
                        "sizeBytes": 15559099
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
                            "quay.io/giantswarm/cert-exporter@sha256:cbcb2ce76df7bdcc61ed6cc447a8e32a00dd8ad1aab7fea781c0861abfb27bcc",
                            "quay.io/giantswarm/cert-exporter:afbc779ea4835d5fa6e7d3c3450846a3a701b835"
                        ],
                        "sizeBytes": 14489266
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:6ca6e09972957c61964d904ca371fe778a2c710f2cdfca6fc98b29b19f3d92b8",
                            "quay.io/giantswarm/cert-exporter:84f8029d53a89549301fa4c1148571ca9aac8175"
                        ],
                        "sizeBytes": 14489266
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:8dd95249482ccbf77014b47d48d309c083bf9c69ae71caa0ca8c3bc35ae7d93b",
                            "quay.io/giantswarm/cert-exporter:8ade19af74402d643bb90162b23a175728d4d8e9"
                        ],
                        "sizeBytes": 14471305
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:434c272bfc2f17d40b8c583b03b3b02f3d7d7a6c4e7f81d931938a815622f5cc",
                            "quay.io/giantswarm/cert-exporter:e9062664cc0b2d89c80e31c69453b0e5d73a35f5"
                        ],
                        "sizeBytes": 14471305
                    }
                ],
                "nodeInfo": {
                    "architecture": "amd64",
                    "bootID": "0d588ad5-2590-4581-bf22-12fde4d8855d",
                    "containerRuntimeVersion": "docker://18.6.3",
                    "kernelVersion": "4.19.50-coreos-r1",
                    "kubeProxyVersion": "v1.14.3",
                    "kubeletVersion": "v1.14.3",
                    "machineID": "0ddc0c742bea4698993705304979a4d1",
                    "operatingSystem": "linux",
                    "osImage": "Container Linux by CoreOS 2135.4.0 (Rhyolite)",
                    "systemUUID": "6def17a5-ff20-b448-a511-23a36260d237"
                }
            }
        },
        {
            "apiVersion": "v1",
            "kind": "Node",
            "metadata": {
                "annotations": {
                    "node.alpha.kubernetes.io/ttl": "0",
                    "volumes.kubernetes.io/controller-managed-attach-detach": "true"
                },
                "creationTimestamp": "2019-07-15T09:57:48Z",
                "labels": {
                    "azure-operator.giantswarm.io/version": "2.4.0",
                    "beta.kubernetes.io/arch": "amd64",
                    "beta.kubernetes.io/instance-type": "Standard_A2_v2",
                    "beta.kubernetes.io/os": "linux",
                    "failure-domain.beta.kubernetes.io/region": "westeurope",
                    "failure-domain.beta.kubernetes.io/zone": "0",
                    "giantswarm.io/provider": "azure",
                    "ip": "10.15.1.8",
                    "kubernetes.io/arch": "amd64",
                    "kubernetes.io/hostname": "6iec4-worker-00000j",
                    "kubernetes.io/os": "linux",
                    "kubernetes.io/role": "worker",
                    "node-role.kubernetes.io/worker": "",
                    "node.kubernetes.io/worker": "",
                    "role": "worker"
                },
                "name": "6iec4-worker-00000j",
                "resourceVersion": "87461883",
                "selfLink": "/api/v1/nodes/6iec4-worker-00000j",
                "uid": "0098bc0c-a6e7-11e9-aef7-000d3a43f679"
            },
            "spec": {
                "podCIDR": "10.15.132.0/24",
                "providerID": "azure:///subscriptions/1be3b2e6-497b-45b9-915f-eb35cae23c6a/resourceGroups/6iec4/providers/Microsoft.Compute/virtualMachineScaleSets/6iec4-worker/virtualMachines/19"
            },
            "status": {
                "addresses": [
                    {
                        "address": "10.15.1.8",
                        "type": "InternalIP"
                    },
                    {
                        "address": "6iec4-worker-00000j",
                        "type": "Hostname"
                    }
                ],
                "allocatable": {
                    "attachable-volumes-azure-disk": "4",
                    "cpu": "2",
                    "ephemeral-storage": "28454196Ki",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "3811668Ki",
                    "pods": "110"
                },
                "capacity": {
                    "attachable-volumes-azure-disk": "4",
                    "cpu": "2",
                    "ephemeral-storage": "28454196Ki",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "4016468Ki",
                    "pods": "110"
                },
                "conditions": [
                    {
                        "lastHeartbeatTime": "2019-07-15T09:59:08Z",
                        "lastTransitionTime": "2019-07-15T09:59:08Z",
                        "message": "RouteController created a route",
                        "reason": "RouteCreated",
                        "status": "False",
                        "type": "NetworkUnavailable"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:50Z",
                        "lastTransitionTime": "2019-07-15T09:57:48Z",
                        "message": "kubelet has sufficient memory available",
                        "reason": "KubeletHasSufficientMemory",
                        "status": "False",
                        "type": "MemoryPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:50Z",
                        "lastTransitionTime": "2019-07-15T09:57:48Z",
                        "message": "kubelet has no disk pressure",
                        "reason": "KubeletHasNoDiskPressure",
                        "status": "False",
                        "type": "DiskPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:50Z",
                        "lastTransitionTime": "2019-07-15T09:57:48Z",
                        "message": "kubelet has sufficient PID available",
                        "reason": "KubeletHasSufficientPID",
                        "status": "False",
                        "type": "PIDPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:50Z",
                        "lastTransitionTime": "2019-07-15T09:58:09Z",
                        "message": "kubelet is posting ready status",
                        "reason": "KubeletReady",
                        "status": "True",
                        "type": "Ready"
                    }
                ],
                "daemonEndpoints": {
                    "kubeletEndpoint": {
                        "Port": 10250
                    }
                },
                "images": [
                    {
                        "names": [
                            "quay.io/giantswarm/nginx-ingress-controller@sha256:9086f624cbddb4265a4d8213a0aaeab6795e274004797cc45932354606853406",
                            "quay.io/giantswarm/nginx-ingress-controller:0.24.1"
                        ],
                        "sizeBytes": 631358200
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/hyperkube@sha256:b6cfcecaa8275fd0921bc9c7d839775830d7ab20fd675e2f0fe90c9cec297dcb",
                            "quay.io/giantswarm/hyperkube:v1.14.3"
                        ],
                        "sizeBytes": 603973662
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/sitesearch@sha256:a098b0bb4c6b1474ab8de20ad4a3d77c8ef2e6ee709477c86ef1eb70bfb49c0d",
                            "quay.io/giantswarm/sitesearch:4bce24d4bc8693b57f638af8b797ff9bf738b145"
                        ],
                        "sizeBytes": 157330257
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:a5f2238a37df9aa9819b131a305b953793cb15886f9ebd89984a268344e21cd7",
                            "quay.io/giantswarm/giantswarmio-webapp:ecac78ce5ce988211be5524a9f53e664bb12a218"
                        ],
                        "sizeBytes": 153717267
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:13e08c633812c8a377c13136ab8a1be8fd0af567d6b5dd0b3d3a10ac481369cb",
                            "quay.io/giantswarm/giantswarmio-webapp:649894e9a518865a4f69c670c7a7354bade38576"
                        ],
                        "sizeBytes": 147224012
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:1db7f44195ece5b30506f95dfede4d4b25a91f8ff12955e02a23a8a9f423e324",
                            "quay.io/giantswarm/giantswarmio-webapp:028f4913fbf3fe164c836c5c88b922b685e0042f"
                        ],
                        "sizeBytes": 147145376
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:053a22c29a4463af5491b13c0798ce500c9959e12af68b56b3a2560ba43dc3c9",
                            "quay.io/giantswarm/giantswarmio-webapp:0539c69b5a0e28d52d0a9b6c627ea65ad73a4f3b"
                        ],
                        "sizeBytes": 147136389
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:4ade91acc6dea784fc881f43ced0cabdec3d452ad5f6040915778ac31a869b48",
                            "quay.io/giantswarm/giantswarmio-webapp:9a9e790b1891fa0546f2bdb3dda75e6ab33bf062"
                        ],
                        "sizeBytes": 147128659
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:333d96c3dcc2cf3fd62df953edd4c796914c16b1df7bf208682e56781444b1e5",
                            "quay.io/giantswarm/giantswarmio-webapp:f61c1bb9f8e061d683473b49a2409782e9490e18"
                        ],
                        "sizeBytes": 147120531
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:238ff04669a44dd32742bb24581753b2e55a703f583fbda482e7c6a78fab79d0",
                            "quay.io/giantswarm/giantswarmio-webapp:4948a84ca125a8eda713d21c5b0bd9680a7942b0"
                        ],
                        "sizeBytes": 147116635
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:395d6a4abbeeba7b7a9bf9c8baac13a241d714d4d5d0eddcd4a7f2956f313b9e",
                            "quay.io/giantswarm/giantswarmio-webapp:8f3bfdd13fbbda37a5183406290b848f9fd4a95c"
                        ],
                        "sizeBytes": 147107617
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:038fd8be194c1c8211d0f5579191c664f8aa92173e63ac08d603cfb95e76b0d3",
                            "quay.io/giantswarm/giantswarmio-webapp:19f36b61d0600b95f7205dbc14e4d2054326cf52"
                        ],
                        "sizeBytes": 147082779
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs-indexer@sha256:38cddbe7f78d581a699da0d0d6ad64be0e1754f60a1df3a98ea221d5488eff69",
                            "quay.io/giantswarm/docs-indexer:8ba9d3f3f9cbb6d713e478add6c0f2a55d73242c"
                        ],
                        "sizeBytes": 115198415
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/node@sha256:4735634d37bb1b15611096617e0d23d0bba6ee187f8ea2a2324ae8901e402716",
                            "quay.io/giantswarm/node:v3.5.1"
                        ],
                        "sizeBytes": 75919845
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cni@sha256:7dd62f4d47fbb2920a0cb54bfc45dc56920c7aaf8150b9096dacd91f66b39657",
                            "quay.io/giantswarm/cni:v3.5.1"
                        ],
                        "sizeBytes": 75434816
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:179d9ef76c74ab394b77b1e20cf170ce0c077d4cf69ef2c61f4231410061840a",
                            "quay.io/giantswarm/docs:45071e90b3a0b4fba8ea0b848695e1cc958bea7a"
                        ],
                        "sizeBytes": 69751371
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:4a20140aaf748eea1b8035ba4c47e7ac3cd4ccb4f70d38d2eb0d37cf764fa6c5",
                            "quay.io/giantswarm/docs:4224533c62ffd39d709f9901fd11309289bbdc95"
                        ],
                        "sizeBytes": 69749459
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:a2a1b9897d3329a94af682a083f09a8d9f1c570b494b0015b3919292d7bb55e0",
                            "quay.io/giantswarm/docs:c6330f3f8170d2353f0d6791f9e226c4c2144d9c"
                        ],
                        "sizeBytes": 69749459
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:00035fd9ab10a93fabd68aed0d3658b86b463bb17987f756c2f14b6197ad149a",
                            "quay.io/giantswarm/docs:6357e17a68bcc2722b0292fb7f380207d8f99a7e"
                        ],
                        "sizeBytes": 69749049
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:12d6697067c6fbdc8da84b8c2cd5ca1d691234fa6bc111ede9170241a2dd6d65",
                            "quay.io/giantswarm/docs:b47e6ca9762ff7df9b10abee2ff4819fda8678c9"
                        ],
                        "sizeBytes": 69624200
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
                            "quay.io/giantswarm/net-exporter@sha256:4a3ea165bae3d1c91c36cdb30821ca459b8e3d8c255091bb460af8c21a7a264f",
                            "quay.io/giantswarm/net-exporter:7f178ed2127c0397c9d443bde38beabe435f603a"
                        ],
                        "sizeBytes": 34907651
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:fd333ca6c6b0268bebfca03550a07b959c9fafed2bddf8c559d9ebf050bc855b",
                            "quay.io/giantswarm/net-exporter:1ad313bfd2d835e4b50a606bc510d697976daa93"
                        ],
                        "sizeBytes": 34907614
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:41d00544b36ca1da0a62e326524ebda2e7d94937c81e10872b94306990e55d26",
                            "quay.io/giantswarm/net-exporter:753e1afd06b01c8add545d0673a7dbc5d55880a5"
                        ],
                        "sizeBytes": 32702797
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:600d067e2a257bc036719ae5c0387f00ca64f5ed65ff5bdad8e65e42b2483a2f",
                            "quay.io/giantswarm/net-exporter:dc80ec9f5e2603bc01626fe8ffa7e5f2c2c6f005"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:8bc26239349b8ab0f93e7b00683b79616a5b3424d595a8553086404c13eda46b",
                            "quay.io/giantswarm/net-exporter:90df7e1dce71e7cf6fcb924e93dcdea2306fc6a2"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:80461ff6f6af959cb1a4de14799da04263787b44e9784182a6a35bb7650462bc",
                            "quay.io/giantswarm/net-exporter:7f2d2252afe1856292c2c2748d02a34b7149b2cb"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:32839177290abb6d2f455c09b7d37d9400e29bd30a5a55d8f021195743ce5879",
                            "quay.io/giantswarm/net-exporter:b2fb626131b1ebc941d3fc7f0eac89cd4abfbf41"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:94e1a7e56048e7889a340c953a64ee2f0963cbd52d289515b6c8e291d27a4894",
                            "quay.io/giantswarm/net-exporter:9609ec286a4b55bdaa86fdd8867164e8fe9273d8"
                        ],
                        "sizeBytes": 32559983
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:5421106383aa8be7d7482e3c1da6defab40b0f904c18942652520dba84ec95cb",
                            "quay.io/giantswarm/net-exporter:441558237c0023a6a54d11fc4f1d1b2deec3fca9"
                        ],
                        "sizeBytes": 32559983
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:a6f03bb5dfb4b8fc6a40f54299b49602529f9495e7c9e42f22c6d92b948abee7",
                            "quay.io/giantswarm/net-exporter:812f6c6a537582ea0a1011236f76da15bd4f8146"
                        ],
                        "sizeBytes": 32294939
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:88d962ca443a338309228922b8fc851da6569056c7144efb389701811098a6fe",
                            "quay.io/giantswarm/net-exporter:3a18405a2445d217c423ba4bca915cb1047d20d0"
                        ],
                        "sizeBytes": 32294939
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
                            "quay.io/giantswarm/api-spec@sha256:1e8393bb5686784ff82ea51a313358b087f9c8bfdf39534c9f10212ee554c6c8",
                            "quay.io/giantswarm/api-spec:8bafe1a294707b9d01f98bc391e70f4a2e9886d0"
                        ],
                        "sizeBytes": 23536240
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
                            "quay.io/giantswarm/api-spec@sha256:f333e185ee9472f147b32b0a90f1c3cd1b3c1b8b5bd4c7ef373f5fe7c4c92038",
                            "quay.io/giantswarm/api-spec:df4eea4c399c25c00b0b676f2ee790fdfeec7389"
                        ],
                        "sizeBytes": 22727827
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-nginx@sha256:f3227a01064d1ef02fc9b84e44190ae895c3ec580d062c7d5240ab6dfe775795",
                            "quay.io/giantswarm/giantswarmio-nginx:ab9e03be4cb7ab38457e56aba3466fd116fe2cf4"
                        ],
                        "sizeBytes": 21970854
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/redis@sha256:604e75bf72bb6c80023f500e8d36135d02553e94cdf77eb4fffaa7a549ca10d7",
                            "quay.io/giantswarm/redis:3.2.11-alpine"
                        ],
                        "sizeBytes": 20658687
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/oncall-scheduler@sha256:c3c4612d6646f624830d9dfd1cdc0452fdad90178f10083d08a8a16f8d372639",
                            "quay.io/giantswarm/oncall-scheduler:latest"
                        ],
                        "sizeBytes": 15559099
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
                            "quay.io/giantswarm/cert-exporter@sha256:6ca6e09972957c61964d904ca371fe778a2c710f2cdfca6fc98b29b19f3d92b8",
                            "quay.io/giantswarm/cert-exporter:84f8029d53a89549301fa4c1148571ca9aac8175"
                        ],
                        "sizeBytes": 14489266
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:cbcb2ce76df7bdcc61ed6cc447a8e32a00dd8ad1aab7fea781c0861abfb27bcc",
                            "quay.io/giantswarm/cert-exporter:afbc779ea4835d5fa6e7d3c3450846a3a701b835"
                        ],
                        "sizeBytes": 14489266
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:8dd95249482ccbf77014b47d48d309c083bf9c69ae71caa0ca8c3bc35ae7d93b",
                            "quay.io/giantswarm/cert-exporter:8ade19af74402d643bb90162b23a175728d4d8e9"
                        ],
                        "sizeBytes": 14471305
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:434c272bfc2f17d40b8c583b03b3b02f3d7d7a6c4e7f81d931938a815622f5cc",
                            "quay.io/giantswarm/cert-exporter:e9062664cc0b2d89c80e31c69453b0e5d73a35f5"
                        ],
                        "sizeBytes": 14471305
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:8ff36bdfbc5df23835d5c9f1e0c8f5e34e9192e6bb0d04ddec3313fa26eebad4",
                            "quay.io/giantswarm/cert-exporter:93996b25368a5e72c4e8d08670deef3cb1afaa19"
                        ],
                        "sizeBytes": 14471305
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:82ffa782d890234291bbf4ae1d6ab89ade761656f33e17db8882363103d0e1d8",
                            "quay.io/giantswarm/cert-exporter:a84beb354d6db93301b6957d1a1daf35003bf45a"
                        ],
                        "sizeBytes": 14471305
                    },
                    {
                        "names": [
                            "quay.io/jetstack/cert-manager-acmesolver@sha256:77e99c22313379d1d4b9cf3faa5a33fba9debdccac92a4a3848fde3f953f93bb",
                            "quay.io/jetstack/cert-manager-acmesolver:v0.4.1"
                        ],
                        "sizeBytes": 10839715
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/web-assets@sha256:0da22cee22c4a4ca98a0586550f0deaecde311c2097e9defb2e1bf73a8f9bac7",
                            "quay.io/giantswarm/web-assets:d458ce96b9dc9ccc50c8640190f7bc1d75077478"
                        ],
                        "sizeBytes": 9337158
                    }
                ],
                "nodeInfo": {
                    "architecture": "amd64",
                    "bootID": "f79eff7e-99dc-4e0a-9bbd-dbf2e30f9145",
                    "containerRuntimeVersion": "docker://18.6.3",
                    "kernelVersion": "4.19.50-coreos-r1",
                    "kubeProxyVersion": "v1.14.3",
                    "kubeletVersion": "v1.14.3",
                    "machineID": "7306f9854a684a1c840039b47be9fb3d",
                    "operatingSystem": "linux",
                    "osImage": "Container Linux by CoreOS 2135.4.0 (Rhyolite)",
                    "systemUUID": "ef99aeca-56cd-2049-9cd4-d0b8a7efdb6a"
                }
            }
        },
        {
            "apiVersion": "v1",
            "kind": "Node",
            "metadata": {
                "annotations": {
                    "node.alpha.kubernetes.io/ttl": "0",
                    "volumes.kubernetes.io/controller-managed-attach-detach": "true"
                },
                "creationTimestamp": "2019-07-15T10:39:52Z",
                "labels": {
                    "azure-operator.giantswarm.io/version": "2.4.0",
                    "beta.kubernetes.io/arch": "amd64",
                    "beta.kubernetes.io/instance-type": "Standard_A2_v2",
                    "beta.kubernetes.io/os": "linux",
                    "failure-domain.beta.kubernetes.io/region": "westeurope",
                    "failure-domain.beta.kubernetes.io/zone": "1",
                    "giantswarm.io/provider": "azure",
                    "ip": "10.15.1.9",
                    "kubernetes.io/arch": "amd64",
                    "kubernetes.io/hostname": "6iec4-worker-00000k",
                    "kubernetes.io/os": "linux",
                    "kubernetes.io/role": "worker",
                    "node-role.kubernetes.io/worker": "",
                    "node.kubernetes.io/worker": "",
                    "role": "worker"
                },
                "name": "6iec4-worker-00000k",
                "resourceVersion": "87461880",
                "selfLink": "/api/v1/nodes/6iec4-worker-00000k",
                "uid": "e0b4c51b-a6ec-11e9-aef7-000d3a43f679"
            },
            "spec": {
                "podCIDR": "10.15.133.0/24",
                "providerID": "azure:///subscriptions/1be3b2e6-497b-45b9-915f-eb35cae23c6a/resourceGroups/6iec4/providers/Microsoft.Compute/virtualMachineScaleSets/6iec4-worker/virtualMachines/20"
            },
            "status": {
                "addresses": [
                    {
                        "address": "10.15.1.9",
                        "type": "InternalIP"
                    },
                    {
                        "address": "6iec4-worker-00000k",
                        "type": "Hostname"
                    }
                ],
                "allocatable": {
                    "attachable-volumes-azure-disk": "4",
                    "cpu": "2",
                    "ephemeral-storage": "28454196Ki",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "3811668Ki",
                    "pods": "110"
                },
                "capacity": {
                    "attachable-volumes-azure-disk": "4",
                    "cpu": "2",
                    "ephemeral-storage": "28454196Ki",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "4016468Ki",
                    "pods": "110"
                },
                "conditions": [
                    {
                        "lastHeartbeatTime": "2019-07-15T10:41:08Z",
                        "lastTransitionTime": "2019-07-15T10:41:08Z",
                        "message": "RouteController created a route",
                        "reason": "RouteCreated",
                        "status": "False",
                        "type": "NetworkUnavailable"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:49Z",
                        "lastTransitionTime": "2019-07-15T10:39:52Z",
                        "message": "kubelet has sufficient memory available",
                        "reason": "KubeletHasSufficientMemory",
                        "status": "False",
                        "type": "MemoryPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:49Z",
                        "lastTransitionTime": "2019-07-15T10:39:52Z",
                        "message": "kubelet has no disk pressure",
                        "reason": "KubeletHasNoDiskPressure",
                        "status": "False",
                        "type": "DiskPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:49Z",
                        "lastTransitionTime": "2019-07-15T10:39:52Z",
                        "message": "kubelet has sufficient PID available",
                        "reason": "KubeletHasSufficientPID",
                        "status": "False",
                        "type": "PIDPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:49Z",
                        "lastTransitionTime": "2019-07-15T10:40:22Z",
                        "message": "kubelet is posting ready status",
                        "reason": "KubeletReady",
                        "status": "True",
                        "type": "Ready"
                    }
                ],
                "daemonEndpoints": {
                    "kubeletEndpoint": {
                        "Port": 10250
                    }
                },
                "images": [
                    {
                        "names": [
                            "quay.io/giantswarm/nginx-ingress-controller@sha256:9086f624cbddb4265a4d8213a0aaeab6795e274004797cc45932354606853406",
                            "quay.io/giantswarm/nginx-ingress-controller:0.24.1"
                        ],
                        "sizeBytes": 631358200
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/hyperkube@sha256:b6cfcecaa8275fd0921bc9c7d839775830d7ab20fd675e2f0fe90c9cec297dcb",
                            "quay.io/giantswarm/hyperkube:v1.14.3"
                        ],
                        "sizeBytes": 603973662
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:097ce3ee8be6f7557e546a8bdbbf993b66c9f69d84869afe3872b00ef5459c81",
                            "quay.io/giantswarm/giantswarmio-webapp:4a112738beef5ba5fe95458c12b96f291078fb28"
                        ],
                        "sizeBytes": 153632225
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:9f314d5a95524530c3e46eb4139dccb7581eca834be23cdd1108e00bd7056052",
                            "quay.io/giantswarm/giantswarmio-webapp:7473ec5666fd310dcbca17b23632edefbd62ca3b"
                        ],
                        "sizeBytes": 153632160
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:14c3c780fbb5081cd8c45447242b25374bc7adfddfc26b0026d3ced8ca3db3db",
                            "quay.io/giantswarm/giantswarmio-webapp:81090449165423c36387c48f5cd639d2e24d6274"
                        ],
                        "sizeBytes": 147214080
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:191cb859d624cb0a5e47642ba8dfb5378532c2192f8e565d028afae5807162c1",
                            "quay.io/giantswarm/giantswarmio-webapp:ab3e26030c29e7a1d1987ab8ff787a1b6513dd81"
                        ],
                        "sizeBytes": 147165646
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:1db7f44195ece5b30506f95dfede4d4b25a91f8ff12955e02a23a8a9f423e324",
                            "quay.io/giantswarm/giantswarmio-webapp:028f4913fbf3fe164c836c5c88b922b685e0042f"
                        ],
                        "sizeBytes": 147145376
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:df0d937111e3b4f87fe468f3c420288da04928460fb14a879270658a5d855127",
                            "quay.io/giantswarm/giantswarmio-webapp:f7bbf78c49b83e728e644a09e857ebebb2366d93"
                        ],
                        "sizeBytes": 147145121
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:5f1805ed757a261be995e0d70971c595484fc0ede743f35d29236912a94923ae",
                            "quay.io/giantswarm/giantswarmio-webapp:960ab74ae0e379b69188e9bbd455ff85656af9ab"
                        ],
                        "sizeBytes": 147135789
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:7760315471d5c7499d1aa354ca389a6f8e761f4d5d95d68280c65b0a2bba034a",
                            "quay.io/giantswarm/giantswarmio-webapp:1bee90c89c7c1e46bd02acb8a9df8bac005ccd1e"
                        ],
                        "sizeBytes": 147134765
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:69a83fa3a5a40ec68c41ec1b43d46e007eed1a07017b79ba58247111ca3857ff",
                            "quay.io/giantswarm/giantswarmio-webapp:03c853e56743aa585a7ec19e4cdac833e89d49f2"
                        ],
                        "sizeBytes": 147132682
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:4ade91acc6dea784fc881f43ced0cabdec3d452ad5f6040915778ac31a869b48",
                            "quay.io/giantswarm/giantswarmio-webapp:9a9e790b1891fa0546f2bdb3dda75e6ab33bf062"
                        ],
                        "sizeBytes": 147128659
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:9c479b566994ebfd7b899a6800522b183869ce3d2110160c54c00b51950e757b",
                            "quay.io/giantswarm/giantswarmio-webapp:a81b18d438212d8497f54ae146add487a52d1c66"
                        ],
                        "sizeBytes": 147128122
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:ec6ccc8abdaf1ec31983247f34ca9fa827f8fc58d1d65c458dc556d1e0635dac",
                            "quay.io/giantswarm/giantswarmio-webapp:8c7a9b1e31d687af7fa1e6b5d75334a3a6200218"
                        ],
                        "sizeBytes": 147111609
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:c91557f5e7e09ec14cba05b142de249a7103f988fa44b66d75e25ef8c6b21134",
                            "quay.io/giantswarm/giantswarmio-webapp:6c0bac466231e3831fa1b952b4c653be67d6f30a"
                        ],
                        "sizeBytes": 147111540
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:395d6a4abbeeba7b7a9bf9c8baac13a241d714d4d5d0eddcd4a7f2956f313b9e",
                            "quay.io/giantswarm/giantswarmio-webapp:8f3bfdd13fbbda37a5183406290b848f9fd4a95c"
                        ],
                        "sizeBytes": 147107617
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:3a97718451fea83b498b7b6dec4e7f3c75fc5e57caf881de8fe6267e8172b17b",
                            "quay.io/giantswarm/giantswarmio-webapp:c2d0e1616565bfec161378aae0d143c87af78046"
                        ],
                        "sizeBytes": 147102708
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:038fd8be194c1c8211d0f5579191c664f8aa92173e63ac08d603cfb95e76b0d3",
                            "quay.io/giantswarm/giantswarmio-webapp:19f36b61d0600b95f7205dbc14e4d2054326cf52"
                        ],
                        "sizeBytes": 147082779
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs-indexer@sha256:38cddbe7f78d581a699da0d0d6ad64be0e1754f60a1df3a98ea221d5488eff69",
                            "quay.io/giantswarm/docs-indexer:8ba9d3f3f9cbb6d713e478add6c0f2a55d73242c"
                        ],
                        "sizeBytes": 115198415
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/node@sha256:4735634d37bb1b15611096617e0d23d0bba6ee187f8ea2a2324ae8901e402716",
                            "quay.io/giantswarm/node:v3.5.1"
                        ],
                        "sizeBytes": 75919845
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cni@sha256:7dd62f4d47fbb2920a0cb54bfc45dc56920c7aaf8150b9096dacd91f66b39657",
                            "quay.io/giantswarm/cni:v3.5.1"
                        ],
                        "sizeBytes": 75434816
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:a2a1b9897d3329a94af682a083f09a8d9f1c570b494b0015b3919292d7bb55e0",
                            "quay.io/giantswarm/docs:c6330f3f8170d2353f0d6791f9e226c4c2144d9c"
                        ],
                        "sizeBytes": 69749459
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:8b365eeaab9ccbb542449869b0d5f334cf794561daf9d62450be91a309484b18",
                            "quay.io/giantswarm/docs:8b09158ed582ad5326cd7661452e4cb36f3ac4c0"
                        ],
                        "sizeBytes": 69606280
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:b21deb4691862992e7de5752619cccadfc5db5ae4016e7565a1d2192dca03458",
                            "quay.io/giantswarm/docs:5ec7929c0c3c401ea0c5a9fab850d26022bcc45a"
                        ],
                        "sizeBytes": 69605018
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:d2c7b8f094bdc85925a0b32d14e0daa0f0fc698c7fe83e839038961a5f46a94c",
                            "quay.io/giantswarm/docs:35ab8cf866fe625abbaa195333a111e034cae19c"
                        ],
                        "sizeBytes": 69121282
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/external-dns@sha256:f2c54f3b288d26195b8d0286aec7ab61ec273356552d4a137a86f34a1c920a79",
                            "quay.io/giantswarm/external-dns:v0.5.11"
                        ],
                        "sizeBytes": 52422138
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
                            "quay.io/giantswarm/net-exporter@sha256:4a3ea165bae3d1c91c36cdb30821ca459b8e3d8c255091bb460af8c21a7a264f",
                            "quay.io/giantswarm/net-exporter:7f178ed2127c0397c9d443bde38beabe435f603a"
                        ],
                        "sizeBytes": 34907651
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:fd333ca6c6b0268bebfca03550a07b959c9fafed2bddf8c559d9ebf050bc855b",
                            "quay.io/giantswarm/net-exporter:1ad313bfd2d835e4b50a606bc510d697976daa93"
                        ],
                        "sizeBytes": 34907614
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:41d00544b36ca1da0a62e326524ebda2e7d94937c81e10872b94306990e55d26",
                            "quay.io/giantswarm/net-exporter:753e1afd06b01c8add545d0673a7dbc5d55880a5"
                        ],
                        "sizeBytes": 32702797
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:600d067e2a257bc036719ae5c0387f00ca64f5ed65ff5bdad8e65e42b2483a2f",
                            "quay.io/giantswarm/net-exporter:dc80ec9f5e2603bc01626fe8ffa7e5f2c2c6f005"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:80461ff6f6af959cb1a4de14799da04263787b44e9784182a6a35bb7650462bc",
                            "quay.io/giantswarm/net-exporter:7f2d2252afe1856292c2c2748d02a34b7149b2cb"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:32839177290abb6d2f455c09b7d37d9400e29bd30a5a55d8f021195743ce5879",
                            "quay.io/giantswarm/net-exporter:b2fb626131b1ebc941d3fc7f0eac89cd4abfbf41"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:8bc26239349b8ab0f93e7b00683b79616a5b3424d595a8553086404c13eda46b",
                            "quay.io/giantswarm/net-exporter:90df7e1dce71e7cf6fcb924e93dcdea2306fc6a2"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:94e1a7e56048e7889a340c953a64ee2f0963cbd52d289515b6c8e291d27a4894",
                            "quay.io/giantswarm/net-exporter:9609ec286a4b55bdaa86fdd8867164e8fe9273d8"
                        ],
                        "sizeBytes": 32559983
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:5421106383aa8be7d7482e3c1da6defab40b0f904c18942652520dba84ec95cb",
                            "quay.io/giantswarm/net-exporter:441558237c0023a6a54d11fc4f1d1b2deec3fca9"
                        ],
                        "sizeBytes": 32559983
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:a6f03bb5dfb4b8fc6a40f54299b49602529f9495e7c9e42f22c6d92b948abee7",
                            "quay.io/giantswarm/net-exporter:812f6c6a537582ea0a1011236f76da15bd4f8146"
                        ],
                        "sizeBytes": 32294939
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:88d962ca443a338309228922b8fc851da6569056c7144efb389701811098a6fe",
                            "quay.io/giantswarm/net-exporter:3a18405a2445d217c423ba4bca915cb1047d20d0"
                        ],
                        "sizeBytes": 32294939
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
                            "quay.io/giantswarm/api-spec@sha256:fea4a288f39b4c6a32894aba36ae42d52c183b1f4f99fef7303574498b1897ed",
                            "quay.io/giantswarm/api-spec:d485fbcc933b8ce5c3bbff3c5bef185e1b8dca51"
                        ],
                        "sizeBytes": 23539782
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
                            "quay.io/giantswarm/api-spec@sha256:f333e185ee9472f147b32b0a90f1c3cd1b3c1b8b5bd4c7ef373f5fe7c4c92038",
                            "quay.io/giantswarm/api-spec:df4eea4c399c25c00b0b676f2ee790fdfeec7389"
                        ],
                        "sizeBytes": 22727827
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/redis@sha256:604e75bf72bb6c80023f500e8d36135d02553e94cdf77eb4fffaa7a549ca10d7",
                            "quay.io/giantswarm/redis:3.2.11-alpine"
                        ],
                        "sizeBytes": 20658687
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs-proxy@sha256:a1a1d9a8a70b4d10b6956a34c297fb2ec81d0cd0eeeecc26b0213fe5c953a0b5",
                            "quay.io/giantswarm/docs-proxy:f68e0e1b306452505ebcd901863756f6bacee3dd"
                        ],
                        "sizeBytes": 17408306
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/oncall-scheduler@sha256:c3c4612d6646f624830d9dfd1cdc0452fdad90178f10083d08a8a16f8d372639",
                            "quay.io/giantswarm/oncall-scheduler:latest"
                        ],
                        "sizeBytes": 15559099
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
                            "quay.io/giantswarm/cert-exporter@sha256:cbcb2ce76df7bdcc61ed6cc447a8e32a00dd8ad1aab7fea781c0861abfb27bcc",
                            "quay.io/giantswarm/cert-exporter:afbc779ea4835d5fa6e7d3c3450846a3a701b835"
                        ],
                        "sizeBytes": 14489266
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:6ca6e09972957c61964d904ca371fe778a2c710f2cdfca6fc98b29b19f3d92b8",
                            "quay.io/giantswarm/cert-exporter:84f8029d53a89549301fa4c1148571ca9aac8175"
                        ],
                        "sizeBytes": 14489266
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cert-exporter@sha256:8dd95249482ccbf77014b47d48d309c083bf9c69ae71caa0ca8c3bc35ae7d93b",
                            "quay.io/giantswarm/cert-exporter:8ade19af74402d643bb90162b23a175728d4d8e9"
                        ],
                        "sizeBytes": 14471305
                    }
                ],
                "nodeInfo": {
                    "architecture": "amd64",
                    "bootID": "26d03867-50e7-49ad-80a5-2c4bcb5c0676",
                    "containerRuntimeVersion": "docker://18.6.3",
                    "kernelVersion": "4.19.50-coreos-r1",
                    "kubeProxyVersion": "v1.14.3",
                    "kubeletVersion": "v1.14.3",
                    "machineID": "83aa371dfab6418f9861d4e85885c982",
                    "operatingSystem": "linux",
                    "osImage": "Container Linux by CoreOS 2135.4.0 (Rhyolite)",
                    "systemUUID": "a226ae85-be22-e14b-9ebd-845ef572cc20"
                }
            }
        },
        {
            "apiVersion": "v1",
            "kind": "Node",
            "metadata": {
                "annotations": {
                    "node.alpha.kubernetes.io/ttl": "0",
                    "volumes.kubernetes.io/controller-managed-attach-detach": "true"
                },
                "creationTimestamp": "2019-07-15T10:46:30Z",
                "labels": {
                    "azure-operator.giantswarm.io/version": "2.4.0",
                    "beta.kubernetes.io/arch": "amd64",
                    "beta.kubernetes.io/instance-type": "Standard_A2_v2",
                    "beta.kubernetes.io/os": "linux",
                    "failure-domain.beta.kubernetes.io/region": "westeurope",
                    "failure-domain.beta.kubernetes.io/zone": "2",
                    "giantswarm.io/provider": "azure",
                    "ip": "10.15.1.11",
                    "kubernetes.io/arch": "amd64",
                    "kubernetes.io/hostname": "6iec4-worker-00000m",
                    "kubernetes.io/os": "linux",
                    "kubernetes.io/role": "worker",
                    "node-role.kubernetes.io/worker": "",
                    "node.kubernetes.io/worker": "",
                    "role": "worker"
                },
                "name": "6iec4-worker-00000m",
                "resourceVersion": "87461885",
                "selfLink": "/api/v1/nodes/6iec4-worker-00000m",
                "uid": "cdc2ca54-a6ed-11e9-aef7-000d3a43f679"
            },
            "spec": {
                "podCIDR": "10.15.134.0/24",
                "providerID": "azure:///subscriptions/1be3b2e6-497b-45b9-915f-eb35cae23c6a/resourceGroups/6iec4/providers/Microsoft.Compute/virtualMachineScaleSets/6iec4-worker/virtualMachines/22"
            },
            "status": {
                "addresses": [
                    {
                        "address": "10.15.1.11",
                        "type": "InternalIP"
                    },
                    {
                        "address": "6iec4-worker-00000m",
                        "type": "Hostname"
                    }
                ],
                "allocatable": {
                    "attachable-volumes-azure-disk": "4",
                    "cpu": "2",
                    "ephemeral-storage": "28454196Ki",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "3811668Ki",
                    "pods": "110"
                },
                "capacity": {
                    "attachable-volumes-azure-disk": "4",
                    "cpu": "2",
                    "ephemeral-storage": "28454196Ki",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "4016468Ki",
                    "pods": "110"
                },
                "conditions": [
                    {
                        "lastHeartbeatTime": "2019-07-15T10:47:09Z",
                        "lastTransitionTime": "2019-07-15T10:47:09Z",
                        "message": "RouteController created a route",
                        "reason": "RouteCreated",
                        "status": "False",
                        "type": "NetworkUnavailable"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:50Z",
                        "lastTransitionTime": "2019-07-15T10:46:30Z",
                        "message": "kubelet has sufficient memory available",
                        "reason": "KubeletHasSufficientMemory",
                        "status": "False",
                        "type": "MemoryPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:50Z",
                        "lastTransitionTime": "2019-07-15T10:46:30Z",
                        "message": "kubelet has no disk pressure",
                        "reason": "KubeletHasNoDiskPressure",
                        "status": "False",
                        "type": "DiskPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:50Z",
                        "lastTransitionTime": "2019-07-15T10:46:30Z",
                        "message": "kubelet has sufficient PID available",
                        "reason": "KubeletHasSufficientPID",
                        "status": "False",
                        "type": "PIDPressure"
                    },
                    {
                        "lastHeartbeatTime": "2019-10-01T09:09:50Z",
                        "lastTransitionTime": "2019-07-15T10:46:50Z",
                        "message": "kubelet is posting ready status",
                        "reason": "KubeletReady",
                        "status": "True",
                        "type": "Ready"
                    }
                ],
                "daemonEndpoints": {
                    "kubeletEndpoint": {
                        "Port": 10250
                    }
                },
                "images": [
                    {
                        "names": [
                            "quay.io/giantswarm/nginx-ingress-controller@sha256:9086f624cbddb4265a4d8213a0aaeab6795e274004797cc45932354606853406",
                            "quay.io/giantswarm/nginx-ingress-controller:0.24.1"
                        ],
                        "sizeBytes": 631358200
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/hyperkube@sha256:b6cfcecaa8275fd0921bc9c7d839775830d7ab20fd675e2f0fe90c9cec297dcb",
                            "quay.io/giantswarm/hyperkube:v1.14.3"
                        ],
                        "sizeBytes": 603973662
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/sitesearch@sha256:a098b0bb4c6b1474ab8de20ad4a3d77c8ef2e6ee709477c86ef1eb70bfb49c0d",
                            "quay.io/giantswarm/sitesearch:4bce24d4bc8693b57f638af8b797ff9bf738b145"
                        ],
                        "sizeBytes": 157330257
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:cf17afb6c6b97d87570d17e90464186d8fbfc722192d40db7eb03993a7cf5d46",
                            "quay.io/giantswarm/giantswarmio-webapp:540e23d2208d41e6c5bc8b0861b3da60729e20ec"
                        ],
                        "sizeBytes": 153717050
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:9f314d5a95524530c3e46eb4139dccb7581eca834be23cdd1108e00bd7056052",
                            "quay.io/giantswarm/giantswarmio-webapp:7473ec5666fd310dcbca17b23632edefbd62ca3b"
                        ],
                        "sizeBytes": 153632160
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:191cb859d624cb0a5e47642ba8dfb5378532c2192f8e565d028afae5807162c1",
                            "quay.io/giantswarm/giantswarmio-webapp:ab3e26030c29e7a1d1987ab8ff787a1b6513dd81"
                        ],
                        "sizeBytes": 147165646
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:7760315471d5c7499d1aa354ca389a6f8e761f4d5d95d68280c65b0a2bba034a",
                            "quay.io/giantswarm/giantswarmio-webapp:1bee90c89c7c1e46bd02acb8a9df8bac005ccd1e"
                        ],
                        "sizeBytes": 147134765
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:dfcb79b116e07bdf5b18ba6c83f16b5123ce43e364da98853d0ceb6d91899ff6",
                            "quay.io/giantswarm/giantswarmio-webapp:ac0f3a19cfc4ab2aa66be0506ac9268c97102775"
                        ],
                        "sizeBytes": 147132137
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-webapp@sha256:238ff04669a44dd32742bb24581753b2e55a703f583fbda482e7c6a78fab79d0",
                            "quay.io/giantswarm/giantswarmio-webapp:4948a84ca125a8eda713d21c5b0bd9680a7942b0"
                        ],
                        "sizeBytes": 147116635
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs-indexer@sha256:38cddbe7f78d581a699da0d0d6ad64be0e1754f60a1df3a98ea221d5488eff69",
                            "quay.io/giantswarm/docs-indexer:8ba9d3f3f9cbb6d713e478add6c0f2a55d73242c"
                        ],
                        "sizeBytes": 115198415
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/node@sha256:4735634d37bb1b15611096617e0d23d0bba6ee187f8ea2a2324ae8901e402716",
                            "quay.io/giantswarm/node:v3.5.1"
                        ],
                        "sizeBytes": 75919845
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cni@sha256:7dd62f4d47fbb2920a0cb54bfc45dc56920c7aaf8150b9096dacd91f66b39657",
                            "quay.io/giantswarm/cni:v3.5.1"
                        ],
                        "sizeBytes": 75434816
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:fb7eaba9a6cb8daedadec8cd6fa6504c5cfde93be7cf947df8c77272b74a2d12",
                            "quay.io/giantswarm/docs:eed835d6dcc366b4ce9b3abb8fa0b42bca3d4b65"
                        ],
                        "sizeBytes": 69751927
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:c944c511e8392298f9a7a529daacd80bf7ebf731c93322f2482090f92d100a2c",
                            "quay.io/giantswarm/docs:766d83031cb60787290adf065ede7126e08bd964"
                        ],
                        "sizeBytes": 69750504
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:c7ecc7c809a0cbfa2f7533dcdc434b05e9f17a2dfec0059f4acf312528ff1982",
                            "quay.io/giantswarm/docs:6ff58ccfb6541e89fa4e541792d1eef2d16cd574"
                        ],
                        "sizeBytes": 69749376
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:1deac5cb09a7be41c70a6d7ec9c9e6f1a1fcb8510ab31093178293ba59bb0df8",
                            "quay.io/giantswarm/docs:56adfdb23b183848a8f38835bd2d84054a09a16a"
                        ],
                        "sizeBytes": 69605018
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:ebb5f2a40a0e9f7f4008280ace9381cc00c5c2ed30a63368a5628c7fe5a546b2",
                            "quay.io/giantswarm/docs:a319fdd9880a9c9620ddc36e0a602fad48fbaffc"
                        ],
                        "sizeBytes": 69220401
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/docs@sha256:d2c7b8f094bdc85925a0b32d14e0daa0f0fc698c7fe83e839038961a5f46a94c",
                            "quay.io/giantswarm/docs:35ab8cf866fe625abbaa195333a111e034cae19c"
                        ],
                        "sizeBytes": 69121282
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/external-dns@sha256:f2c54f3b288d26195b8d0286aec7ab61ec273356552d4a137a86f34a1c920a79",
                            "quay.io/giantswarm/external-dns:v0.5.11"
                        ],
                        "sizeBytes": 52422138
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
                            "quay.io/giantswarm/net-exporter@sha256:4a3ea165bae3d1c91c36cdb30821ca459b8e3d8c255091bb460af8c21a7a264f",
                            "quay.io/giantswarm/net-exporter:7f178ed2127c0397c9d443bde38beabe435f603a"
                        ],
                        "sizeBytes": 34907651
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:fd333ca6c6b0268bebfca03550a07b959c9fafed2bddf8c559d9ebf050bc855b",
                            "quay.io/giantswarm/net-exporter:1ad313bfd2d835e4b50a606bc510d697976daa93"
                        ],
                        "sizeBytes": 34907614
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:41d00544b36ca1da0a62e326524ebda2e7d94937c81e10872b94306990e55d26",
                            "quay.io/giantswarm/net-exporter:753e1afd06b01c8add545d0673a7dbc5d55880a5"
                        ],
                        "sizeBytes": 32702797
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:80461ff6f6af959cb1a4de14799da04263787b44e9784182a6a35bb7650462bc",
                            "quay.io/giantswarm/net-exporter:7f2d2252afe1856292c2c2748d02a34b7149b2cb"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:8bc26239349b8ab0f93e7b00683b79616a5b3424d595a8553086404c13eda46b",
                            "quay.io/giantswarm/net-exporter:90df7e1dce71e7cf6fcb924e93dcdea2306fc6a2"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:32839177290abb6d2f455c09b7d37d9400e29bd30a5a55d8f021195743ce5879",
                            "quay.io/giantswarm/net-exporter:b2fb626131b1ebc941d3fc7f0eac89cd4abfbf41"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:600d067e2a257bc036719ae5c0387f00ca64f5ed65ff5bdad8e65e42b2483a2f",
                            "quay.io/giantswarm/net-exporter:dc80ec9f5e2603bc01626fe8ffa7e5f2c2c6f005"
                        ],
                        "sizeBytes": 32577822
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:94e1a7e56048e7889a340c953a64ee2f0963cbd52d289515b6c8e291d27a4894",
                            "quay.io/giantswarm/net-exporter:9609ec286a4b55bdaa86fdd8867164e8fe9273d8"
                        ],
                        "sizeBytes": 32559983
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:5421106383aa8be7d7482e3c1da6defab40b0f904c18942652520dba84ec95cb",
                            "quay.io/giantswarm/net-exporter:441558237c0023a6a54d11fc4f1d1b2deec3fca9"
                        ],
                        "sizeBytes": 32559983
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:a6f03bb5dfb4b8fc6a40f54299b49602529f9495e7c9e42f22c6d92b948abee7",
                            "quay.io/giantswarm/net-exporter:812f6c6a537582ea0a1011236f76da15bd4f8146"
                        ],
                        "sizeBytes": 32294939
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/net-exporter@sha256:88d962ca443a338309228922b8fc851da6569056c7144efb389701811098a6fe",
                            "quay.io/giantswarm/net-exporter:3a18405a2445d217c423ba4bca915cb1047d20d0"
                        ],
                        "sizeBytes": 32294939
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
                            "quay.io/giantswarm/api-spec@sha256:ea4fc9a60f411a427a879e414cab0d15de69c0de278b221b59a822b2a7d0cc37",
                            "quay.io/giantswarm/api-spec:a3d9968a6d82afaa061856e11d55b5f894bebad5"
                        ],
                        "sizeBytes": 23536484
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/api-spec@sha256:1e8393bb5686784ff82ea51a313358b087f9c8bfdf39534c9f10212ee554c6c8",
                            "quay.io/giantswarm/api-spec:8bafe1a294707b9d01f98bc391e70f4a2e9886d0"
                        ],
                        "sizeBytes": 23536240
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/api-spec@sha256:312c9da3cb762618d3760742b036ec1a82ba91e3e16938bc64c18f4ec2fc0369",
                            "quay.io/giantswarm/api-spec:67949a50b8bd55f956c56494bbb3443bc70bab55"
                        ],
                        "sizeBytes": 23536020
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/api-spec@sha256:7032f467f4f762958fb9626ee22713e2d216994d405e6e4f63d9b53500f179e8",
                            "quay.io/giantswarm/api-spec:85ad5086cee3b02960715a555dc0d56bfbb45f63"
                        ],
                        "sizeBytes": 23527771
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
                            "quay.io/giantswarm/giantswarmio-nginx@sha256:023261b50767a7b9b8b8fd835246570d368e2c0d2a119eb0e453c2671a47b322",
                            "quay.io/giantswarm/giantswarmio-nginx:0ee7977c8e10b41b513c5edbd391a44ba11bb119"
                        ],
                        "sizeBytes": 21970970
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-nginx@sha256:47142466cffa86dfa99cbd6e3510e79acb0d5e23d48fcdd7ffc3a430bbac960d",
                            "quay.io/giantswarm/giantswarmio-nginx:1a5c26c638caa7e1c5f6d3c1acf9f8b2620b4b75"
                        ],
                        "sizeBytes": 21970970
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/giantswarmio-nginx@sha256:f3227a01064d1ef02fc9b84e44190ae895c3ec580d062c7d5240ab6dfe775795",
                            "quay.io/giantswarm/giantswarmio-nginx:ab9e03be4cb7ab38457e56aba3466fd116fe2cf4"
                        ],
                        "sizeBytes": 21970854
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/redis@sha256:604e75bf72bb6c80023f500e8d36135d02553e94cdf77eb4fffaa7a549ca10d7",
                            "quay.io/giantswarm/redis:3.2.11-alpine"
                        ],
                        "sizeBytes": 20658687
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/auto-oncall@sha256:0f04fc04fa7f2d16974a421f4e57d2d98f7393e48770bf5e6599f0ec4350c207",
                            "quay.io/giantswarm/auto-oncall:8083aaf1ec5ee00705eec2053f3628f4bca8ebbb"
                        ],
                        "sizeBytes": 18745529
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/auto-oncall@sha256:2f083bcefe4ed729563ba92ac8fbf1d0487107b541efebd1905a2a1936483917",
                            "quay.io/giantswarm/auto-oncall:dbef151a94d3ce79bdedab06fcf32488056d33f5"
                        ],
                        "sizeBytes": 17593873
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/auto-oncall@sha256:f480a0632d535eb3d6581cd153e6a38955a1cdeb62b3f269c3bf52ce744a3a87",
                            "quay.io/giantswarm/auto-oncall:fa1f92eeba8d1ed258d67188cf481af2820039df"
                        ],
                        "sizeBytes": 17583891
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/oncall-scheduler@sha256:c3c4612d6646f624830d9dfd1cdc0452fdad90178f10083d08a8a16f8d372639",
                            "quay.io/giantswarm/oncall-scheduler:latest"
                        ],
                        "sizeBytes": 15559099
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
                            "quay.io/giantswarm/cert-exporter@sha256:cbcb2ce76df7bdcc61ed6cc447a8e32a00dd8ad1aab7fea781c0861abfb27bcc",
                            "quay.io/giantswarm/cert-exporter:afbc779ea4835d5fa6e7d3c3450846a3a701b835"
                        ],
                        "sizeBytes": 14489266
                    }
                ],
                "nodeInfo": {
                    "architecture": "amd64",
                    "bootID": "ca2cb440-3431-44f2-867e-4459643d435f",
                    "containerRuntimeVersion": "docker://18.6.3",
                    "kernelVersion": "4.19.50-coreos-r1",
                    "kubeProxyVersion": "v1.14.3",
                    "kubeletVersion": "v1.14.3",
                    "machineID": "c276027411724184b24cb6a2e26fd8c1",
                    "operatingSystem": "linux",
                    "osImage": "Container Linux by CoreOS 2135.4.0 (Rhyolite)",
                    "systemUUID": "62e2daf9-bf40-ca4a-a87e-04a34d280d82"
                }
            }
        }
    ],
    "kind": "List",
    "metadata": {
        "resourceVersion": "",
        "selfLink": ""
    }
}
`
