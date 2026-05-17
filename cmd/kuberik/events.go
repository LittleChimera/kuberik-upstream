package main

import (
	"github.com/spf13/cobra"
)

var eventsWatch bool

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Show cluster events for Kuberik resources",
	Long: `Show Kubernetes events whose involved object is a Kuberik resource
(Rollout, RolloutGate, HealthCheck, RolloutSchedule).

Useful for diagnosing why a Rollout is not advancing:
controller events explain gate decisions, bake failures, and
target-patch errors.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireKubectl(); err != nil {
			return err
		}
		kargs := []string{
			"get", "events",
			"--field-selector",
			"involvedObject.apiVersion=kuberik.com/v1alpha1",
			"--sort-by", ".lastTimestamp",
		}
		kargs = append(kargs, scope()...)
		if eventsWatch {
			kargs = append(kargs, "--watch")
		}
		return kubectl(kargs...).Run()
	},
}

func init() {
	eventsCmd.Flags().BoolVarP(&eventsWatch, "watch", "w", false, "Stream new events as they arrive")
	eventsCmd.Flags().BoolVarP(&allNamespaces, "all-namespaces", "A", false, "List events across all namespaces")
	rootCmd.AddCommand(eventsCmd)
}
