/*
Copyright © 2023 Rory McCune <rorym@mccune.org.uk>
*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// escalateusersCmd represents the escalateusers command
var escalateusersCmd = &cobra.Command{
	Use:   "escalateusers",
	Short: "Lists users/groups/service accounts with access to the escalate verb",
	Long:  `Lists users/groups/service accounts with access to the escalate verb`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		escalateUsersList := eathar.EscalateUsers(options)
		eathar.ReportRBAC(escalateUsersList, options, "Users with access to escalate")
	},
}

func init() {
	rbacCmd.AddCommand(escalateusersCmd)

}
