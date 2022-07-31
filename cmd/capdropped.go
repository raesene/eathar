/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// capdroppedCmd represents the capdropped command
var capdroppedCmd = &cobra.Command{
	Use:   "capdropped",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		eathar.DroppedCapabilities(options)
	},
}

func init() {
	rootCmd.AddCommand(capdroppedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// capdroppedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// capdroppedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
