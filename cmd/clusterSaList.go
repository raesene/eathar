/*
Copyright © 2023 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// clusterSaListCmd represents the clusterSaList command
var clusterSaListCmd = &cobra.Command{
	Use:   "clusterSaList",
	Short: "A list of Service Accounts defined in cluster role bindings",
	Long: `this command provides a list of any service accounts defined
	in cluster role bindings.`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		groupListSlice := eathar.PrincipalList(options, "ServiceAccount")
		eathar.ReportPrincipal(groupListSlice, options, "Cluster Service Account List")
	},
}

func init() {
	infoCmd.AddCommand(clusterSaListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterSaListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clusterSaListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
