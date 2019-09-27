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
{}
`
