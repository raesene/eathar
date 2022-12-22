/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// wildcardusersCmd represents the wildcarduser command
var wildcardusersCmd = &cobra.Command{
	Use:   "wildcardusers",
	Short: "List all users with wildcard permissions to all resources",
	Long: `This command finds clusterroles that provide access to all resources
	via wildcard (*), and then lists all users/groups/service accounts associated
	with those clusterroles via clusterrolebindings.`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		wildcardUsersList := eathar.WildcardAccess(options)
		eathar.ReportRBAC(wildcardUsersList, options, "Users with wildcard access to all resources")
	},
}

func init() {
	rbacCmd.AddCommand(wildcardusersCmd)

}
