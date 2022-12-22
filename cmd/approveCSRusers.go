/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// approvecsrusersCmd represents the approveCSRusers command
var approvecsrusersCmd = &cobra.Command{
	Use:   "approvecsrusers",
	Short: "Lists users who can approve CSRs via update access to the CSR resource",
	Long: `Lists users/groups/service accounts that can approve CSRs via update access to the CSR resource
	This is a slight hack as more rights are needed to approve a CSR than just update access to the CSR resource
	but update is the key permission`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		updateCSRUsersList := eathar.UpdateCSRApproval(options)
		eathar.ReportRBAC(updateCSRUsersList, options, "Users with update rights to CSR approvals")
	},
}

func init() {
	rbacCmd.AddCommand(approvecsrusersCmd)
}
