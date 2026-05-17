# Roadmap

The Kuberik roadmap is intentionally short. We focus on a few themes per release rather than a long backlog of speculative items. For day-to-day issue tracking see [GitHub Issues](https://github.com/kuberik/kuberik/issues); for design proposals see [RFCs](rfcs/README.md); for open conversation see [GitHub Discussions](https://github.com/kuberik/kuberik/discussions).

## Themes

### Stability and observability

- Harden the rollout-controller reconciliation loop under partial outages of integration controllers
- Improve `HealthCheck` status reporting (history, transition reasons, witness semantics)
- Surface rollout state and gate decisions in the dashboard with a single-pane-of-glass view

### Operator UX

- Expand the `kuberik` CLI with more day-2 commands (`logs`, `events`, `describe`)
- Helm chart for installation (alternative to the kustomize bundle)
- First-class shell completions and structured output (`-o json`, `-o yaml`)

### Ecosystem integrations

- More health-check sources beyond Datadog/Prometheus (CloudWatch, Grafana, custom HTTP)
- More gate sources (manual approvals via Slack, ChatOps, webhook conditions)
- Tighter Flux integration: per-environment auto-promotion based on upstream image changes

### Multi-cluster

- Reference architecture for orchestrating rollouts across multiple clusters
- Optional control plane for organizations running Kuberik at fleet scale

## Out of scope

- Replacing Flux as a GitOps engine - Kuberik composes with Flux; it does not aim to replace `kustomize-controller` or `helm-controller`
- Replacing Argo Rollouts traffic-shifting primitives - Kuberik integrates with strategy controllers (e.g. OpenKruise) rather than implementing traffic management itself
- A built-in CI engine - rollouts are gated by external signals; we do not orchestrate build pipelines

## Contributing to the roadmap

Roadmap themes are revisited at the start of each release cycle. To propose a new theme, open a discussion and tag it with `roadmap`. Substantial changes should be written up as an [RFC](rfcs/README.md).
