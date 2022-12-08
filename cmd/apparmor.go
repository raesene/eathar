/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// apparmorCmd represents the apparmor command
var apparmorCmd = &cobra.Command{
	Use:   "apparmor",
	Short: "List pods without apparmor profiles",
	Long: `This command will list pods that do not have apparmor profiles
	assigned to them. Apparmor is part of the layers of isolation which should be
	applied to all containers`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		eathar.Apparmor(options)
	},
}

func init() {
	rootCmd.AddCommand(apparmorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apparmorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apparmorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
