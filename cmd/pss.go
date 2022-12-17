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
		allowprivesccont := eathar.AllowPrivEsc(options)
		eathar.ReportPSS(allowprivesccont, options, "Allow Privilege Escalation")
		apparmor := eathar.Apparmor(options)
		eathar.ReportPSS(apparmor, options, "Apparmor Disabled")
		capadded := eathar.AddedCapabilities(options)
		eathar.ReportPSS(capadded, options, "Added Capabilities")
		capdropped := eathar.DroppedCapabilities(options)
		eathar.ReportPSS(capdropped, options, "Dropped Capabilities")
		hostipccont := eathar.Hostipc(options)
		eathar.ReportPSS(hostipccont, options, "Host IPC")
		hostnetcont := eathar.Hostnet(options)
		eathar.ReportPSS(hostnetcont, options, "Host Network")
		hostpath := eathar.HostPath(options)
		eathar.ReportPSS(hostpath, options, "Host Path")
		hostpidcont := eathar.Hostpid(options)
		eathar.ReportPSS(hostpidcont, options, "Host PID")
		hostports := eathar.HostPorts(options)
		eathar.ReportPSS(hostports, options, "Host Ports")
		hostprocesscont := eathar.HostProcess(options)
		eathar.ReportPSS(hostprocesscont, options, "Host Process")
		privcont := eathar.Privileged(options)
		eathar.ReportPSS(privcont, options, "Privileged Container")
		seccomp := eathar.Seccomp(options)
		eathar.ReportPSS(seccomp, options, "Seccomp Disabled")
		unmaskedproc := eathar.Procmount(options)
		eathar.ReportPSS(unmaskedproc, options, "Unmasked Procmount")
		sysctls := eathar.Sysctl(options)
		eathar.ReportPSS(sysctls, options, "Unsafe Sysctl")
	},
}

func init() {
	rootCmd.AddCommand(pssCmd)

}
