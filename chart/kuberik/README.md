# Kuberik Helm Chart

A Helm chart that installs the Kuberik rollout-controller and optionally the integration controllers. Alternative to the [kustomize bundle](../../config/install/kustomization.yaml).

The chart installs:

- 7 CRDs (in `crds/` so Helm applies them before any template): Kuberik core + openkruise + environment.
- rollout-controller: ServiceAccount, leader-election Role/RoleBinding, 10 ClusterRoles and 2 ClusterRoleBindings, metrics Service, Deployment.
- Optional integration controllers (toggle in `values.yaml`):
  - **datadog-controller** - templated
  - **openkruise-controller** - templated
  - **environment-controller** - templated
  - **prometheus-controller** - no upstream release yet; NOTES tells you to build from source
- Dashboard toggle prints the kubectl apply line in NOTES until templates are added.

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
| `dashboard.ingress.enabled` | `false` | Expose the dashboard via an Ingress |
| `dashboard.ingress.host` | `""` | Required when ingress is enabled |
| `metrics.serviceMonitor.enabled` | `false` | Prometheus Operator ServiceMonitor for the rollout-controller |
| `networkPolicy.enabled` | `false` | Restrict ingress/egress for the controller pods |
| `rolloutController.podDisruptionBudget.enabled` | `false` | Emit a PDB (requires replicas > 1) |
| `createNamespace` | `true` | Have the chart create the Namespace. Set false when using `helm install --create-namespace` |

See [values.yaml](values.yaml) for the full schema.

## Common scenarios

### Production (HA, hardened, metrics, ingress)

```bash
helm install kuberik ./chart/kuberik \
  --namespace kuberik-system --create-namespace \
  --set createNamespace=false \
  --set rolloutController.replicas=3 \
  --set rolloutController.podDisruptionBudget.enabled=true \
  --set metrics.serviceMonitor.enabled=true \
  --set networkPolicy.enabled=true \
  --set dashboard.enabled=true \
  --set dashboard.ingress.enabled=true \
  --set dashboard.ingress.host=dashboard.kuberik.example.com \
  --set dashboard.ingress.className=nginx
```

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
  clusterrolloutschedules.kuberik.com \
  rollouttests.rollout.kuberik.com \
  environments.environments.kuberik.com
```

This also deletes every `Rollout`, `RolloutGate`, `HealthCheck`, `RolloutTest`, and `Environment` in the cluster.

## Enabling everything

```bash
helm install kuberik ./chart/kuberik \
  --namespace kuberik-system --create-namespace \
  --set integrations.datadog.enabled=true \
  --set integrations.openkruise.enabled=true \
  --set integrations.environment.enabled=true
```
