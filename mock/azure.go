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
            "apiVersion": "provider.giantswarm.io/v1alpha1",
            "kind": "AzureConfig",
            "metadata": {
                "creationTimestamp": "2019-08-14T12:46:24Z",
                "finalizers": [
                    "operatorkit.giantswarm.io/azure-operator"
                ],
                "generation": 1,
                "name": "0xhzq",
                "namespace": "default",
                "resourceVersion": "141247887",
                "selfLink": "/apis/provider.giantswarm.io/v1alpha1/namespaces/default/azureconfigs/0xhzq",
                "uid": "8679c86e-be91-11e9-ac0f-000d3a43f713"
            },
            "spec": {
                "azure": {
                    "dnsZones": {
                        "api": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        },
                        "etcd": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        },
                        "ingress": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        }
                    },
                    "masters": [
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_D2s_v3"
                        }
                    ],
                    "virtualNetwork": {
                        "calicoSubnetCIDR": "10.14.128.0/17",
                        "cidr": "10.14.0.0/16",
                        "masterSubnetCIDR": "10.14.0.0/24",
                        "workerSubnetCIDR": "10.14.1.0/24"
                    },
                    "workers": [
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_A2_v2"
                        }
                    ]
                },
                "cluster": {
                    "calico": {
                        "cidr": 0,
                        "mtu": 1430,
                        "subnet": ""
                    },
                    "docker": {
                        "daemon": {
                            "cidr": "172.17.0.1/16"
                        }
                    },
                    "etcd": {
                        "altNames": "",
                        "domain": "etcd.0xhzq.k8s.gollum.westeurope.azure.gigantic.io",
                        "port": 2379,
                        "prefix": "giantswarm.io"
                    },
                    "id": "0xhzq",
                    "kubernetes": {
                        "api": {
                            "clusterIPRange": "172.31.0.0/16",
                            "domain": "api.0xhzq.k8s.gollum.westeurope.azure.gigantic.io",
                            "securePort": 443
                        },
                        "cloudProvider": "azure",
                        "dns": {
                            "ip": "172.31.0.10"
                        },
                        "domain": "cluster.local",
                        "ingressController": {
                            "docker": {
                                "image": "quay.io/giantswarm/nginx-ingress-controller:0.9.0-beta.11"
                            },
                            "domain": "ingress.0xhzq.k8s.gollum.westeurope.azure.gigantic.io",
                            "insecurePort": 30010,
                            "securePort": 30011,
                            "wildcardDomain": "*.0xhzq.k8s.gollum.westeurope.azure.gigantic.io"
                        },
                        "kubelet": {
                            "altNames": "kubernetes,kubernetes.default,kubernetes.default.svc,kubernetes.default.svc.cluster.local",
                            "domain": "worker.0xhzq.k8s.gollum.westeurope.azure.gigantic.io",
                            "labels": "azure-operator.giantswarm.io/version=2.5.0,giantswarm.io/provider=azure",
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
                            "id": "xdkp7"
                        }
                    ],
                    "scaling": {
                        "max": 1,
                        "min": 1
                    },
                    "version": "",
                    "workers": [
                        {
                            "id": "gjcw9"
                        }
                    ]
                },
                "versionBundle": {
                    "version": "2.5.0"
                }
            },
            "status": {
                "cluster": {
                    "conditions": [
                        {
                            "lastTransitionTime": "2019-08-14T13:03:28.213061393Z",
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
                                "azure-operator.giantswarm.io/version": "2.5.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D2s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "0",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.14.0.5",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "0xhzq-master-000000",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "master",
                                "node-role.kubernetes.io/master": "",
                                "node.kubernetes.io/master": "",
                                "role": "master"
                            },
                            "lastTransitionTime": "2019-08-14T12:59:49.025688165Z",
                            "name": "0xhzq-master-000000",
                            "version": "2.5.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.5.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_A2_v2",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "0",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.14.1.4",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "0xhzq-worker-000000",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-08-14T12:59:49.025689665Z",
                            "name": "0xhzq-worker-000000",
                            "version": "2.5.0"
                        }
                    ],
                    "resources": [
                        {
                            "conditions": [
                                {
                                    "lastTransitionTime": "0001-01-01T00:00:00Z",
                                    "status": "InstancesUpgrading",
                                    "type": "Stage"
                                }
                            ],
                            "name": "instancev9"
                        }
                    ],
                    "scaling": {
                        "desiredCapacity": 0
                    },
                    "versions": [
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-08-14T13:03:33.817482151Z",
                            "semver": "2.5.0"
                        }
                    ]
                }
            }
        },
        {
            "apiVersion": "provider.giantswarm.io/v1alpha1",
            "kind": "AzureConfig",
            "metadata": {
                "creationTimestamp": "2018-12-20T15:07:47Z",
                "finalizers": [
                    "operatorkit.giantswarm.io/azure-operator"
                ],
                "generation": 12,
                "name": "4jr2w",
                "namespace": "default",
                "resourceVersion": "134033420",
                "selfLink": "/apis/provider.giantswarm.io/v1alpha1/namespaces/default/azureconfigs/4jr2w",
                "uid": "02b6d75c-0469-11e9-8aae-000d3a2c9706"
            },
            "spec": {
                "azure": {
                    "dnsZones": {
                        "api": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        },
                        "etcd": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        },
                        "ingress": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        }
                    },
                    "masters": [
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_D2s_v3"
                        }
                    ],
                    "virtualNetwork": {
                        "calicoSubnetCIDR": "10.1.128.0/17",
                        "cidr": "10.1.0.0/16",
                        "masterSubnetCIDR": "10.1.0.0/24",
                        "workerSubnetCIDR": "10.1.1.0/24"
                    },
                    "workers": [
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_D4s_v3"
                        },
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_D4s_v3"
                        },
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_D4s_v3"
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
                        "id": "giantswarm-shared"
                    },
                    "docker": {
                        "daemon": {
                            "cidr": "172.17.0.1/16"
                        }
                    },
                    "etcd": {
                        "altNames": "",
                        "domain": "etcd.4jr2w.k8s.gollum.westeurope.azure.gigantic.io",
                        "port": 2379,
                        "prefix": "giantswarm.io"
                    },
                    "id": "4jr2w",
                    "kubernetes": {
                        "api": {
                            "clusterIPRange": "172.31.0.0/16",
                            "domain": "api.4jr2w.k8s.gollum.westeurope.azure.gigantic.io",
                            "securePort": 443
                        },
                        "cloudProvider": "azure",
                        "dns": {
                            "ip": "172.31.0.10"
                        },
                        "domain": "cluster.local",
                        "ingressController": {
                            "docker": {
                                "image": "quay.io/giantswarm/nginx-ingress-controller:0.9.0-beta.11"
                            },
                            "domain": "ingress.4jr2w.k8s.gollum.westeurope.azure.gigantic.io",
                            "insecurePort": 30010,
                            "securePort": 30011,
                            "wildcardDomain": "*.4jr2w.k8s.gollum.westeurope.azure.gigantic.io"
                        },
                        "kubelet": {
                            "altNames": "kubernetes,kubernetes.default,kubernetes.default.svc,kubernetes.default.svc.cluster.local",
                            "domain": "worker.4jr2w.k8s.gollum.westeurope.azure.gigantic.io",
                            "labels": "azure-operator.giantswarm.io/version=2.4.0,giantswarm.io/provider=azure",
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
                            "id": "qfxx2"
                        }
                    ],
                    "scaling": {
                        "max": 3,
                        "min": 3
                    },
                    "version": "",
                    "workers": [
                        {
                            "id": "bqgm5"
                        },
                        {
                            "id": "b5bpt"
                        },
                        {
                            "id": "z020a"
                        }
                    ]
                },
                "versionBundle": {
                    "version": "2.4.0"
                }
            },
            "status": {
                "cluster": {
                    "conditions": [
                        {
                            "lastTransitionTime": "2019-08-01T11:55:43.293843758Z",
                            "status": "True",
                            "type": "Updated"
                        },
                        {
                            "lastTransitionTime": "2018-12-20T16:01:57.244890524Z",
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
                                "azure-operator.giantswarm.io/version": "2.4.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D2s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "0",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.1.0.5",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "4jr2w-master-000000",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "master",
                                "node-role.kubernetes.io/master": "",
                                "node.kubernetes.io/master": "",
                                "role": "master"
                            },
                            "lastTransitionTime": "2019-08-01T11:52:23.419123865Z",
                            "name": "4jr2w-master-000000",
                            "version": "2.4.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.4.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D4s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "1",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.1.1.5",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "4jr2w-worker-000001",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-08-01T11:52:23.419125265Z",
                            "name": "4jr2w-worker-000001",
                            "version": "2.4.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.4.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D4s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "2",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.1.1.6",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "4jr2w-worker-000002",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-08-01T11:52:23.419125965Z",
                            "name": "4jr2w-worker-000002",
                            "version": "2.4.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.4.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D4s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "4",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.1.1.8",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "4jr2w-worker-000004",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-08-01T11:52:23.419126765Z",
                            "name": "4jr2w-worker-000004",
                            "version": "2.4.0"
                        }
                    ],
                    "resources": [
                        {
                            "conditions": [],
                            "name": ""
                        },
                        {
                            "conditions": [
                                {
                                    "lastTransitionTime": "0001-01-01T00:00:00Z",
                                    "status": "DeploymentInitialized",
                                    "type": "Stage"
                                }
                            ],
                            "name": "instancev8"
                        }
                    ],
                    "scaling": {
                        "desiredCapacity": 0
                    },
                    "versions": [
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-02-25T14:56:18.881455541Z",
                            "semver": "2.0.2"
                        },
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-03-12T10:47:54.325454199Z",
                            "semver": "2.1.0"
                        },
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-03-20T14:28:30.179280551Z",
                            "semver": "2.2.0"
                        },
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-05-14T19:53:55.736509419Z",
                            "semver": "2.3.0"
                        },
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-08-01T11:55:43.29389796Z",
                            "semver": "2.4.0"
                        }
                    ]
                }
            }
        },
        {
            "apiVersion": "provider.giantswarm.io/v1alpha1",
            "kind": "AzureConfig",
            "metadata": {
                "creationTimestamp": "2018-07-19T14:09:16Z",
                "finalizers": [
                    "operatorkit.giantswarm.io/azure-operator"
                ],
                "generation": 28,
                "name": "6iec4",
                "namespace": "default",
                "resourceVersion": "134032347",
                "selfLink": "/apis/provider.giantswarm.io/v1alpha1/namespaces/default/azureconfigs/6iec4",
                "uid": "521af926-8b5d-11e8-91e5-000d3a2c9706"
            },
            "spec": {
                "azure": {
                    "dnsZones": {
                        "api": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        },
                        "etcd": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        },
                        "ingress": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        }
                    },
                    "masters": [
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_D2s_v3"
                        }
                    ],
                    "virtualNetwork": {
                        "calicoSubnetCIDR": "10.15.128.0/17",
                        "cidr": "10.15.0.0/16",
                        "masterSubnetCIDR": "10.15.0.0/24",
                        "workerSubnetCIDR": "10.15.1.0/24"
                    },
                    "workers": [
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_A2_v2"
                        },
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_A2_v2"
                        },
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_A2_v2"
                        },
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_A2_v2"
                        },
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_A2_v2"
                        },
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_A2_v2"
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
                    "etcd": {
                        "altNames": "",
                        "domain": "etcd.6iec4.k8s.gollum.westeurope.azure.gigantic.io",
                        "port": 2379,
                        "prefix": "giantswarm.io"
                    },
                    "id": "6iec4",
                    "kubernetes": {
                        "api": {
                            "clusterIPRange": "172.31.0.0/16",
                            "domain": "api.6iec4.k8s.gollum.westeurope.azure.gigantic.io",
                            "securePort": 443
                        },
                        "cloudProvider": "azure",
                        "dns": {
                            "ip": "172.31.0.10"
                        },
                        "domain": "cluster.local",
                        "ingressController": {
                            "docker": {
                                "image": "quay.io/giantswarm/nginx-ingress-controller:0.9.0-beta.11"
                            },
                            "domain": "ingress.6iec4.k8s.gollum.westeurope.azure.gigantic.io",
                            "insecurePort": 30010,
                            "securePort": 30011,
                            "wildcardDomain": "*.6iec4.k8s.gollum.westeurope.azure.gigantic.io"
                        },
                        "kubelet": {
                            "altNames": "kubernetes,kubernetes.default,kubernetes.default.svc,kubernetes.default.svc.cluster.local",
                            "domain": "worker.6iec4.k8s.gollum.westeurope.azure.gigantic.io",
                            "labels": "azure-operator.giantswarm.io/version=2.4.0,giantswarm.io/provider=azure",
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
                            "id": "gh6mb"
                        }
                    ],
                    "scaling": {
                        "max": 6,
                        "min": 6
                    },
                    "version": "",
                    "workers": [
                        {
                            "id": "o7tve"
                        },
                        {
                            "id": "j6q42"
                        },
                        {
                            "id": "78j05"
                        },
                        {
                            "id": "zwn6n"
                        },
                        {
                            "id": "51b99"
                        },
                        {
                            "id": "z7mjv"
                        }
                    ]
                },
                "versionBundle": {
                    "version": "2.4.0"
                }
            },
            "status": {
                "cluster": {
                    "conditions": [
                        {
                            "lastTransitionTime": "2019-07-15T10:50:15.521705572Z",
                            "status": "True",
                            "type": "Updated"
                        },
                        {
                            "lastTransitionTime": "2018-11-15T13:37:39.863296744Z",
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
                            "lastTransitionTime": "2019-07-15T10:46:56.485518113Z",
                            "name": "6iec4-master-000000",
                            "version": "2.4.0"
                        },
                        {
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
                            "lastTransitionTime": "2019-07-15T10:46:56.485519313Z",
                            "name": "6iec4-worker-000003",
                            "version": "2.4.0"
                        },
                        {
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
                            "lastTransitionTime": "2019-07-15T10:46:56.485519913Z",
                            "name": "6iec4-worker-000004",
                            "version": "2.4.0"
                        },
                        {
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
                            "lastTransitionTime": "2019-07-15T10:46:56.485520513Z",
                            "name": "6iec4-worker-000009",
                            "version": "2.4.0"
                        },
                        {
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
                            "lastTransitionTime": "2019-07-15T10:46:56.485520913Z",
                            "name": "6iec4-worker-00000j",
                            "version": "2.4.0"
                        },
                        {
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
                            "lastTransitionTime": "2019-07-15T10:46:56.485521513Z",
                            "name": "6iec4-worker-00000k",
                            "version": "2.4.0"
                        },
                        {
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
                            "lastTransitionTime": "2019-07-15T10:46:56.485521813Z",
                            "name": "6iec4-worker-00000m",
                            "version": "2.4.0"
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
                            "name": "instancev8"
                        }
                    ],
                    "scaling": {
                        "desiredCapacity": 0
                    },
                    "versions": [
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2018-12-06T11:22:15.474857203Z",
                            "semver": "2.0.2"
                        },
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-02-26T11:10:55.044159067Z",
                            "semver": "2.1.0"
                        },
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-03-12T12:15:40.25307362Z",
                            "semver": "2.0.1"
                        },
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-03-14T09:37:37.996706165Z",
                            "semver": "2.2.0"
                        },
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-07-15T10:50:15.521772475Z",
                            "semver": "2.4.0"
                        }
                    ]
                }
            }
        },
        {
            "apiVersion": "provider.giantswarm.io/v1alpha1",
            "kind": "AzureConfig",
            "metadata": {
                "creationTimestamp": "2018-11-16T10:30:05Z",
                "finalizers": [
                    "operatorkit.giantswarm.io/azure-operator"
                ],
                "generation": 9,
                "name": "g5igw",
                "namespace": "default",
                "resourceVersion": "141248576",
                "selfLink": "/apis/provider.giantswarm.io/v1alpha1/namespaces/default/azureconfigs/g5igw",
                "uid": "9588a2b4-e98a-11e8-8104-000d3a2c9706"
            },
            "spec": {
                "azure": {
                    "dnsZones": {
                        "api": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        },
                        "etcd": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        },
                        "ingress": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        }
                    },
                    "masters": [
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_D2s_v3"
                        }
                    ],
                    "virtualNetwork": {
                        "calicoSubnetCIDR": "10.21.128.0/17",
                        "cidr": "10.21.0.0/16",
                        "masterSubnetCIDR": "10.21.0.0/24",
                        "workerSubnetCIDR": "10.21.1.0/24"
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
                        "id": "giantswarm-shared"
                    },
                    "docker": {
                        "daemon": {
                            "cidr": "172.17.0.1/16"
                        }
                    },
                    "etcd": {
                        "altNames": "",
                        "domain": "etcd.g5igw.k8s.gollum.westeurope.azure.gigantic.io",
                        "port": 2379,
                        "prefix": "giantswarm.io"
                    },
                    "id": "g5igw",
                    "kubernetes": {
                        "api": {
                            "clusterIPRange": "172.31.0.0/16",
                            "domain": "api.g5igw.k8s.gollum.westeurope.azure.gigantic.io",
                            "securePort": 443
                        },
                        "cloudProvider": "azure",
                        "dns": {
                            "ip": "172.31.0.10"
                        },
                        "domain": "cluster.local",
                        "ingressController": {
                            "docker": {
                                "image": "quay.io/giantswarm/nginx-ingress-controller:0.9.0-beta.11"
                            },
                            "domain": "ingress.g5igw.k8s.gollum.westeurope.azure.gigantic.io",
                            "insecurePort": 30010,
                            "securePort": 30011,
                            "wildcardDomain": "*.g5igw.k8s.gollum.westeurope.azure.gigantic.io"
                        },
                        "kubelet": {
                            "altNames": "kubernetes,kubernetes.default,kubernetes.default.svc,kubernetes.default.svc.cluster.local",
                            "domain": "worker.g5igw.k8s.gollum.westeurope.azure.gigantic.io",
                            "labels": "azure-operator.giantswarm.io/version=2.4.0,giantswarm.io/provider=azure",
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
                            "id": "zmg33"
                        }
                    ],
                    "scaling": {
                        "max": 3,
                        "min": 3
                    },
                    "version": "",
                    "workers": [
                        {
                            "id": "86uqh"
                        },
                        {
                            "id": "h4zuf"
                        },
                        {
                            "id": "zn920"
                        }
                    ]
                },
                "versionBundle": {
                    "version": "2.4.0"
                }
            },
            "status": {
                "cluster": {
                    "conditions": [
                        {
                            "lastTransitionTime": "2019-08-01T14:10:23.761473292Z",
                            "status": "True",
                            "type": "Updated"
                        },
                        {
                            "lastTransitionTime": "2018-11-16T11:04:23.938921322Z",
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
                                "azure-operator.giantswarm.io/version": "2.4.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D2s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "0",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.21.0.5",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "g5igw-master-000000",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "master",
                                "node-role.kubernetes.io/master": "",
                                "node.kubernetes.io/master": "",
                                "role": "master"
                            },
                            "lastTransitionTime": "2019-08-01T14:10:14.702580179Z",
                            "name": "g5igw-master-000000",
                            "version": "2.4.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.4.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D2s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "0",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.21.1.4",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "g5igw-worker-000000",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-08-01T14:10:14.702581079Z",
                            "name": "g5igw-worker-000000",
                            "version": "2.4.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.4.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D2s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "1",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.21.1.5",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "g5igw-worker-000001",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-08-01T14:10:14.702581579Z",
                            "name": "g5igw-worker-000001",
                            "version": "2.4.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.4.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D2s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "4",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.21.1.8",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "g5igw-worker-000004",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-08-01T14:10:14.702582279Z",
                            "name": "g5igw-worker-000004",
                            "version": "2.4.0"
                        }
                    ],
                    "resources": [
                        {
                            "conditions": null,
                            "name": ""
                        },
                        {
                            "conditions": [
                                {
                                    "lastTransitionTime": "0001-01-01T00:00:00Z",
                                    "status": "InstancesUpgrading",
                                    "type": "Stage"
                                }
                            ],
                            "name": "instancev8"
                        }
                    ],
                    "scaling": {
                        "desiredCapacity": 0
                    },
                    "versions": [
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-03-15T16:29:12.903976449Z",
                            "semver": "2.3.0"
                        },
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-08-01T12:38:21.229697215Z",
                            "semver": "2.2.0"
                        },
                        {
                            "date": "2018-11-16T11:10:09.631838622Z",
                            "lastTransitionTime": "0001-01-01T00:00:00Z",
                            "semver": "2.0.0"
                        },
                        {
                            "date": "2018-12-07T19:19:42.768964688Z",
                            "lastTransitionTime": "0001-01-01T00:00:00Z",
                            "semver": "2.0.1"
                        },
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-08-01T14:10:23.761526494Z",
                            "semver": "2.4.0"
                        }
                    ]
                }
            }
        },
        {
            "apiVersion": "provider.giantswarm.io/v1alpha1",
            "kind": "AzureConfig",
            "metadata": {
                "creationTimestamp": "2018-11-26T11:56:34Z",
                "finalizers": [
                    "operatorkit.giantswarm.io/azure-operator"
                ],
                "generation": 9,
                "name": "gv83s",
                "namespace": "default",
                "resourceVersion": "131361989",
                "selfLink": "/apis/provider.giantswarm.io/v1alpha1/namespaces/default/azureconfigs/gv83s",
                "uid": "527da9df-f172-11e8-b03b-000d3a2c9706"
            },
            "spec": {
                "azure": {
                    "dnsZones": {
                        "api": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        },
                        "etcd": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        },
                        "ingress": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        }
                    },
                    "masters": [
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_D2s_v3"
                        }
                    ],
                    "virtualNetwork": {
                        "calicoSubnetCIDR": "10.20.128.0/17",
                        "cidr": "10.20.0.0/16",
                        "masterSubnetCIDR": "10.20.0.0/24",
                        "workerSubnetCIDR": "10.20.1.0/24"
                    },
                    "workers": [
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_A8_v2"
                        },
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_A8_v2"
                        },
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_A8_v2"
                        },
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_A8_v2"
                        },
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_A8_v2"
                        }
                    ]
                },
                "cluster": {
                    "calico": {
                        "cidr": 0,
                        "mtu": 1430,
                        "subnet": ""
                    },
                    "docker": {
                        "daemon": {
                            "cidr": "172.17.0.1/16"
                        }
                    },
                    "etcd": {
                        "altNames": "",
                        "domain": "etcd.gv83s.k8s.gollum.westeurope.azure.gigantic.io",
                        "port": 2379,
                        "prefix": "giantswarm.io"
                    },
                    "id": "gv83s",
                    "kubernetes": {
                        "api": {
                            "clusterIPRange": "172.31.0.0/16",
                            "domain": "api.gv83s.k8s.gollum.westeurope.azure.gigantic.io",
                            "securePort": 443
                        },
                        "cloudProvider": "azure",
                        "dns": {
                            "ip": "172.31.0.10"
                        },
                        "domain": "cluster.local",
                        "ingressController": {
                            "docker": {
                                "image": "quay.io/giantswarm/nginx-ingress-controller:0.9.0-beta.11"
                            },
                            "domain": "ingress.gv83s.k8s.gollum.westeurope.azure.gigantic.io",
                            "insecurePort": 30010,
                            "securePort": 30011,
                            "wildcardDomain": "*.gv83s.k8s.gollum.westeurope.azure.gigantic.io"
                        },
                        "kubelet": {
                            "altNames": "kubernetes,kubernetes.default,kubernetes.default.svc,kubernetes.default.svc.cluster.local",
                            "domain": "worker.gv83s.k8s.gollum.westeurope.azure.gigantic.io",
                            "labels": "azure-operator.giantswarm.io/version=2.4.0,giantswarm.io/provider=azure",
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
                            "id": "bhmp3"
                        }
                    ],
                    "scaling": {
                        "max": 5,
                        "min": 5
                    },
                    "version": "",
                    "workers": [
                        {
                            "id": "766nk"
                        },
                        {
                            "id": "if4fi"
                        },
                        {
                            "id": "nry7r"
                        },
                        {
                            "id": "f4abk"
                        },
                        {
                            "id": "j44cy"
                        }
                    ]
                },
                "versionBundle": {
                    "version": "2.4.0"
                }
            },
            "status": {
                "cluster": {
                    "conditions": [
                        {
                            "lastTransitionTime": "2019-07-23T12:22:03.977875834Z",
                            "status": "True",
                            "type": "Updated"
                        },
                        {
                            "lastTransitionTime": "2018-11-26T13:01:50.721409609Z",
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
                                "azure-operator.giantswarm.io/version": "2.4.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D2s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "0",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.20.0.5",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "gv83s-master-000000",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "master",
                                "node-role.kubernetes.io/master": "",
                                "node.kubernetes.io/master": "",
                                "role": "master"
                            },
                            "lastTransitionTime": "2019-07-23T12:18:31.191415928Z",
                            "name": "gv83s-master-000000",
                            "version": "2.4.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.4.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_A8_v2",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "0",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.20.1.4",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "gv83s-worker-000000",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-07-23T12:18:31.191417328Z",
                            "name": "gv83s-worker-000000",
                            "version": "2.4.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.4.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_A8_v2",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "1",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.20.1.5",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "gv83s-worker-000001",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-07-23T12:18:31.191418128Z",
                            "name": "gv83s-worker-000001",
                            "version": "2.4.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.4.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_A8_v2",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "3",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.20.1.7",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "gv83s-worker-000007",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-07-23T12:18:31.191418828Z",
                            "name": "gv83s-worker-000007",
                            "version": "2.4.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.4.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_A8_v2",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "2",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.20.1.6",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "gv83s-worker-00000b",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-07-23T12:18:31.191419328Z",
                            "name": "gv83s-worker-00000b",
                            "version": "2.4.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.4.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_A8_v2",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "0",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.20.1.9",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "gv83s-worker-00000d",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-07-23T12:18:31.191420928Z",
                            "name": "gv83s-worker-00000d",
                            "version": "2.4.0"
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
                            "name": ""
                        },
                        {
                            "conditions": [
                                {
                                    "lastTransitionTime": "0001-01-01T00:00:00Z",
                                    "status": "DeploymentInitialized",
                                    "type": "Stage"
                                }
                            ],
                            "name": "instancev8"
                        }
                    ],
                    "scaling": {
                        "desiredCapacity": 0
                    },
                    "versions": [
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-03-12T19:19:05.540435958Z",
                            "semver": "2.2.0"
                        },
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-05-03T10:48:12.213123007Z",
                            "semver": "2.3.0"
                        },
                        {
                            "date": "2018-11-26T13:05:40.226575708Z",
                            "lastTransitionTime": "0001-01-01T00:00:00Z",
                            "semver": "2.0.0"
                        },
                        {
                            "date": "2018-12-10T18:42:43.013006029Z",
                            "lastTransitionTime": "0001-01-01T00:00:00Z",
                            "semver": "2.0.1"
                        },
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-07-23T12:22:03.977936136Z",
                            "semver": "2.4.0"
                        }
                    ]
                }
            }
        },
        {
            "apiVersion": "provider.giantswarm.io/v1alpha1",
            "kind": "AzureConfig",
            "metadata": {
                "creationTimestamp": "2019-09-03T15:29:23Z",
                "finalizers": [
                    "operatorkit.giantswarm.io/azure-operator"
                ],
                "generation": 1,
                "name": "t9nz5",
                "namespace": "default",
                "resourceVersion": "141247370",
                "selfLink": "/apis/provider.giantswarm.io/v1alpha1/namespaces/default/azureconfigs/t9nz5",
                "uid": "9b3ff7b9-ce5f-11e9-9965-000d3a39f600"
            },
            "spec": {
                "azure": {
                    "dnsZones": {
                        "api": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        },
                        "etcd": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        },
                        "ingress": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        }
                    },
                    "masters": [
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_D2s_v3"
                        }
                    ],
                    "virtualNetwork": {
                        "calicoSubnetCIDR": "10.16.128.0/17",
                        "cidr": "10.16.0.0/16",
                        "masterSubnetCIDR": "10.16.0.0/24",
                        "workerSubnetCIDR": "10.16.1.0/24"
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
                    "docker": {
                        "daemon": {
                            "cidr": "172.17.0.1/16"
                        }
                    },
                    "etcd": {
                        "altNames": "",
                        "domain": "etcd.t9nz5.k8s.gollum.westeurope.azure.gigantic.io",
                        "port": 2379,
                        "prefix": "giantswarm.io"
                    },
                    "id": "t9nz5",
                    "kubernetes": {
                        "api": {
                            "clusterIPRange": "172.31.0.0/16",
                            "domain": "api.t9nz5.k8s.gollum.westeurope.azure.gigantic.io",
                            "securePort": 443
                        },
                        "cloudProvider": "azure",
                        "dns": {
                            "ip": "172.31.0.10"
                        },
                        "domain": "cluster.local",
                        "ingressController": {
                            "docker": {
                                "image": "quay.io/giantswarm/nginx-ingress-controller:0.9.0-beta.11"
                            },
                            "domain": "ingress.t9nz5.k8s.gollum.westeurope.azure.gigantic.io",
                            "insecurePort": 30010,
                            "securePort": 30011,
                            "wildcardDomain": "*.t9nz5.k8s.gollum.westeurope.azure.gigantic.io"
                        },
                        "kubelet": {
                            "altNames": "kubernetes,kubernetes.default,kubernetes.default.svc,kubernetes.default.svc.cluster.local",
                            "domain": "worker.t9nz5.k8s.gollum.westeurope.azure.gigantic.io",
                            "labels": "azure-operator.giantswarm.io/version=2.5.0,giantswarm.io/provider=azure",
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
                            "id": "7etmr"
                        }
                    ],
                    "scaling": {
                        "max": 3,
                        "min": 3
                    },
                    "version": "",
                    "workers": [
                        {
                            "id": "yt3g9"
                        },
                        {
                            "id": "mcm2z"
                        },
                        {
                            "id": "5faue"
                        }
                    ]
                },
                "versionBundle": {
                    "version": "2.5.0"
                }
            },
            "status": {
                "cluster": {
                    "conditions": [
                        {
                            "lastTransitionTime": "2019-09-03T15:43:09.967172769Z",
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
                                "azure-operator.giantswarm.io/version": "2.5.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D2s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "0",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.16.0.5",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "t9nz5-master-000000",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "master",
                                "node-role.kubernetes.io/master": "",
                                "node.kubernetes.io/master": "",
                                "role": "master"
                            },
                            "lastTransitionTime": "2019-09-03T15:39:44.926197163Z",
                            "name": "t9nz5-master-000000",
                            "version": "2.5.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.5.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D2s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "1",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.16.1.5",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "t9nz5-worker-000001",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-09-03T15:39:44.926198063Z",
                            "name": "t9nz5-worker-000001",
                            "version": "2.5.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.5.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D2s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "3",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.16.1.7",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "t9nz5-worker-000003",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-09-03T15:39:44.926198663Z",
                            "name": "t9nz5-worker-000003",
                            "version": "2.5.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.5.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D2s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "4",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.16.1.8",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "t9nz5-worker-000004",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-09-03T15:39:44.926199163Z",
                            "name": "t9nz5-worker-000004",
                            "version": "2.5.0"
                        }
                    ],
                    "resources": [
                        {
                            "conditions": [
                                {
                                    "lastTransitionTime": "0001-01-01T00:00:00Z",
                                    "status": "InstancesUpgrading",
                                    "type": "Stage"
                                }
                            ],
                            "name": "instancev9"
                        }
                    ],
                    "scaling": {
                        "desiredCapacity": 0
                    },
                    "versions": [
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-09-03T15:46:55.355086379Z",
                            "semver": "2.5.0"
                        }
                    ]
                }
            }
        },
        {
            "apiVersion": "provider.giantswarm.io/v1alpha1",
            "kind": "AzureConfig",
            "metadata": {
                "creationTimestamp": "2019-09-19T08:28:48Z",
                "finalizers": [
                    "operatorkit.giantswarm.io/azure-operator"
                ],
                "generation": 1,
                "name": "zf3n9",
                "namespace": "default",
                "resourceVersion": "141248500",
                "selfLink": "/apis/provider.giantswarm.io/v1alpha1/namespaces/default/azureconfigs/zf3n9",
                "uid": "80a59b45-dab7-11e9-b513-000d3a2c9706"
            },
            "spec": {
                "azure": {
                    "dnsZones": {
                        "api": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        },
                        "etcd": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        },
                        "ingress": {
                            "name": "gollum.westeurope.azure.gigantic.io",
                            "resourceGroup": "gollum"
                        }
                    },
                    "masters": [
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_D2s_v3"
                        }
                    ],
                    "virtualNetwork": {
                        "calicoSubnetCIDR": "10.17.128.0/17",
                        "cidr": "10.17.0.0/16",
                        "masterSubnetCIDR": "10.17.0.0/24",
                        "workerSubnetCIDR": "10.17.1.0/24"
                    },
                    "workers": [
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_D8s_v3"
                        },
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_D8s_v3"
                        },
                        {
                            "dockerVolumeSizeGB": 0,
                            "vmSize": "Standard_D8s_v3"
                        }
                    ]
                },
                "cluster": {
                    "calico": {
                        "cidr": 0,
                        "mtu": 1430,
                        "subnet": ""
                    },
                    "docker": {
                        "daemon": {
                            "cidr": "172.17.0.1/16"
                        }
                    },
                    "etcd": {
                        "altNames": "",
                        "domain": "etcd.zf3n9.k8s.gollum.westeurope.azure.gigantic.io",
                        "port": 2379,
                        "prefix": "giantswarm.io"
                    },
                    "id": "zf3n9",
                    "kubernetes": {
                        "api": {
                            "clusterIPRange": "172.31.0.0/16",
                            "domain": "api.zf3n9.k8s.gollum.westeurope.azure.gigantic.io",
                            "securePort": 443
                        },
                        "cloudProvider": "azure",
                        "dns": {
                            "ip": "172.31.0.10"
                        },
                        "domain": "cluster.local",
                        "ingressController": {
                            "docker": {
                                "image": "quay.io/giantswarm/nginx-ingress-controller:0.9.0-beta.11"
                            },
                            "domain": "ingress.zf3n9.k8s.gollum.westeurope.azure.gigantic.io",
                            "insecurePort": 30010,
                            "securePort": 30011,
                            "wildcardDomain": "*.zf3n9.k8s.gollum.westeurope.azure.gigantic.io"
                        },
                        "kubelet": {
                            "altNames": "kubernetes,kubernetes.default,kubernetes.default.svc,kubernetes.default.svc.cluster.local",
                            "domain": "worker.zf3n9.k8s.gollum.westeurope.azure.gigantic.io",
                            "labels": "azure-operator.giantswarm.io/version=2.6.0,giantswarm.io/provider=azure",
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
                            "id": "ufut0"
                        }
                    ],
                    "scaling": {
                        "max": 3,
                        "min": 3
                    },
                    "version": "",
                    "workers": [
                        {
                            "id": "m3jcv"
                        },
                        {
                            "id": "9m0q3"
                        },
                        {
                            "id": "vvq5p"
                        }
                    ]
                },
                "versionBundle": {
                    "version": "2.6.0"
                }
            },
            "status": {
                "cluster": {
                    "conditions": [
                        {
                            "lastTransitionTime": "2019-09-19T08:41:09.082711765Z",
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
                                "azure-operator.giantswarm.io/version": "2.6.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D2s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "0",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.17.0.5",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "zf3n9-master-000000",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "master",
                                "node-role.kubernetes.io/master": "",
                                "node.kubernetes.io/master": "",
                                "role": "master"
                            },
                            "lastTransitionTime": "2019-09-19T08:37:30.846734912Z",
                            "name": "zf3n9-master-000000",
                            "version": "2.6.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.6.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D8s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "1",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.17.1.5",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "zf3n9-worker-000001",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-09-19T08:37:30.846735812Z",
                            "name": "zf3n9-worker-000001",
                            "version": "2.6.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.6.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D8s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "2",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.17.1.6",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "zf3n9-worker-000002",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-09-19T08:37:30.846736312Z",
                            "name": "zf3n9-worker-000002",
                            "version": "2.6.0"
                        },
                        {
                            "labels": {
                                "azure-operator.giantswarm.io/version": "2.6.0",
                                "beta.kubernetes.io/arch": "amd64",
                                "beta.kubernetes.io/instance-type": "Standard_D8s_v3",
                                "beta.kubernetes.io/os": "linux",
                                "failure-domain.beta.kubernetes.io/region": "westeurope",
                                "failure-domain.beta.kubernetes.io/zone": "3",
                                "giantswarm.io/provider": "azure",
                                "ip": "10.17.1.7",
                                "kubernetes.io/arch": "amd64",
                                "kubernetes.io/hostname": "zf3n9-worker-000003",
                                "kubernetes.io/os": "linux",
                                "kubernetes.io/role": "worker",
                                "node-role.kubernetes.io/worker": "",
                                "node.kubernetes.io/worker": "",
                                "role": "worker"
                            },
                            "lastTransitionTime": "2019-09-19T08:37:30.846736812Z",
                            "name": "zf3n9-worker-000003",
                            "version": "2.6.0"
                        }
                    ],
                    "resources": null,
                    "scaling": {
                        "desiredCapacity": 0
                    },
                    "versions": [
                        {
                            "date": "0001-01-01T00:00:00Z",
                            "lastTransitionTime": "2019-09-19T08:44:23.75660127Z",
                            "semver": "2.6.0"
                        }
                    ]
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
