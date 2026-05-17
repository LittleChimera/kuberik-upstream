# Rollout with Manual Approval Gate

Builds on [examples/basic-rollout](../basic-rollout) by adding a `RolloutGate` that requires manual approval before a new release is promoted.

## Apply

```bash
# Apply the base rollout first
kubectl apply -f ../basic-rollout/

# Then add the gate
kubectl apply -f gate.yaml
```

## Approve a release

```bash
kuberik approve my-app-approval
```

The Rollout will not promote new versions while `spec.passing` is `false`. Once approved, the controller patches the target OCIRepository and starts the bake period.
