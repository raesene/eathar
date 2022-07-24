/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// hostpidCmd represents the hostpid command
var hostpidCmd = &cobra.Command{
	Use:   "hostpid",
	Short: "List pods with host PID access",
	Long: `This command lists pods which have host PID access
	This access could be misused by an attacker to affect processes
	in other containers on running on the host.`,
	Run: func(cmd *cobra.Command, args []string) {
		kubeconfig, _ := cmd.Flags().GetString("kubeconfig")
		eathar.Hostpid(kubeconfig)
	},
}

func init() {
	rootCmd.AddCommand(hostpidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostpidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostpidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
