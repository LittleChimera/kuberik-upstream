# GitHub Workflows

Reference for every workflow in this directory.

| Workflow | Trigger | Role |
|---|---|---|
| `ci.yaml` | push to main, PR | Build/vet/test/lint CLI; validate kustomize install manifest; lint, template, `helm install`, and `helm upgrade` the chart on kind |
| `e2e.yaml` | push to main, PR | Stand up a kind cluster, install Flux + Kuberik, run smoke tests |
| `codeql.yaml` | push, PR, weekly | CodeQL static analysis of Go source |
| `scorecard.yaml` | push to main, weekly | OpenSSF Scorecard security posture analysis |
| `trivy.yaml` | push, Dockerfile/go.mod PRs, weekly | Trivy filesystem vuln scan, SARIF to code-scanning |
| `lint-docs.yaml` | push, PR touching `**.md` | markdownlint-cli2 against the configured rules |
| `chart-release.yaml` | push to main touching `chart/**` | Publish the Helm chart to `gh-pages` via chart-releaser |
| `release.yaml` | push of `v*` tag | Goreleaser run: CLI binaries, Homebrew formula, ghcr.io images |
| `sync-labels.yaml` | push to main touching `labels.yaml` | Apply `.github/labels.yaml` to issue/PR label set |
| `pr-labeler.yaml` | PR opened / synchronize | Auto-apply `area/*` labels to PRs based on `.github/labeler.yml` path patterns |
| `stale.yaml` | daily cron | Mark and close inactive issues / PRs |
| `check-component-releases.yaml` | weekly cron | Diff pinned controller versions vs upstream; open tracking issue if behind |

## Adding a workflow

- Pin third-party actions to a SHA (Dependabot keeps them current).
- Declare top-level `permissions: read-all` and grant elevated permissions per-job, not at the workflow scope.
- Prefer `actions/setup-go` with `go-version-file: go.mod` so the Go version stays in sync with the module.
- Add the workflow here in the table.
