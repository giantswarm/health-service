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
