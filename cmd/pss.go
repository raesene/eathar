/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// pssCmd represents the pss command
var pssCmd = &cobra.Command{
	Use:   "pss",
	Short: "Runs all the PSS checks",
	Long: `This command runs all the available Pod Security checks on the target cluster.
	  You can individual checks by using the subcommands`,
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
	rootCmd.AddCommand(pssCmd)

}
