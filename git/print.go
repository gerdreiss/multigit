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

const (
	red   = "\033[31m"
	green = "\033[32m"
	blue  = "\033[34m"
	grey  = "\033[90m"
	reset = "\033[0m"
)

func PrintRepoWithBranchName(repoPath string, listLocalBranches bool, opts *CommonOptions) error {
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

	remoteBranches, err := getRemoteBranchNames(repo, opts.RemoteName)
	if err != nil {
		remoteBranches = []string{}
	}

	isCurrentBranchClean := status.IsClean()
	isCurrentBranchTracked := slices.Contains(remoteBranches, currentBranch)

	// Print out all the info
	fmt.Printf(
		"%s %s[%s%s%s%s][%s%s%s]%s\n",
		repoPath,
		red,
		green,
		currentBranch,
		red,
		helpers.IfElse(isCurrentBranchClean, "", "*"),
		helpers.IfElse(isCurrentBranchTracked, blue, grey),
		helpers.IfElse(isCurrentBranchTracked, opts.RemoteName+"/"+currentBranch, "untracked"),
		red,
		reset,
	)

	lineLength := len(repoPath) + len(currentBranch) + helpers.IfElse(isCurrentBranchClean, 1, 2)

	if listLocalBranches {
		localBranches, err := getLocalBranchNames(repo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Failed to determine the local branches for %s: %v\n", repoName, err)
			localBranches = []string{}
		}

		localBranches = slices.DeleteFunc(localBranches, func(b string) bool { return b == currentBranch })
		if len(localBranches) > 0 {
			for _, localBranch := range localBranches {
				spaces := strings.Repeat(" ", lineLength-len(localBranch))
				isLocalBranchTracked := slices.Contains(remoteBranches, localBranch)

				fmt.Printf(
					"%s%s[%s%s%s][%s%s%s]%s\n",
					spaces,
					red,
					green,
					localBranch,
					red,
					helpers.IfElse(isLocalBranchTracked, blue, grey),
					helpers.IfElse(isLocalBranchTracked, opts.RemoteName+"/"+localBranch, "untracked"),
					red,
					reset,
				)
			}
		}
	}

	return nil
}
