/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// clusterUserListCmd represents the clusterUserList command
var clusterUserListCmd = &cobra.Command{
	Use:   "clusterUserList",
	Short: "A list of users defined in cluster role bindings",
	Long: `this command provides a list of any users defined
	in cluster role bindings.`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		userListSlice := eathar.PrincipalList(options, "User")
		eathar.ReportPrincipal(userListSlice, options, "Cluster User List")
	},
}

func init() {
	infoCmd.AddCommand(clusterUserListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterUserListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clusterUserListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
