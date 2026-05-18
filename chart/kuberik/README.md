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

### gh-pages helm repo

```bash
helm repo add kuberik https://kuberik.github.io/kuberik
helm repo update
helm install kuberik kuberik/kuberik \
  --namespace kuberik-system --create-namespace \
  --set createNamespace=false
```

### OCI registry (Helm 3.8+)

```bash
helm install kuberik oci://ghcr.io/kuberik/charts/kuberik \
  --namespace kuberik-system --create-namespace \
  --set createNamespace=false
```

### From this directory directly

```bash
helm install kuberik ./chart/kuberik \
  --namespace kuberik-system
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

## Troubleshooting

### `namespaces "kuberik-system" already exists` during install

You used `helm install --create-namespace` while the chart also tried to create the namespace. Pick one:

```bash
# A) Let helm own the namespace
helm install kuberik ./chart/kuberik \
  --namespace kuberik-system --create-namespace \
  --set createNamespace=false

# B) Let the chart own the namespace
helm install kuberik ./chart/kuberik \
  --namespace kuberik-system
```

### `helm test` pod can't list deployments

The default ServiceAccount in the namespace usually cannot list deployments. The chart's test pod uses `rollout-controller-controller-manager` by default, which has the necessary permissions. If you override `tests.serviceAccountName`, make sure it can `get`/`list`/`watch` Deployments.

### Image pull failures for `rollout-controller`

The chart defaults `rolloutController.image.tag` to `v` + `Chart.AppVersion`. If you are testing a fork with a non-`v`-prefixed tag, set the tag explicitly:

```bash
--set rolloutController.image.tag=0.7.0
```

### ServiceMonitor not picked up by Prometheus

The ServiceMonitor lives in `.Values.namespace` by default. If your Prometheus Operator scrapes a different namespace, set:

```bash
--set metrics.serviceMonitor.namespace=monitoring
```
