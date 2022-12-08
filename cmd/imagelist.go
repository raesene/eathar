/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
		eathar.ImageList(options)
	},
}

func init() {
	rootCmd.AddCommand(imagelistCmd)
}
