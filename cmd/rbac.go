/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rbacCmd represents the rbac command
var rbacCmd = &cobra.Command{
	Use:   "rbac",
	Short: "Checks related to cluster RBAC",
	Long: `This command runs RBAC checks on the target cluster., 
	you can use the all command to run all the checks, or run each check individually`,
	Run: func(cmd *cobra.Command, args []string) {
		//return help for RBAC command
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(rbacCmd)

}
