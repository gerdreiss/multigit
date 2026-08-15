/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/gerdreiss/mgit/exe"
	"github.com/gerdreiss/mgit/helpers"
	"github.com/spf13/cobra"
)

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Delete local branches in all git repositories in the given directory that don't have a corresponding remote branch",
	Long:  "Delete all local branches in all git repositories in the given directory that don't have a corresponding remote branch",
	Args:  cobra.ExactArgs(1),
	Run:   purgeLocalBranches,
}

func purgeLocalBranches(cmd *cobra.Command, args []string) {
	rootDir, err := filepath.Abs(helpers.IfElse(len(args) > 0, args[0], "."))
	if err != nil {
		fmt.Printf("❌ Error determining root directory: %v\n", err)
		return
	}

	exe.PurgeLocalBranches(rootDir)
}

func init() {
	rootCmd.AddCommand(purgeCmd)
}
