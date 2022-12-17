/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// imagelistCmd represents the imagelist command
var imagelistCmd = &cobra.Command{
	Use:   "imagelist",
	Short: "List images used in the cluster",
	Long:  `This will provide a list of images used in the cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		imageListSlice := eathar.ImageList(options)
		eathar.ReportImage(imageListSlice, options, "Image List")
	},
}

func init() {
	infoCmd.AddCommand(imagelistCmd)
}
