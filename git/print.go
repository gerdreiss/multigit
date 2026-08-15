/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package git

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gerdreiss/mgit/helpers"
	"github.com/go-git/go-git/v5"
)

const red = "\033[31m"
const green = "\033[32m"
const reset = "\033[0m"

func PrintRepoWithBranchName(repoPath string, listLocalBranches bool) error {
	repoName := filepath.Base(repoPath)
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ failed to open repo %s: %v\n", repoName, err)
		return nil
	}

	// Get the worktree
	worktree, err := repo.Worktree()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ failed to get worktree of %s: %v\n", repoName, err)
		return nil
	}

	// Check for uncommitted changes
	status, err := worktree.Status()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ failed to get status of %s: %v\n", repoName, err)
		return nil
	}

	// Get current branch name
	branch, err := getCurrentBranchName(repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to determine the current branch name for %s: %v\n", repoName, err)
		return nil
	}

	// Print out all the info
	fmt.Printf(
		"%s %s[%s%s%s%s]%s\n",
		repoPath,
		red,
		green,
		branch,
		red,
		helpers.IfElse(status.IsClean(), "", "*"),
		reset,
	)

	var branches []string
	if listLocalBranches {
		branches, err = getLocalBranchNames(repo, branch)
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Failed to determine the local branches for %s: %v\n", repoName, err)
			branches = []string{}
		}
		if len(branches) > 0 {
			spaces := strings.Repeat(" ", len(repoPath)-3)
			for _, branch := range branches {
				fmt.Printf(
					"%s -- %s[%s%s%s]%s\n",
					spaces,
					red,
					green,
					branch,
					red,
					reset,
				)
			}
		}
	}

	return nil
}
