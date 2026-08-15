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

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Executes git pull in all git repositories in the given directory",
	Long:  "Executes git pull in all git repositories in the given directory",
	Run:   PullAll,
}

func PullAll(cmd *cobra.Command, args []string) {
	rootDir, err := filepath.Abs(helpers.IfElse(len(args) > 0, args[0], "."))
	if err != nil {
		fmt.Printf("❌ Error determining root directory: %v\n", err)
		return
	}

	checkoutDefault, err := cmd.Flags().GetBool("default")
	if err != nil {
		fmt.Println("⚠️ 'default' flag of indeterminate value - using default value")
		checkoutDefault = false
	}

	force, err := cmd.Flags().GetBool("force")
	if err != nil {
		fmt.Println("⚠️ 'force' flag of indeterminate value - using default value")
		force = false
	}

	exclude, err := cmd.Flags().GetStringSlice("exclude")
	if err != nil {
		fmt.Println("⚠️ 'exclude' flag of indeterminate value - using default value")
		exclude = []string{}
	}

	exe.PullAll(rootDir, checkoutDefault, force, exclude)
}

func init() {
	rootCmd.AddCommand(pullCmd)

	pullCmd.Flags().BoolP("force", "f", false, "Force pulling or checking out the default branch when current branch is dirty.")
	pullCmd.Flags().BoolP("default", "d", false, "Check out the default branch (e.g. main) before pulling the latest changes.")
	pullCmd.Flags().StringSliceP("exclude", "x", []string{}, "Exclude the repositories with the names given here from pulling.")
}
