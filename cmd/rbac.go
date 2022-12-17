/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// rbacCmd represents the rbac command
var rbacCmd = &cobra.Command{
	Use:   "rbac",
	Short: "Runs all the RBAC checks",
	Long: `This runs all the RBAC commands, 
	you can run each check individually if you wish, using the subcommands`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		clusterAdminRoleBindingList := eathar.GetClusterAdminUsers(options)
		eathar.ReportRBAC(clusterAdminRoleBindingList, options, "Cluster Admin Users")
		getSecretsUsersList := eathar.GetSecretsUsers(options)
		eathar.ReportRBAC(getSecretsUsersList, options, "Users with access to secrets")
		createPVUsersList := eathar.CreatePVUsers(options)
		eathar.ReportRBAC(createPVUsersList, options, "Users with access to create persistent volumes")
		impersonateUsersList := eathar.ImpersonateUsers(options)
		eathar.ReportRBAC(impersonateUsersList, options, "Users with access to impersonate")
		escalateUsersList := eathar.EscalateUsers(options)
		eathar.ReportRBAC(escalateUsersList, options, "Users with access to escalate")
		bindUsersList := eathar.BindUsers(options)
		eathar.ReportRBAC(bindUsersList, options, "Users with access to bind")
		validatingWebhookUsersList := eathar.ValidatingWebhookUsers(options)
		eathar.ReportRBAC(validatingWebhookUsersList, options, "Users with access to create or modify validatingadmissionwebhookconfigurations")
		mutatingWebhookUsersList := eathar.MutatingWebhookUsers(options)
		eathar.ReportRBAC(mutatingWebhookUsersList, options, "Users with access to create or modify mutatingadmissionwebhookconfigurations")
	},
}

func init() {
	rootCmd.AddCommand(rbacCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rbacCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rbacCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
