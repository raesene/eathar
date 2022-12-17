/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// procmountCmd represents the procmount command
var procmountCmd = &cobra.Command{
	Use:   "procmount",
	Short: "List containers with unmasked proc mounts",
	Long: `This command lists containers with unmasked proc mounts. This is a security risk as it allows
	access to the proc filesystem on the host which can contain sensitive information`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		unmaskedproc := eathar.Procmount(options)
		eathar.ReportPSS(unmaskedproc, options, "Unmasked Procmount")
	},
}

func init() {
	pssCmd.AddCommand(procmountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// procmountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// procmountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
