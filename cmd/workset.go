// Package cmd
/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package cmd

import (
	"github.com/gerdreiss/mgit/exe"
	"github.com/spf13/cobra"
)

// config command (parent of add/remove/delete)
var worksetCmd = &cobra.Command{
	Use:   "workset",
	Short: "Manage worksets",
	Long:  "Commands to manage worksets",
}

// add subcommand
var addCmd = &cobra.Command{
	Use:   "add <workset-id> <path> [path ...]",
	Short: "Add repo(s) to workset",
	Long: `Add repo(s) to a workset.

Examples:
  mgit workset add my-workset .
  mgit workset add my-workset ./this-repo ./that-repo`,

	Args: cobra.MinimumNArgs(2),
	Run:  exe.AddRepoToWorkset,
}

// remove subcommand
var rmCmd = &cobra.Command{
	Use:   "remove <workset-id> <path> [path ...]",
	Short: "Remove repo from workset",
	Long: `Remove one or multiple repos from workset.

Examples:
  mgit workset remove my-workset .
  mgit workset remove my-workset ./this-repo ./that-repo`,

	Args: cobra.MinimumNArgs(2),
	Run:  exe.RemoveRepoFromWorkset,
}

// delete subcommand
var delCmd = &cobra.Command{
	Use:   "delete <workset-id>",
	Short: "Delete a workset",
	Long:  "Remove an entire workset.",
	Args:  cobra.ExactArgs(1),
	Run:   exe.DeleteWorkset,
}

func init() {
	worksetCmd.AddCommand(addCmd, rmCmd, delCmd)
	rootCmd.AddCommand(worksetCmd)
}
