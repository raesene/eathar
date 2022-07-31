/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// capaddedCmd represents the capadded command
var capaddedCmd = &cobra.Command{
	Use:   "capadded",
	Short: "Check for containers with added capabilities",
	Long: `Adding Capabilities to containers over the base set provided
	by the CRI can risk container breakout. This command lists all the
	containers with added capabilities`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		eathar.AddedCapabilities(options)
	},
}

func init() {
	rootCmd.AddCommand(capaddedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// capaddedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// capaddedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
