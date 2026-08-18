/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

	exclude, err := cmd.Flags().GetStringSlice("exclude")
	if err != nil {
		fmt.Println("⚠️ 'exclude' flag of indeterminate value - using default value")
		exclude = []string{}
	}

	if sureToProceed() {
		exe.PurgeLocalBranches(rootDir, exclude)
	}
}

func init() {
	rootCmd.AddCommand(purgeCmd)

	purgeCmd.Flags().StringSliceP("exclude", "x", []string{}, "Exclude the repositories with the names given here from purging.")
}

func sureToProceed() bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("⚠️ Are you sure to proceed? Have you checked with `mgit list -b`? (N/y) ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error reading input: %s\n", err)
		return false
	}

	response := strings.ToLower(strings.TrimSpace(input))

	return response == "y"
}
