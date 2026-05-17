package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the kuberik CLI version",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("kuberik %s\n", Version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
