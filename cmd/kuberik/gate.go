package main

import (
	"github.com/spf13/cobra"
)

var approveCmd = &cobra.Command{
	Use:   "approve GATE",
	Short: "Set a RolloutGate to passing",
	Long: `Patch a RolloutGate so that spec.passing is true, allowing the rollout
to proceed past the gate.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireKubectl(); err != nil {
			return err
		}
		return kubectl("-n", namespace, "patch", "rolloutgate.kuberik.com", args[0],
			"--type", "merge", "-p", `{"spec":{"passing":true}}`).Run()
	},
}

var rejectCmd = &cobra.Command{
	Use:   "reject GATE",
	Short: "Set a RolloutGate to not passing",
	Long: `Patch a RolloutGate so that spec.passing is false, blocking the rollout
from proceeding past the gate.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireKubectl(); err != nil {
			return err
		}
		return kubectl("-n", namespace, "patch", "rolloutgate.kuberik.com", args[0],
			"--type", "merge", "-p", `{"spec":{"passing":false}}`).Run()
	},
}

func init() {
	rootCmd.AddCommand(approveCmd, rejectCmd)
}
