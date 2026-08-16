/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package git

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/gerdreiss/mgit/helpers"
	"github.com/go-git/go-git/v5"
)

const red = "\033[31m"
const green = "\033[32m"
const blue = "\033[34m"
const grey = "\033[90m"
const reset = "\033[0m"

func PrintRepoWithBranchName(repoPath string, listLocalBranches bool, remoteName string) error {
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

	// Get current currentBranch name
	currentBranch, err := getCurrentBranchName(repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to determine the current branch name for %s: %v\n", repoName, err)
		return nil
	}

	remoteBranches, err := getRemoteBranchNames(repo, remoteName)
	if err != nil {
		remoteBranches = []string{}
	}

	// Print out all the info
	fmt.Printf(
		"%s %s[%s%s%s%s][%s%s%s]%s\n",
		repoPath,
		red,
		green,
		currentBranch,
		red,
		helpers.IfElse(status.IsClean(), "", "*"),
		helpers.IfElse(slices.Contains(remoteBranches, currentBranch), blue, grey),
		helpers.IfElse(slices.Contains(remoteBranches, currentBranch), remoteName+"/"+currentBranch, "untracked"),
		red,
		reset,
	)

	if listLocalBranches {
		localBranches, err := getLocalBranchNames(repo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Failed to determine the local branches for %s: %v\n", repoName, err)
			localBranches = []string{}
		}
		localBranches = slices.DeleteFunc(localBranches, func(b string) bool { return b == currentBranch })
		if len(localBranches) > 0 {
			for _, localBranch := range localBranches {
				spaces := strings.Repeat(" ", len(repoPath)-len(localBranch)+len(currentBranch)+helpers.IfElse(status.IsClean(), 1, 2))
				fmt.Printf(
					"%s%s[%s%s%s][%s%s%s]%s\n",
					spaces,
					red,
					green,
					localBranch,
					red,
					helpers.IfElse(slices.Contains(remoteBranches, localBranch), blue, grey),
					helpers.IfElse(slices.Contains(remoteBranches, localBranch), remoteName+"/"+localBranch, "untracked"),
					red,
					reset,
				)
			}
		}
	}

	return nil
}
