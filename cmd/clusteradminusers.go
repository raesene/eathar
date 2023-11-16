/*
Copyright © 2023 Rory McCune <rorym@mccune.org.uk>
*/

package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// clusteradminusersCmd represents the clusteradminusers command
var clusteradminusersCmd = &cobra.Command{
	Use:   "clusteradminusers",
	Short: "A list of users/groups/service accounts with cluster-admin role",
	Long:  `This provides a list of users/groups/service accounts with cluster-admin role`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		clusterAdminRoleBindingList := eathar.GetClusterAdminUsers(options)
		eathar.ReportRBAC(clusterAdminRoleBindingList, options, "Cluster Admin Users")
	},
}

func init() {
	rbacCmd.AddCommand(clusteradminusersCmd)
}
