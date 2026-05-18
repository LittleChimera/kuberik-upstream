package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	initName                 string
	initImagePolicy          string
	initImagePolicyNamespace string
	initBakeTime             string
)

var initCmd = &cobra.Command{
	Use:   "init KIND",
	Short: "Scaffold a starter manifest for a Kuberik resource",
	Long: `Print a starter manifest to stdout for one of the Kuberik kinds.

Examples:
  kuberik init rollout --name my-app > rollout.yaml
  kuberik init gate --name my-app-approval > gate.yaml
  kuberik init schedule --name business-hours > schedule.yaml

The KIND must be one of: rollout, gate, schedule.`,
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	ValidArgs: []string{"rollout", "gate", "schedule"},
	RunE: func(cmd *cobra.Command, args []string) error {
		var tmpl string
		switch args[0] {
		case "rollout":
			tmpl = fmt.Sprintf(rolloutTemplate, initName, namespace, initImagePolicy, initBakeTime)
		case "gate":
			tmpl = fmt.Sprintf(gateTemplate, initName, namespace, initName)
		case "schedule":
			tmpl = fmt.Sprintf(scheduleTemplate, initName, namespace, initName)
		default:
			return fmt.Errorf("unknown kind: %s", args[0])
		}
		_, err := os.Stdout.WriteString(tmpl)
		return err
	},
}

const rolloutTemplate = `apiVersion: kuberik.com/v1alpha1
kind: Rollout
metadata:
  name: %s
  namespace: %s
spec:
  releasesImagePolicy:
    name: %s
  versionHistoryLimit: 10
  bakeTime: %s
`

const gateTemplate = `apiVersion: kuberik.com/v1alpha1
kind: RolloutGate
metadata:
  name: %s
  namespace: %s
spec:
  rolloutRef:
    name: %s
  # Flip to true (e.g. via kuberik approve) when ready to promote.
  passing: false
`

const scheduleTemplate = `apiVersion: kuberik.com/v1alpha1
kind: RolloutSchedule
metadata:
  name: %s
  namespace: %s
spec:
  rolloutRef:
    name: %s
  # Allow: open during the window, closed outside.
  # Deny:  closed during the window, open outside.
  action: Allow
  schedule:
    - days: [Mon, Tue, Wed, Thu, Fri]
      start: "09:00"
      end: "17:00"
      timezone: "America/New_York"
`

func init() {
	initCmd.Flags().StringVar(&initName, "name", "my-app", "Name to set in metadata.name")
	initCmd.Flags().StringVar(&initImagePolicy, "image-policy", "my-app", "ImagePolicy name (Rollout only). Must live in the same namespace as the Rollout in the current CRD; this flag stays for forward-compat.")
	initCmd.Flags().StringVar(&initImagePolicyNamespace, "image-policy-namespace", "", "(unused) reserved for forward-compat")
	initCmd.Flags().StringVar(&initBakeTime, "bake-time", "10m", "Bake duration (Rollout only)")
	rootCmd.AddCommand(initCmd)
}
