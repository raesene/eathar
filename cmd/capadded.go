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
		capadded := eathar.AddedCapabilities(options)
		eathar.ReportPSS(capadded, options, "Added Capabilities")
	},
}

func init() {
	pssCmd.AddCommand(capaddedCmd)

}
