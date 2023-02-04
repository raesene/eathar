/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// allPSSCmd represents the allPSS command
var allPSSCmd = &cobra.Command{
	Use:   "all",
	Short: "Runs all the PSS commands",
	Long:  `Runs all the checks in the PSS group.`,
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
	pssCmd.AddCommand(allPSSCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allPSSCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allPSSCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
