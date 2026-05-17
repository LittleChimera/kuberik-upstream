# Rollouts on a Schedule

Use a `RolloutSchedule` to control _when_ a rollout is allowed to proceed, independent of approval. Two patterns:

## Business-hours-only deployments

`schedule.yaml` allows promotions only between 9am and 5pm Eastern on weekdays. Outside that window the auto-managed gate is closed and the Rollout waits.

```bash
kubectl apply -f schedule.yaml
```

## Holiday change freeze

`freeze.yaml` is a `ClusterRolloutSchedule` that blocks promotions in every namespace labeled `kuberik.com/freezable=true` for a one-week window. Use this for end-of-year freezes, audit windows, on-call coverage gaps.

```bash
kubectl label namespace production kuberik.com/freezable=true
kubectl apply -f freeze.yaml
```

## Combining with manual approval

A Rollout can be gated by **both** a schedule and a manual approval. The Rollout proceeds only when all gates pass, so a release that needs to go out off-hours requires both flipping the manual gate and either waiting for the schedule window or temporarily removing the schedule.
