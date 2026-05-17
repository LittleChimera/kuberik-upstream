# Kuberik RFCs

Many new features and enhancements are discussed in [GitHub Discussions](https://github.com/kuberik/kuberik/discussions) or filed directly as issues. For **substantial** changes we use an RFC (Request For Comments) process so users, contributors, and maintainers can align on the direction Kuberik should take.

## When to write an RFC

Substantial changes include:

- New CRDs or new relationships between existing APIs
- Breaking API changes (required fields, removals, kind renames)
- Cross-component changes touching more than one controller
- Security boundaries (RBAC, tenant isolation, secret handling)
- Major UX changes (e.g. new CLI subcommands that change install flow)

If you are unsure, open a discussion first and ask whether your idea needs an RFC.

## RFC Process

1. Open a discussion in [kuberik/kuberik discussions](https://github.com/kuberik/kuberik/discussions) describing the proposal and motivation.
2. Find a maintainer willing to sponsor the RFC. The sponsor will help shepherd the proposal through review.
3. Open a pull request adding a new directory under `rfcs/` using [RFC-0000](RFC-0000/README.md) as template. Use a placeholder number until the PR is ready to merge.
4. Iterate on feedback. Push additional commits rather than force-pushing so the discussion history is preserved.
5. At least one maintainer must approve the proposal before merge. The sponsor assigns a final RFC number and rebases on `main` before merging.
6. Once merged, track implementation against the RFC number (use it as a prefix in issue/PR titles).
7. After implementation lands in a release, update the RFC's "Implementation History" with the version.

An RFC may be discarded during implementation if security, performance, or scope concerns surface. The "Implementation History" section should record the rejection rationale, and a follow-up RFC can address those concerns.
