/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// persistentvolumecreationuserCmd represents the persistentvolumecreationuser command
var persistentvolumecreationusersCmd = &cobra.Command{
	Use:   "persistentvolumecreationusers",
	Short: "Lists users/groups/service accounts with access to create persistent volumes",
	Long:  `This command lists users/groups/service accounts with access to create persistent volumes`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		createPVUsersList := eathar.CreatePVUsers(options)
		eathar.ReportRBAC(createPVUsersList, options, "Users with access to create persistent volumes")
	},
}

func init() {
	rbacCmd.AddCommand(persistentvolumecreationusersCmd)

}
