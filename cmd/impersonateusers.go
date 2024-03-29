/*
Copyright © 2023 Rory McCune <rorym@mccune.org.uk>
*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// impersonateusersCmd represents the impersonateusers command
var impersonateusersCmd = &cobra.Command{
	Use:   "impersonateusers",
	Short: "Lists users/groups/service accounts with access to the impersonate verb",
	Long:  `Lists users/groups/service accounts with access to the impersonate verb`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		impersonateUsersList := eathar.ImpersonateUsers(options)
		eathar.ReportRBAC(impersonateUsersList, options, "Users with access to impersonate")
	},
}

func init() {
	rbacCmd.AddCommand(impersonateusersCmd)

}
