# FAQ

## How is Kuberik different from Argo Rollouts?

Argo Rollouts owns the **deployment strategy** (canary, blue-green) and traffic shifting for a single workload. Kuberik owns **promotion control** - when a new image is allowed to enter the deployment, under what gating and bake conditions. Many teams run both. See [Migration Guide](migration.md) for the detailed mapping.

## How is Kuberik different from Flagger?

Flagger automates progressive delivery on a single workload (traffic shift + metric analysis + rollback) in a monolithic CRD. Kuberik decomposes those concerns into independent CRDs (`Rollout`, `RolloutGate`, `HealthCheck`) so producers (Datadog, Prometheus, custom controllers) can compose freely.

## Do I have to use Flux?

Yes, currently. Kuberik reads release information from Flux's `ImagePolicy` resource and patches Flux's `OCIRepository` / `Kustomization` when promoting. Future versions may add adapters for other GitOps engines, but Flux is the only first-class integration today.

## Do I have to use all the controllers?

No. The **rollout-controller** is the only required component. Integration controllers (datadog, prometheus, openkruise, environment) are optional - install only the ones you need.

## Can I write my own integration controller?

Yes - that's the design. Any controller that creates `HealthCheck` or `RolloutGate` resources can participate. See [docs/healthchecks.md](healthchecks.md) for the contract.

## Why does my rollout never promote?

Most common causes, in order:

1. The `ImagePolicy` has no `latestRef.tag` set - check the upstream `ImageRepository` and your registry.
2. A `RolloutGate` is `passing: false` - run `kuberik tree my-app` to see which.
3. A `HealthCheck` is unhealthy or `lastErrorTime` is more recent than bake start.
4. The target `OCIRepository` / `Kustomization` is missing the `rollout.kuberik.com/rollout` annotation.

Full diagnostic guide in [docs/troubleshooting.md](troubleshooting.md).

## Why does the bake fail even though my metrics recovered?

`HealthCheck.status.lastErrorTime` is a **witness** that survives recovery. If a metric flapped during the bake window, even briefly, the rollout-controller treats that as a bake failure. This is intentional - "saw a real error during bake" is a stronger signal than "currently healthy at this exact moment." See [docs/healthchecks.md](healthchecks.md).

## Can I use Kuberik without the CLI?

Yes. Everything Kuberik does is reachable through `kubectl`. The CLI is a convenience layer.

## How do I bypass gates for an emergency deploy?

```bash
kubectl annotate rollout my-app rollout.kuberik.com/bypass-gates="<version>"
```

This deploys the specified version regardless of gate state. **Use carefully** - it skips your safety net. Remove the annotation when the emergency is resolved.

## Why is the Helm chart at 0.x?

The chart's API (values keys) is still settling. Once we have a few production users for ~2 quarters and the integration controller templates are stable, we'll tag 1.0.

## Where can I get help?

- [GitHub Discussions](https://github.com/kuberik/kuberik/discussions) - usage questions, design ideas
- [GitHub Issues](https://github.com/kuberik/kuberik/issues) - reproducible bugs, feature requests
- For private security reports: [Report a vulnerability](https://github.com/kuberik/kuberik/security/advisories/new)

## How do I contribute?

See [CONTRIBUTING.md](../CONTRIBUTING.md). Substantial changes go through the [RFC process](../rfcs/README.md).
