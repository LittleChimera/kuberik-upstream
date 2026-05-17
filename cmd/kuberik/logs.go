package main

import (
	"github.com/spf13/cobra"
)

var (
	logsFollow    bool
	logsTail      int
	logsContainer string
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Tail the rollout-controller logs",
	Long: `Stream logs from the rollout-controller pod(s) in the kuberik-system namespace.

Equivalent to:
  kubectl logs -n kuberik-system -l app.kubernetes.io/name=rollout-controller`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requireKubectl(); err != nil {
			return err
		}
		kargs := []string{
			"logs", "-n", "kuberik-system",
			"-l", "app.kubernetes.io/name=rollout-controller",
			"--tail", itoa(logsTail),
		}
		if logsContainer != "" {
			kargs = append(kargs, "-c", logsContainer)
		}
		if logsFollow {
			kargs = append(kargs, "-f")
		}
		return kubectl(kargs...).Run()
	},
}

// itoa is a small int->string helper to avoid pulling in strconv via fmt only here.
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}

func init() {
	logsCmd.Flags().BoolVarP(&logsFollow, "follow", "f", false, "Stream logs as they are written")
	logsCmd.Flags().IntVar(&logsTail, "tail", 100, "Number of recent log lines to show before streaming")
	logsCmd.Flags().StringVarP(&logsContainer, "container", "c", "", "Container name (if the pod has multiple)")
	rootCmd.AddCommand(logsCmd)
}
