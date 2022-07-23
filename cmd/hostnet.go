/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// hostnetCmd represents the hostnet command
var hostnetCmd = &cobra.Command{
	Use:   "hostnet",
	Short: "list pods with host networking",
	Long: `This command returns a list of all the pods in the cluster
	which have host networking enabled.`,
	Run: func(cmd *cobra.Command, args []string) {
		kubeconfig, _ := cmd.Flags().GetString("kubeconfig")
		eathar.Hostnet(kubeconfig)
	},
}

func init() {
	rootCmd.AddCommand(hostnetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostnetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostnetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
