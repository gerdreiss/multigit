/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/gerdreiss/mgit/exe"
	"github.com/gerdreiss/mgit/helpers"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list  [directory]",
	Short: "Lists GIT reposities with their currently checked out branches in the given directory.",
	Long:  "Lists GIT reposities with their currently checked out branches in the given directory.",
	Args:  cobra.MinimumNArgs(1),
	Run:   listRepos,
}

func listRepos(cmd *cobra.Command, args []string) {
	dirname := helpers.IfElse(len(args) > 0, args[0], ".")

	rootDir, err := filepath.Abs(dirname)
	if err != nil {
		fmt.Printf("❌ Error determining root directory: %v\n", err)
		return
	}

	exe.ListAll(rootDir)
}

func init() {
	rootCmd.AddCommand(listCmd)
}
