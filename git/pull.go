/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package git

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// GitPullWithOptions performs a git pull with custom options
func GitPullWithOptions(repoPath string, opts *PullOptions) error {
	repoName := filepath.Base(repoPath)

	// Open the repository
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return fmt.Errorf("%s - failed to open repository: %w", repoName, err)
	}

	currentBranchName, err := getCurrentBranchName(repo)
	if err != nil {
		return fmt.Errorf("%s - failed to get the current branch name: %w", repoName, err)
	}

	if opts.Default {
		defaultBranchName, err := getDefaultBranchName(repo)
		if err != nil {
			return fmt.Errorf("%s - failed to get the default branch name: %w", repoName, err)
		}
		if currentBranchName != defaultBranchName {
			if err := checkoutBranch(repo, defaultBranchName, opts.Force); err != nil {
				return fmt.Errorf("%s - failed to check out the default branch name: %w", repoName, err)
			}
			currentBranchName = defaultBranchName
		}
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

	if !status.IsClean() && !opts.Force {
		return fmt.Errorf("%s - uncommitted changes present in (use Force=true to override)", repoName)
	}

	// Prepare pull options
	pullOpts := &git.PullOptions{
		RemoteName:    opts.RemoteName,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", currentBranchName)),
		Force:         opts.Force,
		Auth:          opts.Auth,
	}

	// Show progress
	if opts.ShowProgress {
		pullOpts.Progress = os.Stdout
	}

	fmt.Printf("🔄 %s - Pulling %s from %s...\n", repoName, currentBranchName, opts.RemoteName)

	// Perform the pull
	err = worktree.Pull(pullOpts)

	if err != nil {
		if err == git.NoErrAlreadyUpToDate {
			fmt.Printf("✅ %s already up-to-date\n", repoName)
			return nil
		}
		return fmt.Errorf("%s failed to pull: %w", repoName, err)
	}

	fmt.Printf("✅ %s pull successful\n", repoName)
	return nil
}
