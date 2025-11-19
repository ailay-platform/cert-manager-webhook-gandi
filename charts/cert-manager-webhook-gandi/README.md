# cert-manager-webhook-gandi

A Helm chart for cert-manager-webhook-gandi.

## Usage

Install the helm chart repository:

```bash
helm repo add cert-manager-webhook-gandi https://sintef.github.io/cert-manager-webhook-gandi
```

Install the chart:

```bash
helm install cert-manager-webhook-gandi cert-manager-webhook-gandi/cert-manager-webhook-gandi -f gandiPat=XXX_EXAMPLE_XXX
```

*or using the deprecated `gandiApiToken` secret*

```bash
helm install cert-manager-webhook-gandi cert-manager-webhook-gandi/cert-manager-webhook-gandi -f gandiApiToken=XXX_EXAMPLE_XXX
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| certManager.namespace | string | `"cert-manager"` | Namespace of cert-manager |
| certManager.serviceAccountName | string | `"cert-manager"` | Name of cert-manager's service account |
| containerport | int | `8443` | Container port (in case you have restrictions on the listening port) |
| features.apiPriorityAndFairness | bool | `true` | It is enabled by default since a while. |
| fullnameOverride | string | `""` | Set to override the fullname |
| gandiPat | string | `""` | The secret is not created if not set. |
| ~~gandiApiToken~~ | string | `""` | **Deprecated**. The secret is not created if not set. |
| groupName | string | `"acme.bwolf.me"` | "Group is the API group name this server hosts", if you find this description helful. |
| image.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| image.repository | string | `"ghcr.io/sintef/cert-manager-webhook-gandi"` | Image name |
| image.tag | string | `""` | Image tag (default to Chart's appVersion) |
| image.pullSecrets | list | `[]` | Image pull secrets |
| replicaCount | int | `1` | Number of replicas |
| logLevel | int | `2` | Verbosity of the logs. Set to 6 for verbose logs. |
| nameOverride | string | `""` | Set to override the name |
| nodeSelector | object | `{}` |  |
| resources | object | `{}` |  |
| service.port | int | `443` | Service port |
| service.type | string | `"ClusterIP"` | Service type, e.g. ClusterIP, NodePort, LoadBalancer |
| tolerations | list | `[]` |  |

----------------------------------------------
