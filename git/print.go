/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package git

import (
	"fmt"
	"path/filepath"

	"github.com/gerdreiss/mgit/helpers"
	"github.com/go-git/go-git/v5"
)

const red = "\033[31m"
const green = "\033[32m"
const reset = "\033[0m"

func PrintRepoWithBranchName(repoPath string) error {
	repoName := filepath.Base(repoPath)
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return fmt.Errorf("failed to open repo: %w", err)
	}

	// Get the worktree
	worktree, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("%s - failed to get worktree: %w", repoName, err)
	}

	// Check for uncommitted changes
	status, err := worktree.Status()
	if err != nil {
		return fmt.Errorf("%s - failed to get status: %w", repoName, err)
	}

	branch, err := getCurrentBranchName(repo)
	if err != nil {
		return fmt.Errorf("Failed to determine the current branch name for %s: %v", repoName, err)
	}

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
	return nil
}
