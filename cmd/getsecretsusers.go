/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// getsecretsusersCmd represents the getsecretsusers command
var getsecretsusersCmd = &cobra.Command{
	Use:   "getsecretsusers",
	Short: "Lists users/groups/service accounts with access to read secrets",
	Long: `This command lists users/groups/service accounts with access to read secrets
	either via the get verb or via the list verb (both of which allow you to read the contents of the secret).`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		eathar.GetSecretsUsers(options)
	},
}

func init() {
	rbacCmd.AddCommand(getsecretsusersCmd)
}
