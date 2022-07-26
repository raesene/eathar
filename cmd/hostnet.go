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
		//Get the value of the kubeconfig flag so we can pass it to the command
		kubeconfig, _ := cmd.Flags().GetString("kubeconfig")
		jsonrep, _ := cmd.Flags().GetBool("jsonrep")
		eathar.Hostnet(kubeconfig, jsonrep)
	},
}

func init() {
	rootCmd.AddCommand(hostnetCmd)
}
