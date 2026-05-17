package main

import (
	"github.com/spf13/cobra"
)

var allNamespaces bool

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Inspect Kuberik resources",
}

var getRolloutsCmd = &cobra.Command{
	Use:     "rollouts",
	Aliases: []string{"rollout", "ro"},
	Short:   "List Rollout resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireKubectl(); err != nil {
			return err
		}
		return kubectl(append([]string{"get", "rollouts.kuberik.com"}, scope()...)...).Run()
	},
}

var getGatesCmd = &cobra.Command{
	Use:     "gates",
	Aliases: []string{"gate"},
	Short:   "List RolloutGate resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireKubectl(); err != nil {
			return err
		}
		return kubectl(append([]string{"get", "rolloutgates.kuberik.com"}, scope()...)...).Run()
	},
}

var getHealthChecksCmd = &cobra.Command{
	Use:     "healthchecks",
	Aliases: []string{"healthcheck", "hc"},
	Short:   "List HealthCheck resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireKubectl(); err != nil {
			return err
		}
		return kubectl(append([]string{"get", "healthchecks.kuberik.com"}, scope()...)...).Run()
	},
}

// scope returns kubectl flags for namespace selection.
func scope() []string {
	if allNamespaces {
		return []string{"--all-namespaces"}
	}
	return []string{"-n", namespace}
}

func init() {
	getCmd.PersistentFlags().BoolVarP(&allNamespaces, "all-namespaces", "A", false, "List resources across all namespaces")
	getCmd.AddCommand(getRolloutsCmd, getGatesCmd, getHealthChecksCmd)
	rootCmd.AddCommand(getCmd)
}
