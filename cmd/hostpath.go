/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// hostpathCmd represents the hostpath command
var hostpathCmd = &cobra.Command{
	Use:   "hostpath",
	Short: "List pods with hostPath volumes",
	Long: `This will list any pods with hostPath volumes. This is a security
	risk as it allows the container to access the host filesystem`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		hostpath := eathar.HostPath(options)
		eathar.ReportPSS(hostpath, options, "Host Path")
	},
}

func init() {
	pssCmd.AddCommand(hostpathCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostpathCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostpathCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
