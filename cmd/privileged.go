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
		options := cmd.Flags()
		eathar.Privileged(options)
	},
}

func init() {
	rootCmd.AddCommand(privilegedCmd)

}
