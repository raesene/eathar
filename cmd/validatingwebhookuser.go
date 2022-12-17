/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// validatingwebhookuserCmd represents the validatingwebhookuser command
var validatingwebhookuserCmd = &cobra.Command{
	Use:   "validatingwebhookuser",
	Short: "List the users that have access to modify validating webhooks",
	Long:  `List the users that have access to modify validating webhooks`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		validatingWebhookUsersList := eathar.ValidatingWebhookUsers(options)
		eathar.ReportRBAC(validatingWebhookUsersList, options, "Users with access to create or modify validatingadmissionwebhookconfigurations")
	},
}

func init() {
	rbacCmd.AddCommand(validatingwebhookuserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validatingwebhookuserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validatingwebhookuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
