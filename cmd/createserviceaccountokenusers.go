/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// createserviceaccountokenusersCmd represents the createserviceaccountokenusers command
var createserviceaccountokenusersCmd = &cobra.Command{
	Use:   "createserviceaccountokenusers",
	Short: "Lists users who can create service account tokens",
	Long:  `Lists users who can create service account tokens at the cluster level.`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		satokenUsersList := eathar.CreateServiceAccountTokens(options)
		eathar.ReportRBAC(satokenUsersList, options, "Users with create access to service account tokens")
	},
}

func init() {
	rbacCmd.AddCommand(createserviceaccountokenusersCmd)

}
