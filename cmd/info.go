/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Checks that provide general information about a cluster",
	Long: `These commands return general information about a cluster.
	you can use the all command to run all the checks, or run each check individually`,
	Run: func(cmd *cobra.Command, args []string) {
		//return help for info command
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

}
