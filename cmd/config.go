/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package cmd

import (
	"github.com/gerdreiss/mgit/exe"
	"github.com/spf13/cobra"
)

// config command (parent of read/write)
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Read configuration",
	Long:  "Read and display the current configuration values",

	Run: exe.PrintConfig,
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().BoolP("json", "j", false, "Print configuration as JSON.")
}
