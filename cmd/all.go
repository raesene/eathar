/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Runs all the checks",
	Long:  `This command runs all the available checks on the target cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		eathar.AllowPrivEsc(options)
		eathar.Apparmor(options)
		eathar.AddedCapabilities(options)
		eathar.DroppedCapabilities(options)
		eathar.Hostipc(options)
		eathar.Hostnet(options)
		eathar.HostPath(options)
		eathar.Hostpid(options)
		eathar.HostPorts(options)
		eathar.HostProcess(options)
		eathar.Privileged(options)
		eathar.Seccomp(options)
		eathar.Procmount(options)
		eathar.Sysctl(options)
	},
}

func init() {
	rootCmd.AddCommand(allCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
