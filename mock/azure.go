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
{}
`
