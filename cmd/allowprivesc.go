/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// allowprivescCmd represents the allowprivesc command
var allowprivescCmd = &cobra.Command{
	Use:   "allowprivesc",
	Short: "List Pods that allow privilege escalation",
	Long: `This command lists posts that allow privilege escalation.
	This is a default in general for linux container runtimes and allows
	for things like sudo to be used in a container to escalate privileges`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		eathar.AllowPrivEsc(options)
	},
}

func init() {
	rootCmd.AddCommand(allowprivescCmd)

}
