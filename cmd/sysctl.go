/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// sysctlCmd represents the sysctl command
var sysctlCmd = &cobra.Command{
	Use:   "sysctl",
	Short: "List dangerous sysctls",
	Long: `List sysctls set on pods which are not in the 
	'safe' list`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		sysctls := eathar.Sysctl(options)
		eathar.ReportPSS(sysctls, options, "Unsafe Sysctl")
	},
}

func init() {
	pssCmd.AddCommand(sysctlCmd)
}
