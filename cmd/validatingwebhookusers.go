/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// validatingwebhookuserCmd represents the validatingwebhookusers command
var validatingwebhookusersCmd = &cobra.Command{
	Use:   "validatingwebhookusers",
	Short: "List the users that have access to modify validating webhooks",
	Long:  `List the users that have access to modify validating webhooks`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		validatingWebhookUsersList := eathar.ValidatingWebhookUsers(options)
		eathar.ReportRBAC(validatingWebhookUsersList, options, "Users with access to create or modify validatingadmissionwebhookconfigurations")
	},
}

func init() {
	rbacCmd.AddCommand(validatingwebhookusersCmd)

}
