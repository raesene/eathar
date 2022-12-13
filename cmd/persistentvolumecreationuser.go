/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"github.com/raesene/eathar/pkg/eathar"
	"github.com/spf13/cobra"
)

// persistentvolumecreationuserCmd represents the persistentvolumecreationuser command
var persistentvolumecreationuserCmd = &cobra.Command{
	Use:   "persistentvolumecreationuser",
	Short: "Lists users/groups/service accounts with access to create persistent volumes",
	Long:  `This command lists users/groups/service accounts with access to create persistent volumes`,
	Run: func(cmd *cobra.Command, args []string) {
		options := cmd.Flags()
		eathar.CreatePVUsers(options)
	},
}

func init() {
	rbacCmd.AddCommand(persistentvolumecreationuserCmd)

}
