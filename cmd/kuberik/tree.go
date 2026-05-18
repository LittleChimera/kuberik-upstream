package main

import (
	"github.com/spf13/cobra"
)

var treeCmd = &cobra.Command{
	Use:   "tree ROLLOUT",
	Short: "Print the gate and health-check tree for a Rollout",
	Long: `Print a tree view of a Rollout and every RolloutGate / HealthCheck
that references it. Useful for diagnosing why a Rollout is blocked.

Example:
  $ kuberik tree my-app -n production
  Rollout/my-app
  ├── Gates
  │   ├── RolloutGate/business-hours (passing=true)
  │   └── RolloutGate/manual-approval (passing=false)  <-- blocking
  └── HealthChecks
      ├── HealthCheck/error-rate (healthy=true)
      └── HealthCheck/latency-p99 (healthy=true)
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireKubectl(); err != nil {
			return err
		}
		name := args[0]
		out := cmd.OutOrStdout()
		writeln(out, "Rollout/"+name)
		writeln(out, "├── Gates")
		// gates filtered by spec.rolloutRef.name
		if err := kubectl("-n", namespace, "get", "rolloutgates.kuberik.com",
			"-o", "jsonpath={range .items[?(@.spec.rolloutRef.name==\""+name+"\")]}│   ├── RolloutGate/{.metadata.name} (passing={.spec.passing}){\"\\n\"}{end}").Run(); err != nil {
			return err
		}
		writeln(out, "└── HealthChecks")
		if err := kubectl("-n", namespace, "get", "healthchecks.kuberik.com",
			"-o", "jsonpath={range .items[?(@.spec.rolloutRef.name==\""+name+"\")]}    ├── HealthCheck/{.metadata.name} (healthy={.status.healthy}){\"\\n\"}{end}").Run(); err != nil {
			return err
		}
		return nil
	},
}

func writeln(w interface{ Write([]byte) (int, error) }, s string) {
	_, _ = w.Write([]byte(s + "\n"))
}

func init() {
	rootCmd.AddCommand(treeCmd)
}
