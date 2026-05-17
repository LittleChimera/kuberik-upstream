# Rollout with Datadog Health Check

Use a Datadog query monitor as a health signal for a Kuberik Rollout. While the monitor is OK, the bake period proceeds; if Datadog flips the monitor to Alert during bake, the rollout is paused.

## Prerequisites

- The datadog-controller is installed (it ships in the install bundle, or apply only the integration controller from its release).
- Datadog Operator and the `DatadogMonitor` CRD are installed in the cluster.
- The Datadog API key is configured for the Datadog Operator.

## Apply

```bash
kubectl apply -f monitor.yaml
```

The datadog-controller sees the `kuberik.com/health-check: "true"` annotation and creates a `HealthCheck` resource named `my-app-error-rate` mirroring the monitor's state.

## Wire it to your Rollout

The HealthCheck must reference the Rollout. The datadog-controller does this automatically based on label selection - any Rollout labeled to match the HealthCheck's `selector` will pick it up. Adjust labels on your Rollout accordingly.

## Inspect

```bash
kuberik get healthchecks -n production
kubectl describe healthcheck my-app-error-rate -n production
```

See [docs/healthchecks.md](../../docs/healthchecks.md) for how `HealthCheck` interacts with bake periods, including the `lastErrorTime` witness semantics.
