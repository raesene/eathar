/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// allowprivescCmd represents the allowprivesc command
var allowprivescCmd = &cobra.Command{
	Use:   "allowprivesc",
	Short: "List Pods that allow privilege escalation",
	Long: `This command lists posts that allow privilege escalation.
	This is a default in general for linux container runtimes and allows
	for things like sudo to be used in a container to escalate privileges`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		eathar.AllowPrivEsc(options)
	},
}

func init() {
	pssCmd.AddCommand(allowprivescCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allowprivescCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allowprivescCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
