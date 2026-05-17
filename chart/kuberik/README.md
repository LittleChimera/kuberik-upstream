# Kuberik Helm Chart

A Helm chart that installs the Kuberik rollout-controller and optionally the integration controllers. Alternative to the [kustomize bundle](../../config/install/kustomization.yaml).

The chart installs:

- The 5 Kuberik CRDs (from `crds/` so Helm applies them before any template).
- The rollout-controller: ServiceAccount, leader-election Role/RoleBinding, 10 ClusterRoles and 2 ClusterRoleBindings, metrics Service, and Deployment.

Integration controllers (datadog, prometheus, openkruise, environment) and the dashboard have toggles in `values.yaml`. Their workload templates are planned - until then, enabling a toggle prints the matching `kubectl apply` line in the post-install NOTES.

## Install

```bash
helm repo add kuberik https://kuberik.github.io/charts
helm install kuberik kuberik/kuberik \
  --namespace kuberik-system \
  --create-namespace
```

Or from this directory directly:

```bash
helm install kuberik ./chart/kuberik \
  --namespace kuberik-system \
  --create-namespace
```

## Values

| Key | Default | Description |
|---|---|---|
| `namespace` | `kuberik-system` | Namespace the controllers run in |
| `rolloutController.enabled` | `true` | Install the core controller |
| `rolloutController.image.repository` | `ghcr.io/kuberik/rollout-controller` | Controller image |
| `rolloutController.image.tag` | _appVersion_ | Image tag override |
| `rolloutController.replicas` | `1` | Number of controller replicas |
| `rolloutController.logLevel` | `info` | Controller log level |
| `integrations.datadog.enabled` | `false` | Install datadog-controller |
| `integrations.prometheus.enabled` | `false` | Install prometheus-controller |
| `integrations.openkruise.enabled` | `false` | Install openkruise-controller |
| `integrations.environment.enabled` | `false` | Install environment-controller |
| `dashboard.enabled` | `false` | Install rollout-dashboard |

See [values.yaml](values.yaml) for the full schema.

## Uninstall

```bash
helm uninstall kuberik -n kuberik-system
```

CRDs are not removed by `helm uninstall`. To delete them:

```bash
kubectl delete crd \
  rollouts.kuberik.com \
  rolloutgates.kuberik.com \
  healthchecks.kuberik.com \
  rolloutschedules.kuberik.com \
  clusterrolloutschedules.kuberik.com
```

This also deletes every `Rollout`, `RolloutGate`, and `HealthCheck` in the cluster.
