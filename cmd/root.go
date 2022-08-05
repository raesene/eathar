/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "eathar",
	Short: "Kubernetes Security Information Retriever",
	Long: `Eathar is a program designed to pull information that might be
	of interest back from Kubernetes clusters.`,
	Version: "0.1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Get the home directory for default kubeconfig location
	homedir, err := os.UserHomeDir()
	if err != nil {
		os.Exit(1)
	}
	// Set the default path for kubeconfig files
	defaultkubeconfig := homedir + "/.kube/config"
	rootCmd.PersistentFlags().StringP("kubeconfig", "k", defaultkubeconfig,
		"Kubeconfig file")
	rootCmd.PersistentFlags().BoolP("jsonrep", "j", false, "json reporting")
	rootCmd.PersistentFlags().StringP("file", "f", "", "Report file")
}
