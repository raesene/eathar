/*
Copyright © 2023 Rory McCune <rorym@mccune.org.uk>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// pssCmd represents the pss command
var pssCmd = &cobra.Command{
	Use:   "pss",
	Short: "Checks relating to Pod Security Standards",
	Long: `These commands run Pod Security checks on the target cluster.
	  you can use the all command to run all the checks, or run each check individually`,
	Run: func(cmd *cobra.Command, args []string) {
		//return the help for the pss command
		cmd.Help()

	},
}

func init() {
	rootCmd.AddCommand(pssCmd)

}
