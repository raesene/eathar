/*
Copyright © 2023 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// clusterGroupListCmd represents the clusterGroupList command
var clusterGroupListCmd = &cobra.Command{
	Use:   "clusterGroupList",
	Short: "A list of Groups defined in cluster role bindings",
	Long: `this command provides a list of any groups defined
	in cluster role bindings.`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		groupListSlice := eathar.PrincipalList(options, "Group")
		eathar.ReportPrincipal(groupListSlice, options, "Cluster Group List")
	},
}

func init() {
	infoCmd.AddCommand(clusterGroupListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterGroupListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clusterGroupListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
