/*
Copyright © 2023 Rory McCune <rorym@mccune.org.uk>
*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// hostprocessCmd represents the hostprocess command
var hostprocessCmd = &cobra.Command{
	Use:   "hostprocess",
	Short: "List hostProcess Windows pods",
	Long: `Lists hostProcess Windows pods. This is a security risk as it allows
	full access to the underlying node. This is effectively the Windows equivalent
	of privileged containers`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		hostprocesscont := eathar.HostProcess(options)
		eathar.ReportPSS(hostprocesscont, options, "Host Process")
	},
}

func init() {
	pssCmd.AddCommand(hostprocessCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostprocessCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostprocessCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
