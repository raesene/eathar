/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// clusteradminusersCmd represents the clusteradminusers command
var clusteradminusersCmd = &cobra.Command{
	Use:   "clusteradminusers",
	Short: "A list of users/groups/service accounts with cluster-admin role",
	Long:  `This provides a list of users/groups/service accounts with cluster-admin role`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		eathar.GetClusterAdminUsers(options)
	},
}

func init() {
	rootCmd.AddCommand(clusteradminusersCmd)
}
