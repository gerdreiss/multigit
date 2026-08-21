/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/gerdreiss/mgit/exe"
	"github.com/spf13/cobra"
)

// config command (parent of read/write)
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Long:  "Commands to manage application configuration",
}

// config read subcommand
var configReadCmd = &cobra.Command{
	Use:   "read [name]",
	Short: "Read configuration",
	Long: `Read and display the current configuration values.
	When calling without arguments, the entire configuration is displayed. 
	You can display a subselection of the configuration by providing the names of configuration values.

Examples:
  mgit config read
  mgit config read git.1.remote.name
  mgit config read git.1.remote.name git.2.auth.token.token`,

	Run: exe.PrintConfig,
}

// config write subcommand (optional)
var configWriteCmd = &cobra.Command{
	Use:   "write <name=value> [name=value ...]",
	Short: "Write configuration values",
	Long: `Write one or more configuration values in the format name=value.

Examples:
  mgit config write git.1.remote.name=origin
  mgit config write git.2.auth.token.token=xyz`,

	Args: validArgs,
	Run:  exe.WriteConfig,
}

func validArgs(cmd *cobra.Command, args []string) error {
	// Ensure at least one argument is provided
	if len(args) < 1 {
		return fmt.Errorf("at least one name=value argument is required")
	}

	// Validate all arguments are in name=value format
	for _, arg := range args {
		if !strings.Contains(arg, "=") {
			return fmt.Errorf("invalid format: '%s' (expected name=value)", arg)
		}
		// Optional: validate that both name and value are non-empty
		parts := strings.SplitN(arg, "=", 2)
		if parts[0] == "" {
			return fmt.Errorf("empty key in argument: '%s'", arg)
		}
	}
	return nil
}

func init() {
	configCmd.AddCommand(configReadCmd, configWriteCmd)
	rootCmd.AddCommand(configCmd)

	configReadCmd.Flags().BoolP("json", "j", false, "Print configuration as JSON.")
}
