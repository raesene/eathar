/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// mutatingwebhookuserCmd represents the mutatingwebhookuser command
var mutatingwebhookuserCmd = &cobra.Command{
	Use:   "mutatingwebhookuser",
	Short: "List the users that have access to modify mutating webhooks",
	Long:  `List the users that have access to modify mutating webhooks`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		mutatingWebhookUsersList := eathar.MutatingWebhookUsers(options)
		eathar.ReportRBAC(mutatingWebhookUsersList, options, "Users with access to create or modify mutatingadmissionwebhookconfigurations")
	},
}

func init() {
	rbacCmd.AddCommand(mutatingwebhookuserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mutatingwebhookuserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mutatingwebhookuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
