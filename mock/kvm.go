package mock

// KVMHealthy is a JSON string representing a healthy KVM cluster
var KVMHealthy = `
{
	"apiVersion": "provider.giantswarm.io/v1alpha1",
	"kind": "KVMConfig",
	"metadata": {
		"creationTimestamp": "2019-09-23T11:53:38Z",
		"name": "cxx2e",
		"namespace": "default"
	},
	"spec": {
		"cluster": {
			"calico": {
				"cidr": 16,
				"mtu": 1430,
				"subnet": "192.168.0.0"
			},
			"customer": {
				"id": "giantswarm"
			},
			"docker": {
				"daemon": {
					"cidr": "172.17.0.1/16"
				}
			},
			"id": "cxx2e",
			"masters": [
				{
					"id": "hhn5a"
				}
			],
			"scaling": {
				"max": 3,
				"min": 3
			},
			"version": "",
			"workers": [
				{
					"id": "hi5ie"
				},
				{
					"id": "pq295"
				},
				{
					"id": "oh8xv"
				}
			]
		},
		"kvm": {
			"masters": [
				{
					"cpus": 2,
					"disk": 40,
					"dockerVolumeSizeGB": 0,
					"memory": "8G"
				}
			],
			"network": {
				"flannel": {
					"vni": 4
				}
			},
			"workers": [
				{
					"cpus": 2,
					"disk": 40,
					"dockerVolumeSizeGB": 0,
					"memory": "3G"
				},
				{
					"cpus": 2,
					"disk": 40,
					"dockerVolumeSizeGB": 0,
					"memory": "3G"
				},
				{
					"cpus": 2,
					"disk": 40,
					"dockerVolumeSizeGB": 0,
					"memory": "3G"
				}
			]
		},
		"versionBundle": {
			"version": "3.8.0"
		}
	},
	"status": {
		"cluster": {
			"conditions": [
				{
					"lastTransitionTime": "2019-09-23T12:08:24.66025979Z",
					"status": "False",
					"type": "Updating"
				},
				{
					"lastTransitionTime": "2019-09-23T12:01:19.857790161Z",
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
						"beta.kubernetes.io/arch": "amd64",
						"beta.kubernetes.io/os": "linux",
						"giantswarm.io/provider": "kvm",
						"ip": "172.23.0.206",
						"kubernetes.io/arch": "amd64",
						"kubernetes.io/hostname": "master-hhn5a-66b4789bbc-tjvb4",
						"kubernetes.io/os": "linux",
						"kubernetes.io/role": "master",
						"kvm-operator.giantswarm.io/version": "3.8.0",
						"node-role.kubernetes.io/master": "",
						"node.kubernetes.io/master": "",
						"role": "master"
					},
					"lastTransitionTime": "2019-09-23T17:55:06.434468696Z",
					"name": "master-hhn5a-66b4789bbc-tjvb4",
					"version": "3.8.0"
				},
				{
					"labels": {
						"beta.kubernetes.io/arch": "amd64",
						"beta.kubernetes.io/os": "linux",
						"giantswarm.io/provider": "kvm",
						"ip": "172.23.0.154",
						"kubernetes.io/arch": "amd64",
						"kubernetes.io/hostname": "worker-hi5ie-5f79675f59-5wgss",
						"kubernetes.io/os": "linux",
						"kubernetes.io/role": "worker",
						"kvm-operator.giantswarm.io/version": "3.8.0",
						"node-role.kubernetes.io/worker": "",
						"node.kubernetes.io/worker": "",
						"role": "worker"
					},
					"lastTransitionTime": "2019-09-23T17:55:06.434469572Z",
					"name": "worker-hi5ie-5f79675f59-5wgss",
					"version": "3.8.0"
				},
				{
					"labels": {
						"beta.kubernetes.io/arch": "amd64",
						"beta.kubernetes.io/os": "linux",
						"giantswarm.io/provider": "kvm",
						"ip": "172.23.0.242",
						"kubernetes.io/arch": "amd64",
						"kubernetes.io/hostname": "worker-oh8xv-c549d84f-24z2s",
						"kubernetes.io/os": "linux",
						"kubernetes.io/role": "worker",
						"kvm-operator.giantswarm.io/version": "3.5.0",
						"node-role.kubernetes.io/worker": "",
						"node.kubernetes.io/worker": "",
						"role": "worker"
					},
					"lastTransitionTime": "2019-09-23T17:55:06.434470124Z",
					"name": "worker-oh8xv-c549d84f-24z2s",
					"version": "3.5.0"
				},
				{
					"labels": {
						"beta.kubernetes.io/arch": "amd64",
						"beta.kubernetes.io/os": "linux",
						"giantswarm.io/provider": "kvm",
						"ip": "172.23.0.86",
						"kubernetes.io/arch": "amd64",
						"kubernetes.io/hostname": "worker-pq295-6c4bbcccf7-kql2k",
						"kubernetes.io/os": "linux",
						"kubernetes.io/role": "worker",
						"kvm-operator.giantswarm.io/version": "3.5.0",
						"node-role.kubernetes.io/worker": "",
						"node.kubernetes.io/worker": "",
						"role": "worker"
					},
					"lastTransitionTime": "2019-09-23T17:55:06.434470788Z",
					"name": "worker-pq295-6c4bbcccf7-kql2k",
					"version": "3.5.0"
				}
			],
			"resources": null,
			"scaling": {
				"desiredCapacity": 0
			},
			"versions": [
				{
					"date": "0001-01-01T00:00:00Z",
					"lastTransitionTime": "2019-09-23T12:01:19.910404547Z",
					"semver": "3.5.0"
				}
			]
		},
		"kvm": {
			"nodeIndexes": {
				"hhn5a": 1,
				"hi5ie": 2,
				"oh8xv": 4,
				"pq295": 3
			}
		}
	}
}`

