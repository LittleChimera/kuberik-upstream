# Multi-Environment Promotion

Use `Environment` resources to model promotion order across environments. Production waits for staging to deploy successfully before it considers a version.

## Prerequisites

- environment-controller is installed (`integrations.environment.enabled=true` in the Helm chart, or `kuberik install --all`).
- You have a Rollout named `my-app` in both the `staging` and `production` namespaces.
- A `github-token` Secret in each Environment's namespace with a GitHub PAT that can write Deployments to `myorg/my-app`.

## Apply

```bash
kubectl apply -f environments.yaml
```

The environment-controller:

1. Sees the `production` Environment's `relationship.type=After` pointing at `staging`.
2. Creates a `RolloutGate` in the `production` namespace that stays `passing: false` until the `staging` Environment has reported a successful deployment via its GitHub backend.
3. Reports deployment events for each Environment to the GitHub Deployments API.

## Inspect

```bash
kuberik get gates -A
kubectl describe environment production -n production
```

## Adding a chained environment

Each Environment has at most one `relationship` (single, not a list). For chains, point each downstream at its immediate upstream:

```yaml
# canary follows staging
spec:
  relationship: { type: After, environment: staging }
---
# production follows canary
spec:
  relationship: { type: After, environment: canary }
```

## Parallel environments

For environments that should be treated as a single deployment phase (e.g. multi-region production), use `relationship.type: Parallel` pointing at a sibling:

```yaml
spec:
  relationship: { type: Parallel, environment: production-eu }
```
