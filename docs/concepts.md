# Concepts

## Rollout

A `Rollout` defines a release pipeline for one workload. It references a Flux `ImagePolicy` to discover new releases and specifies how to promote them: which Flux resources to patch, how long to observe health (bake time), and which gates must pass before proceeding. The rollout-controller reconciles the `Rollout` continuously, advancing through versions as conditions are met.

## RolloutGate

A `RolloutGate` is a named boolean condition scoped to a namespace. When `spec.passing` is `false`, any `Rollout` that references this gate will not advance to a new release. Gates can be set manually (human approval), created automatically by integration controllers (e.g., environment-controller), or managed by a `RolloutSchedule`. Multiple gates can block a single rollout, and all must pass before promotion proceeds.

## HealthCheck

A `HealthCheck` records the observed health state of a workload component. It is typically created and updated by an integration controller (datadog-controller, prometheus-controller) rather than by hand. The rollout-controller reads `HealthCheck` status during the bake period to determine whether a deployment is stable. An unhealthy `HealthCheck` during bake time can pause or block the current release.

## Environment

An `Environment` (managed by environment-controller) represents a logical deployment target - for example, `staging` or `production`. It links a `Rollout` to a backend like the GitHub Deployments API, so deployment events are recorded in GitHub's environment tracking UI. Environments can declare relationships (`After`, `Parallel`) to model promotion dependencies across environments.

## RolloutSchedule / ClusterRolloutSchedule

A `RolloutSchedule` creates and manages `RolloutGate` resources on a time-based schedule. The `action` field controls the direction: `Allow` means gates pass only during the schedule window (block outside it), `Deny` means gates are blocked during the window (allow outside it). `ClusterRolloutSchedule` applies across multiple namespaces via a `namespaceSelector`. Common uses: business-hours-only deployments, holiday freezes, maintenance windows.
