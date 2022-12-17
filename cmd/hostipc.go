/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// hostipcCmd represents the hostipc command
var hostipcCmd = &cobra.Command{
	Use:   "hostipc",
	Short: "Show containers with hostIPC",
	Long: `Shows containers set to use the host's
	IPC namespace`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		hostipccont := eathar.Hostipc(options)
		eathar.ReportPSS(hostipccont, options, "Host IPC")
	},
}

func init() {
	pssCmd.AddCommand(hostipcCmd)

}
