# Rollout with Prometheus Health Check

Gate the bake period on a PromQL query. The prometheus-controller evaluates the query on a schedule and updates the corresponding `HealthCheck` resource.

## Prerequisites

- prometheus-controller is installed (included in the install bundle).
- Prometheus is reachable from the prometheus-controller pod.

## Apply

```bash
kubectl apply -f healthcheck.yaml
```

This example fails the bake if p99 HTTP latency for `my-app` exceeds 500ms over a 5-minute window.

## Tuning

- `interval`: how often the controller evaluates the query.
- `threshold` + `comparison`: the condition that defines "healthy". Use `GreaterThan` to flag when a value exceeds a budget, or `LessThan` for SLO-style upper bounds.
- For more complex predicates (multiple metrics, alert-state mirroring), see the [prometheus-controller README](https://github.com/kuberik/prometheus-controller).

## Inspect

```bash
kuberik get healthchecks -n production
kubectl describe healthcheck my-app-latency -n production
```
