# Multi-Environment Promotion

Use `Environment` resources to model promotion order across environments. Production waits for staging to deploy successfully before it considers a version.

## Prerequisites

- environment-controller is installed (included in the install bundle).
- You have a Rollout named `my-app` in both the `staging` and `production` namespaces.

## Apply

```bash
kubectl apply -f environments.yaml
```

The environment-controller:

1. Sees that the `production` Environment has `after: [staging]`.
2. Creates a `RolloutGate` in the `production` namespace that stays `passing: false` until the `staging` Environment has deployed the same version.
3. Reports deployment events to the GitHub Deployments API for both environments.

## Inspect

```bash
kuberik get gates -A
kubectl describe environment production -n production
```

## Adding more environments

Add an `after:` entry per upstream environment. Multiple upstreams form an AND - all must deploy before the downstream proceeds.

```yaml
spec:
  after:
    - { name: staging, namespace: staging }
    - { name: canary, namespace: canary }
```

## Parallel environments

For environments that should deploy together (e.g. multi-region production), use `parallel:`:

```yaml
spec:
  parallel:
    - { name: production-eu, namespace: production-eu }
```

The environment-controller treats parallel siblings as a single deployment phase: downstream environments wait for all parallels to deploy.
