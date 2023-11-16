/*
Copyright © 2023 Rory McCune <rorym@mccune.org.uk>
*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// hostportsCmd represents the hostports command
var hostportsCmd = &cobra.Command{
	Use:   "hostports",
	Short: "List pods with hostPorts",
	Long: `This will list any pods with hostPorts. This is a security
	risk as hostPorts cannot be controlled by the network policy engine`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		hostports := eathar.HostPorts(options)
		eathar.ReportPSS(hostports, options, "Host Ports")
	},
}

func init() {
	pssCmd.AddCommand(hostportsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostportsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostportsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
