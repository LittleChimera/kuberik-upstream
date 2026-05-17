package main

import (
	"github.com/spf13/cobra"
)

var suspendCmd = &cobra.Command{
	Use:   "suspend ROLLOUT",
	Short: "Pause a Rollout (no new releases will be promoted)",
	Long: `Annotate a Rollout with rollout.kuberik.com/suspended=true.

The rollout-controller treats a suspended Rollout as paused: gates and
health checks continue to be evaluated, but no new release is patched
onto the target resource until the suspension is lifted.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireKubectl(); err != nil {
			return err
		}
		return kubectl("-n", namespace, "annotate", "--overwrite",
			"rollout.kuberik.com", args[0],
			"rollout.kuberik.com/suspended=true").Run()
	},
}

var resumeCmd = &cobra.Command{
	Use:   "resume ROLLOUT",
	Short: "Resume a previously suspended Rollout",
	Long:  `Remove the rollout.kuberik.com/suspended annotation so the rollout-controller proceeds again.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireKubectl(); err != nil {
			return err
		}
		return kubectl("-n", namespace, "annotate",
			"rollout.kuberik.com", args[0],
			"rollout.kuberik.com/suspended-").Run()
	},
}

var describeCmd = &cobra.Command{
	Use:   "describe KIND NAME",
	Short: "Describe a Kuberik resource (rollout, gate, healthcheck)",
	Long: `Describe a Kuberik resource by short kind name. Equivalent to
"kubectl describe <kind>.kuberik.com <name>".`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireKubectl(); err != nil {
			return err
		}
		kind := args[0]
		switch kind {
		case "rollout", "rollouts", "ro":
			kind = "rollouts.kuberik.com"
		case "gate", "gates":
			kind = "rolloutgates.kuberik.com"
		case "healthcheck", "healthchecks", "hc":
			kind = "healthchecks.kuberik.com"
		case "schedule", "schedules":
			kind = "rolloutschedules.kuberik.com"
		}
		return kubectl("-n", namespace, "describe", kind, args[1]).Run()
	},
}

func init() {
	rootCmd.AddCommand(suspendCmd, resumeCmd, describeCmd)
}
