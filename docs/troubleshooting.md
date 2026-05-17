# Troubleshooting

Common issues with running Kuberik and how to diagnose them.

## A rollout never promotes

Symptoms: `kuberik get rollouts` shows the Rollout but `status.latestRelease` stays empty or stale even though a newer image tag exists.

Check, in order:

1. **The ImagePolicy has a `latestRef`.**
   ```bash
   kubectl get imagepolicy -n flux-system <name> -o yaml
   ```
   `status.latestRef.tag` must be set. If empty, the image-reflector-controller has not seen any matching tags. Check the `ImageRepository` and your image registry.

2. **The Rollout points at the right ImagePolicy.**
   `spec.releasesImagePolicy.{name,namespace}` must match the existing ImagePolicy.

3. **No gate is blocking the rollout.**
   ```bash
   kuberik get gates -n <namespace>
   kubectl describe rollout <name> -n <namespace>
   ```
   The `Promoting` condition lists the blocking gate by name. Approve it with `kuberik approve <gate>` if appropriate.

4. **No HealthCheck is unhealthy.**
   ```bash
   kuberik get healthchecks -n <namespace>
   ```
   A failing health check during bake will pause promotion. Look at the `message` field for the producing controller's reason.

5. **The target resource has the annotation.**
   The OCIRepository or Kustomization the Rollout patches must carry either:
   - `rollout.kuberik.com/rollout: "<rollout-name>"`, or
   - `rollout.kuberik.com/substitute.<VAR>.from: "<image-policy-name>"`
   Without the annotation, the controller has no target to patch.

## `kuberik check` says CRDs are missing

The install did not run, or it ran against a different cluster.

```bash
# Verify current context
kubectl config current-context

# Reapply
kuberik install
```

If you see "ImagePolicy not registered", the Flux image-reflector-controller is missing. Install Flux with the image-reflector component (see [installation](installation.md)).

## A bake fails immediately

If a Rollout enters bake and fails within seconds with no obvious unhealthy HealthCheck:

- Inspect `lastErrorTime` on each HealthCheck referenced by the Rollout. **`lastErrorTime` survives recovery** - if a check went unhealthy at some point and then recovered, `lastErrorTime` is still set and the rollout-controller treats this as a witness of trouble during bake.
- To clear, the producing controller must update the HealthCheck status with a `lastErrorTime` older than the bake start, or you must wait for the bake start to advance past the recorded error.

This is intentional, not a bug. See [docs/healthchecks.md](healthchecks.md).

## Controller pod crashlooping

```bash
kubectl logs -n kuberik-system deploy/rollout-controller --tail=200
```

Common causes:

- **RBAC missing for a CRD.** Usually means an integration controller installed CRDs after the rollout-controller started but the controller's ServiceAccount cannot list them. Restart the rollout-controller pod.
- **Old install bundle.** If you mixed-and-matched per-component install URLs, the version pinning may be inconsistent. Reapply the full bundle:
  ```bash
  kuberik install --all
  ```

## Manual approval gate keeps re-blocking

If `kuberik approve <gate>` is overridden right after you run it, a `RolloutSchedule` is probably reconciling the gate back to its scheduled value. Check:

```bash
kubectl get rolloutschedule -A
kubectl get clusterrolloutschedule
```

If a schedule owns the gate, change the schedule rather than the gate.

## Asking for help

- [GitHub Discussions](https://github.com/kuberik/kuberik/discussions) for usage questions.
- [GitHub Issues](https://github.com/kuberik/kuberik/issues) for reproducible bugs.

Include in your report: Kuberik CLI version, rollout-controller image version, Kubernetes version, and the YAML of the Rollout / Gate / HealthCheck you are seeing problems with.
