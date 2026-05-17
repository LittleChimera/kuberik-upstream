# Governance

Kuberik is an open project with a small, active maintainer group. This document describes how decisions get made and how the project is run today.

## Principles

- **Open by default.** Discussions, RFCs, design notes, and roadmaps live in public.
- **Code review is consensus.** Every change ships through a pull request reviewed by at least one maintainer.
- **Lazy consensus for small things.** A maintainer approves and merges trivial changes (typos, dependency bumps, single-file fixes) without further discussion.
- **RFC for substantial changes.** New CRDs, breaking API changes, and cross-component behavior changes go through the [RFC process](rfcs/README.md).

## Roles

### Contributor

Anyone who opens an issue, joins a discussion, or sends a pull request. There is no formal onboarding - read [CONTRIBUTING.md](CONTRIBUTING.md) and dive in.

### Maintainer

Listed in [MAINTAINERS](MAINTAINERS). Maintainers can:

- Merge pull requests
- Triage issues and assign labels
- Sponsor and approve RFCs
- Cut releases

To become a maintainer, a contributor should demonstrate sustained involvement (typically a few merged PRs or sustained help in Discussions) and be nominated by an existing maintainer. Nominations happen via PR against the [MAINTAINERS](MAINTAINERS) file.

## Decision making

Most decisions are reached through pull request review and inline discussion. For larger questions:

1. Open a [GitHub Discussion](https://github.com/kuberik/kuberik/discussions) and tag the relevant area.
2. If consensus is not reached in discussion, write an [RFC](rfcs/README.md).
3. Maintainers approve the RFC; lazy consensus applies if no maintainer objects within a week.
4. If maintainers disagree, the project lead breaks the tie. The current project lead is the first maintainer listed in [MAINTAINERS](MAINTAINERS).

## Releases

Each Kuberik component (rollout-controller, datadog-controller, etc.) releases independently from its own repository. The `kuberik/kuberik` repo releases the CLI and the bundled install manifest that pins compatible versions of each controller.

Releases follow [SemVer](https://semver.org/). Pre-1.0 versions may include breaking changes between minor versions; we will call those out in release notes.

## Code of conduct

All project participants are expected to follow the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/main/code-of-conduct.md). Code of Conduct issues can be reported privately via [GitHub Security Advisories](https://github.com/kuberik/kuberik/security/advisories/new).
