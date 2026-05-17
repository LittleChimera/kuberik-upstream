# Contributing to Kuberik

## Filing issues

Use [GitHub Issues](https://github.com/kuberik/kuberik/issues) for bugs, feature requests, and questions. Check existing issues before opening a new one.

## Where the code lives

| Component | Repository |
|---|---|
| Core controller (CRDs, reconcilers) | [rollout-controller](https://github.com/kuberik/rollout-controller) |
| Web dashboard | [rollout-dashboard](https://github.com/kuberik/rollout-dashboard) |
| Datadog integration | [datadog-controller](https://github.com/kuberik/datadog-controller) |
| GitHub Deployments integration | [environment-controller](https://github.com/kuberik/environment-controller) |
| OpenKruise integration | [openkruise-controller](https://github.com/kuberik/openkruise-controller) |
| Prometheus integration | [prometheus-controller](https://github.com/kuberik/prometheus-controller) |

This repo (`kuberik/kuberik`) holds the lighthouse README, install manifests, Helm chart, and documentation. For code contributions, open a PR in the relevant sub-repo above.

## Proposing changes

1. Open an issue describing the change and the motivation.
2. Fork the relevant sub-repo and create a branch.
3. Make your changes with tests.
4. Open a pull request referencing the issue.

## Code of conduct

This project follows the [Contributor Covenant](CODE_OF_CONDUCT.md).
