/*
Copyright © 2023 Rory McCune <rorym@mccune.org.uk>
*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// bindusersCmd represents the bindusers command
var bindusersCmd = &cobra.Command{
	Use:   "bindusers",
	Short: "Lists users/groups/service accounts with access to the bind verb",
	Long:  `Lists users/groups/service accounts with access to the bind verb`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		bindUsersList := eathar.BindUsers(options)
		eathar.ReportRBAC(bindUsersList, options, "Users with access to bind")
	},
}

func init() {
	rbacCmd.AddCommand(bindusersCmd)

}
