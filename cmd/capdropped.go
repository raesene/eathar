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
	Short: "List pods and containers that drop capabilities",
	Long: `This will list containers and pods which drop capabilities.
	this is a good hardening measure to ensure that containers run with
	least privilege`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		capdropped := eathar.DroppedCapabilities(options)
		eathar.ReportPSS(capdropped, options, "Dropped Capabilities")
	},
}

func init() {
	pssCmd.AddCommand(capdroppedCmd)

}
