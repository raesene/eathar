/*
Copyright © 2023 Rory McCune <rorym@mccune.org.uk>
*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// allInfoCmd represents the allInfo command
var allInfoCmd = &cobra.Command{
	Use:   "all",
	Short: "Runs all the info checks",
	Long:  `Runs all the checks in the info group.`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		imageListSlice := eathar.ImageList(options)
		eathar.ReportImage(imageListSlice, options, "Image List")
	},
}

func init() {
	infoCmd.AddCommand(allInfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
