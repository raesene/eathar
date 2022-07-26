/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// privilegedCmd represents the privileged command
var privilegedCmd = &cobra.Command{
	Use:   "privileged",
	Short: "List Privileged containers",
	Long: `Lists privileged containers. Containers which run
	as privileged can easily break out to the underlying host
	so should be used only where expicitly required.`,
	Run: func(cmd *cobra.Command, args []string) {
		//Get the value of the kubeconfig flag so we can pass it to the command
		kubeconfig, _ := cmd.Flags().GetString("kubeconfig")
		jsonrep, _ := cmd.Flags().GetBool("jsonrep")
		eathar.Privileged(kubeconfig, jsonrep)
	},
}

func init() {
	rootCmd.AddCommand(privilegedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// privilegedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// privilegedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
