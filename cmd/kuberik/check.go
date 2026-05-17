package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check that the cluster is ready for Kuberik or has Kuberik installed",
	Long: `Verify that the cluster has the required CRDs and controllers running.

Reports the install status of:
  - rollout-controller (required)
  - Kuberik CRDs (Rollout, RolloutGate, HealthCheck, RolloutSchedule)
  - Flux image-reflector-controller (required for ImagePolicy tracking)
  - Integration controllers (optional)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireKubectl(); err != nil {
			return err
		}
		fmt.Println("Checking Kuberik CRDs...")
		if err := kubectl("get", "crd",
			"rollouts.kuberik.com",
			"rolloutgates.kuberik.com",
			"healthchecks.kuberik.com",
			"rolloutschedules.kuberik.com",
		).Run(); err != nil {
			return fmt.Errorf("CRDs not installed: %w", err)
		}
		fmt.Println("\nChecking rollout-controller pod...")
		return kubectl("get", "pods", "-n", "kuberik-system", "-l", "app.kubernetes.io/name=rollout-controller").Run()
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
