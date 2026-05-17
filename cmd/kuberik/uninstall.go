package main

import (
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Remove Kuberik from the current cluster",
	Long: `Remove Kuberik resources from the cluster.

This deletes the rollout-controller and all installed integration controllers
along with their CRDs. Rollout, RolloutGate, and HealthCheck resources will be
deleted with the CRDs.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireKubectl(); err != nil {
			return err
		}
		return kubectl("delete", "-k", allInstallURL, "--ignore-not-found").Run()
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
