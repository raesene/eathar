/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// hostpidCmd represents the hostpid command
var hostpidCmd = &cobra.Command{
	Use:   "hostpid",
	Short: "List pods with host PID access",
	Long: `This command lists pods which have host PID access
	This access could be misused by an attacker to affect processes
	in other containers on running on the host.`,
	Run: func(cmd *cobra.Command, args []string) {
		//options := make(map[string]interface{})
		options := cmd.Flags()
		//Get the value of the kubeconfig flag so we can pass it to the command
		//options["kubeconfig"], _ = cmd.Flags().GetString("kubeconfig")
		//options["jsonrep"], _ = cmd.Flags().GetBool("jsonrep")
		eathar.Hostpid(options)
	},
}

func init() {
	pssCmd.AddCommand(hostpidCmd)

}