// KVMHealthyTC is a JSON string representing a healthy KVM tenant cluster's node list.
var KVMHealthyTC = `
{
    "kind": "NodeList",
    "apiVersion": "v1",
    "metadata": {
        "selfLink": "/api/v1/nodes",
        "resourceVersion": "1066"
    },
    "items": [
        {
            "metadata": {
                "name": "master-8w5xy-58494dd955-7x4px",
                "selfLink": "/api/v1/nodes/master-8w5xy-58494dd955-7x4px",
                "uid": "d72ea8fd-e501-11e9-8b24-deadbe8f243e",
                "resourceVersion": "984",
                "creationTimestamp": "2019-10-02T10:46:07Z",
                "labels": {
                    "beta.kubernetes.io/arch": "amd64",
                    "beta.kubernetes.io/os": "linux",
                    "giantswarm.io/provider": "kvm",
                    "ip": "172.23.2.210",
                    "kubernetes.io/arch": "amd64",
                    "kubernetes.io/hostname": "master-8w5xy-58494dd955-7x4px",
                    "kubernetes.io/os": "linux",
                    "kubernetes.io/role": "master",
                    "kvm-operator.giantswarm.io/version": "3.8.0",
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
                "taints": [
                    {
                        "key": "node-role.kubernetes.io/master",
                        "effect": "NoSchedule"
                    }
                ]
            },
            "status": {
                "capacity": {
                    "cpu": "2",
                    "ephemeral-storage": "5110Mi",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "8168392Ki",
                    "pods": "110"
                },
                "allocatable": {
                    "cpu": "1400m",
                    "ephemeral-storage": "4086Mi",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "6259656Ki",
                    "pods": "110"
                },
                "conditions": [
                    {
                        "type": "NetworkUnavailable",
                        "status": "False",
                        "lastHeartbeatTime": "2019-10-02T10:48:41Z",
                        "lastTransitionTime": "2019-10-02T10:48:41Z",
                        "reason": "CalicoIsUp",
                        "message": "Calico is running on this node"
                    },
                    {
                        "type": "MemoryPressure",
                        "status": "False",
                        "lastHeartbeatTime": "2019-10-02T10:49:39Z",
                        "lastTransitionTime": "2019-10-02T10:45:59Z",
                        "reason": "KubeletHasSufficientMemory",
                        "message": "kubelet has sufficient memory available"
                    },
                    {
                        "type": "DiskPressure",
                        "status": "False",
                        "lastHeartbeatTime": "2019-10-02T10:49:39Z",
                        "lastTransitionTime": "2019-10-02T10:45:59Z",
                        "reason": "KubeletHasNoDiskPressure",
                        "message": "kubelet has no disk pressure"
                    },
                    {
                        "type": "PIDPressure",
                        "status": "False",
                        "lastHeartbeatTime": "2019-10-02T10:49:39Z",
                        "lastTransitionTime": "2019-10-02T10:45:59Z",
                        "reason": "KubeletHasSufficientPID",
                        "message": "kubelet has sufficient PID available"
                    },
                    {
                        "type": "Ready",
                        "status": "True",
                        "lastHeartbeatTime": "2019-10-02T10:49:39Z",
                        "lastTransitionTime": "2019-10-02T10:49:39Z",
                        "reason": "KubeletReady",
                        "message": "kubelet is posting ready status"
                    }
                ],
                "addresses": [
                    {
                        "type": "InternalIP",
                        "address": "172.23.2.210"
                    },
                    {
                        "type": "Hostname",
                        "address": "master-8w5xy-58494dd955-7x4px"
                    }
                ],
                "daemonEndpoints": {
                    "kubeletEndpoint": {
                        "Port": 10250
                    }
                },
                "nodeInfo": {
                    "machineID": "210dc0e1e3334cd6a604b97bca05683e",
                    "systemUUID": "311936e2303b034fe7ef70182235b8cb",
                    "bootID": "b583d701-0843-4755-8656-2f98c77090ca",
                    "kernelVersion": "4.19.50-coreos-r1",
                    "osImage": "Container Linux by CoreOS 2135.4.0 (Rhyolite)",
                    "containerRuntimeVersion": "docker://18.6.3",
                    "kubeletVersion": "v1.14.6",
                    "kubeProxyVersion": "v1.14.6",
                    "operatingSystem": "linux",
                    "architecture": "amd64"
                },
                "images": [
                    {
                        "names": [
                            "quay.io/giantswarm/hyperkube@sha256:e1cf1471c9ae772699fc0fc8a4a53970ca0cb5d6bb0f516ccda91036368dd20f",
                            "quay.io/giantswarm/hyperkube:v1.14.6"
                        ],
                        "sizeBytes": 604022734
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/node@sha256:229f5185643de43c7642d977f08d90bc8878a267a009019ddff0a40d93befb86",
                            "quay.io/giantswarm/node:v3.8.2"
                        ],
                        "sizeBytes": 188867043
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cni@sha256:d769038f47cf5a23bd88f7698e4e5816c220d0801fd9ea2496660809e39448c9",
                            "quay.io/giantswarm/cni:v3.8.2"
                        ],
                        "sizeBytes": 157250943
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
                            "quay.io/giantswarm/docker-kubectl@sha256:a56a4cb53f1668a38b448e77d63f17fd554c085d4be3555a012d3aa2591a0694",
                            "quay.io/giantswarm/docker-kubectl:e777d4eaf369d4dabc393c5da42121c2a725ea6a"
                        ],
                        "sizeBytes": 61002517
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/kube-controllers@sha256:08d1fd5c3fec642fc143dedb6524758856351d372e7c6563dbef906c33b36d20",
                            "quay.io/giantswarm/kube-controllers:v3.8.2"
                        ],
                        "sizeBytes": 46809393
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
                "name": "worker-3505b-57c995659f-55shv",
                "selfLink": "/api/v1/nodes/worker-3505b-57c995659f-55shv",
                "uid": "f02e83ad-e501-11e9-8b24-deadbe8f243e",
                "resourceVersion": "1030",
                "creationTimestamp": "2019-10-02T10:46:49Z",
                "labels": {
                    "beta.kubernetes.io/arch": "amd64",
                    "beta.kubernetes.io/os": "linux",
                    "giantswarm.io/provider": "kvm",
                    "ip": "172.23.2.22",
                    "kubernetes.io/arch": "amd64",
                    "kubernetes.io/hostname": "worker-3505b-57c995659f-55shv",
                    "kubernetes.io/os": "linux",
                    "kubernetes.io/role": "worker",
                    "kvm-operator.giantswarm.io/version": "3.8.0",
                    "node-role.kubernetes.io/worker": "",
                    "node.kubernetes.io/worker": "",
                    "role": "worker"
                },
                "annotations": {
                    "node.alpha.kubernetes.io/ttl": "0",
                    "volumes.kubernetes.io/controller-managed-attach-detach": "true"
                }
            },
            "spec": {},
            "status": {
                "capacity": {
                    "cpu": "2",
                    "ephemeral-storage": "10230Mi",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "3007436Ki",
                    "pods": "110"
                },
                "allocatable": {
                    "cpu": "1500m",
                    "ephemeral-storage": "9206Mi",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "1622988Ki",
                    "pods": "110"
                },
                "conditions": [
                    {
                        "type": "NetworkUnavailable",
                        "status": "False",
                        "lastHeartbeatTime": "2019-10-02T10:48:26Z",
                        "lastTransitionTime": "2019-10-02T10:48:26Z",
                        "reason": "CalicoIsUp",
                        "message": "Calico is running on this node"
                    },
                    {
                        "type": "MemoryPressure",
                        "status": "False",
                        "lastHeartbeatTime": "2019-10-02T10:50:00Z",
                        "lastTransitionTime": "2019-10-02T10:46:49Z",
                        "reason": "KubeletHasSufficientMemory",
                        "message": "kubelet has sufficient memory available"
                    },
                    {
                        "type": "DiskPressure",
                        "status": "False",
                        "lastHeartbeatTime": "2019-10-02T10:50:00Z",
                        "lastTransitionTime": "2019-10-02T10:46:49Z",
                        "reason": "KubeletHasNoDiskPressure",
                        "message": "kubelet has no disk pressure"
                    },
                    {
                        "type": "PIDPressure",
                        "status": "False",
                        "lastHeartbeatTime": "2019-10-02T10:50:00Z",
                        "lastTransitionTime": "2019-10-02T10:46:49Z",
                        "reason": "KubeletHasSufficientPID",
                        "message": "kubelet has sufficient PID available"
                    },
                    {
                        "type": "Ready",
                        "status": "True",
                        "lastHeartbeatTime": "2019-10-02T10:50:00Z",
                        "lastTransitionTime": "2019-10-02T10:48:20Z",
                        "reason": "KubeletReady",
                        "message": "kubelet is posting ready status"
                    }
                ],
                "addresses": [
                    {
                        "type": "InternalIP",
                        "address": "172.23.2.22"
                    },
                    {
                        "type": "Hostname",
                        "address": "worker-3505b-57c995659f-55shv"
                    }
                ],
                "daemonEndpoints": {
                    "kubeletEndpoint": {
                        "Port": 10250
                    }
                },
                "nodeInfo": {
                    "machineID": "d678dd62f9054a56ad4fe6f145167070",
                    "systemUUID": "311936e2303b034fe7ef70182235b8cb",
                    "bootID": "ba285422-8f17-4d54-abf5-8287d7a238f4",
                    "kernelVersion": "4.19.50-coreos-r1",
                    "osImage": "Container Linux by CoreOS 2135.4.0 (Rhyolite)",
                    "containerRuntimeVersion": "docker://18.6.3",
                    "kubeletVersion": "v1.14.6",
                    "kubeProxyVersion": "v1.14.6",
                    "operatingSystem": "linux",
                    "architecture": "amd64"
                },
                "images": [
                    {
                        "names": [
                            "quay.io/giantswarm/hyperkube@sha256:e1cf1471c9ae772699fc0fc8a4a53970ca0cb5d6bb0f516ccda91036368dd20f",
                            "quay.io/giantswarm/hyperkube:v1.14.6"
                        ],
                        "sizeBytes": 604022734
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/node@sha256:229f5185643de43c7642d977f08d90bc8878a267a009019ddff0a40d93befb86",
                            "quay.io/giantswarm/node:v3.8.2"
                        ],
                        "sizeBytes": 188867043
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cni@sha256:d769038f47cf5a23bd88f7698e4e5816c220d0801fd9ea2496660809e39448c9",
                            "quay.io/giantswarm/cni:v3.8.2"
                        ],
                        "sizeBytes": 157250943
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
                "name": "worker-fqu7z-6b4d5d9cf5-vkrxt",
                "selfLink": "/api/v1/nodes/worker-fqu7z-6b4d5d9cf5-vkrxt",
                "uid": "f032005e-e501-11e9-8b24-deadbe8f243e",
                "resourceVersion": "948",
                "creationTimestamp": "2019-10-02T10:46:49Z",
                "labels": {
                    "beta.kubernetes.io/arch": "amd64",
                    "beta.kubernetes.io/os": "linux",
                    "giantswarm.io/provider": "kvm",
                    "ip": "172.23.2.138",
                    "kubernetes.io/arch": "amd64",
                    "kubernetes.io/hostname": "worker-fqu7z-6b4d5d9cf5-vkrxt",
                    "kubernetes.io/os": "linux",
                    "kubernetes.io/role": "worker",
                    "kvm-operator.giantswarm.io/version": "3.8.0",
                    "node-role.kubernetes.io/worker": "",
                    "node.kubernetes.io/worker": "",
                    "role": "worker"
                },
                "annotations": {
                    "node.alpha.kubernetes.io/ttl": "0",
                    "volumes.kubernetes.io/controller-managed-attach-detach": "true"
                }
            },
            "spec": {},
            "status": {
                "capacity": {
                    "cpu": "2",
                    "ephemeral-storage": "10230Mi",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "3007420Ki",
                    "pods": "110"
                },
                "allocatable": {
                    "cpu": "1500m",
                    "ephemeral-storage": "9206Mi",
                    "hugepages-1Gi": "0",
                    "hugepages-2Mi": "0",
                    "memory": "1622972Ki",
                    "pods": "110"
                },
                "conditions": [
                    {
                        "type": "NetworkUnavailable",
                        "status": "False",
                        "lastHeartbeatTime": "2019-10-02T10:48:25Z",
                        "lastTransitionTime": "2019-10-02T10:48:25Z",
                        "reason": "CalicoIsUp",
                        "message": "Calico is running on this node"
                    },
                    {
                        "type": "MemoryPressure",
                        "status": "False",
                        "lastHeartbeatTime": "2019-10-02T10:49:30Z",
                        "lastTransitionTime": "2019-10-02T10:46:49Z",
                        "reason": "KubeletHasSufficientMemory",
                        "message": "kubelet has sufficient memory available"
                    },
                    {
                        "type": "DiskPressure",
                        "status": "False",
                        "lastHeartbeatTime": "2019-10-02T10:49:30Z",
                        "lastTransitionTime": "2019-10-02T10:46:49Z",
                        "reason": "KubeletHasNoDiskPressure",
                        "message": "kubelet has no disk pressure"
                    },
                    {
                        "type": "PIDPressure",
                        "status": "False",
                        "lastHeartbeatTime": "2019-10-02T10:49:30Z",
                        "lastTransitionTime": "2019-10-02T10:46:49Z",
                        "reason": "KubeletHasSufficientPID",
                        "message": "kubelet has sufficient PID available"
                    },
                    {
                        "type": "Ready",
                        "status": "True",
                        "lastHeartbeatTime": "2019-10-02T10:49:30Z",
                        "lastTransitionTime": "2019-10-02T10:48:20Z",
                        "reason": "KubeletReady",
                        "message": "kubelet is posting ready status"
                    }
                ],
                "addresses": [
                    {
                        "type": "InternalIP",
                        "address": "172.23.2.138"
                    },
                    {
                        "type": "Hostname",
                        "address": "worker-fqu7z-6b4d5d9cf5-vkrxt"
                    }
                ],
                "daemonEndpoints": {
                    "kubeletEndpoint": {
                        "Port": 10250
                    }
                },
                "nodeInfo": {
                    "machineID": "b2fba48ba6ba4ba288a6014d966e9432",
                    "systemUUID": "311936e2303b034fe7ef70182235b8cb",
                    "bootID": "960e3b60-52d9-4d17-af48-be3ecf0d01f4",
                    "kernelVersion": "4.19.50-coreos-r1",
                    "osImage": "Container Linux by CoreOS 2135.4.0 (Rhyolite)",
                    "containerRuntimeVersion": "docker://18.6.3",
                    "kubeletVersion": "v1.14.6",
                    "kubeProxyVersion": "v1.14.6",
                    "operatingSystem": "linux",
                    "architecture": "amd64"
                },
                "images": [
                    {
                        "names": [
                            "quay.io/giantswarm/hyperkube@sha256:e1cf1471c9ae772699fc0fc8a4a53970ca0cb5d6bb0f516ccda91036368dd20f",
                            "quay.io/giantswarm/hyperkube:v1.14.6"
                        ],
                        "sizeBytes": 604022734
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/node@sha256:229f5185643de43c7642d977f08d90bc8878a267a009019ddff0a40d93befb86",
                            "quay.io/giantswarm/node:v3.8.2"
                        ],
                        "sizeBytes": 188867043
                    },
                    {
                        "names": [
                            "quay.io/giantswarm/cni@sha256:d769038f47cf5a23bd88f7698e4e5816c220d0801fd9ea2496660809e39448c9",
                            "quay.io/giantswarm/cni:v3.8.2"
                        ],
                        "sizeBytes": 157250943
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
