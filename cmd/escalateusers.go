/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// escalateusersCmd represents the escalateusers command
var escalateusersCmd = &cobra.Command{
	Use:   "escalateusers",
	Short: "Lists users/groups/service accounts with access to the escalate verb",
	Long:  `Lists users/groups/service accounts with access to the escalate verb`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		eathar.EscalateUsers(options)
	},
}

func init() {
	rbacCmd.AddCommand(escalateusersCmd)

}
