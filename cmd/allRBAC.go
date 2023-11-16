/*
Copyright © 2023 Rory McCune <rorym@mccune.org.uk>
*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// allRBACCmd represents the allRBAC command
var allRBACCmd = &cobra.Command{
	Use:   "all",
	Short: "Runs all the RBAC commands",
	Long:  `Runs all the RBAC commands`,
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
		wildcardUsersList := eathar.WildcardAccess(options)
		eathar.ReportRBAC(wildcardUsersList, options, "Users with wildcard access to all resources")
		satokenUsersList := eathar.CreateServiceAccountTokens(options)
		eathar.ReportRBAC(satokenUsersList, options, "Users with create access to service account tokens")
	},
}

func init() {
	rbacCmd.AddCommand(allRBACCmd)

}
