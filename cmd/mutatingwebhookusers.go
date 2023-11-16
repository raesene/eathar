/*
Copyright © 2023 Rory McCune <rorym@mccune.org.uk>
*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// mutatingwebhookuserCmd represents the mutatingwebhookusers command
var mutatingwebhookusersCmd = &cobra.Command{
	Use:   "mutatingwebhookusers",
	Short: "List the users that have access to modify mutating webhooks",
	Long:  `List the users that have access to modify mutating webhooks`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		mutatingWebhookUsersList := eathar.MutatingWebhookUsers(options)
		eathar.ReportRBAC(mutatingWebhookUsersList, options, "Users with access to create or modify mutatingadmissionwebhookconfigurations")
	},
}

func init() {
	rbacCmd.AddCommand(mutatingwebhookusersCmd)

}
