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
		eathar.GetClusterAdminUsers(options)
		eathar.GetSecretsUsers(options)
		eathar.CreatePVUsers(options)
		eathar.ImpersonateUsers(options)
		eathar.EscalateUsers(options)
		eathar.BindUsers(options)
		eathar.ValidatingWebhookUsers(options)
		eathar.MutatingWebhookUsers(options)
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
