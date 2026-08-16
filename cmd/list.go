/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/gerdreiss/mgit/exe"
	"github.com/gerdreiss/mgit/git"
	"github.com/gerdreiss/mgit/helpers"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list  [directory]",
	Short: "Lists GIT reposities with their currently checked out branches in the given directory.",
	Long:  "Lists GIT reposities with their currently checked out branches in the given directory.",
	Args:  cobra.ExactArgs(1),
	Run:   listRepos,
}

func listRepos(cmd *cobra.Command, args []string) {
	rootDir, err := filepath.Abs(helpers.IfElse(len(args) > 0, args[0], "."))
	if err != nil {
		fmt.Printf("❌ Error determining root directory: %v\n", err)
		return
	}

	branches, err := cmd.Flags().GetBool("branches")
	if err != nil {
		fmt.Printf("❌ 'branches' flag of indeterminate value - using default value: %v\n", err)
		branches = false
	}

	exe.ListAll(rootDir, branches, git.DefaultPullOptions())
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("branches", "b", false, "List all local branches.")
}
