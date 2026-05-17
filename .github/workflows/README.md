# GitHub Workflows

Reference for every workflow in this directory.

| Workflow | Trigger | Role |
|---|---|---|
| `ci.yaml` | push to main, PR | Build, vet, test, lint the CLI; validate the install manifest |
| `e2e.yaml` | push to main, PR | Stand up a kind cluster, install Flux + Kuberik, run smoke tests |
| `codeql.yaml` | push, PR, weekly | CodeQL static analysis of Go source |
| `scorecard.yaml` | push to main, weekly | OpenSSF Scorecard security posture analysis |
| `trivy.yaml` | push, Dockerfile/go.mod PRs, weekly | Trivy filesystem vuln scan, SARIF to code-scanning |
| `release.yaml` | push of `v*` tag | Goreleaser run: CLI binaries, Homebrew formula, ghcr.io images |
| `sync-labels.yaml` | push to main touching `labels.yaml` | Apply `.github/labels.yaml` to issue/PR label set |
| `stale.yaml` | daily cron | Mark and close inactive issues / PRs |

## Adding a workflow

- Pin third-party actions to a SHA (Dependabot keeps them current).
- Declare top-level `permissions: read-all` and grant elevated permissions per-job, not at the workflow scope.
- Prefer `actions/setup-go` with `go-version-file: go.mod` so the Go version stays in sync with the module.
- Add the workflow here in the table.
