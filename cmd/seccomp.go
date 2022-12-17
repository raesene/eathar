/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// seccompCmd represents the seccomp command
var seccompCmd = &cobra.Command{
	Use:   "seccomp",
	Short: "Check for disabled seccomp",
	Long: `Checks whether a seccomp profile has been set. By default
	Kubernete disables CRI seccomp profiles (e.g. Docker)`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		seccomp := eathar.Seccomp(options)
		eathar.ReportPSS(seccomp, options, "Seccomp Disabled")
	},
}

func init() {
	pssCmd.AddCommand(seccompCmd)

}
