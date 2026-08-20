/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package git

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gerdreiss/mgit/auth"
	"github.com/gerdreiss/mgit/config"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// GitPullWithOptions performs a git pull with custom options
func GitPullWithOptions(repoPath string, checkoutDefault bool, force bool) error {
	repoName := filepath.Base(repoPath)

	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ failed to open repository %s: %v\n", repoName, err)
		return nil
	}

	remotes, err := repo.Remotes()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ failed to pull the latest changes for %s: %v\n", repoName, err)
		return nil
	}
	if len(remotes) == 0 {
		fmt.Fprintf(os.Stderr, "❌ failed to pull the latest changes for %s: the repo seems to be local only\n", repoName)
		return nil
	}

	currentBranchName, err := getCurrentBranchName(repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ failed to get the current branch name for %s: %v\n", repoName, err)
		return nil
	}

	if checkoutDefault {
		defaultBranchName, err := getDefaultBranchName(repo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ failed to get the default branch name for %s: %v\n", repoName, err)
			return nil
		}
		if currentBranchName != defaultBranchName {
			if err := checkoutBranch(repo, defaultBranchName, force); err != nil {
				fmt.Fprintf(os.Stderr, "❌ failed to check out the default branch name of %s: %v\n", repoName, err)
				return nil
			}
			currentBranchName = defaultBranchName
		}
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

	if !status.IsClean() && !force {
		fmt.Fprintf(os.Stderr, "❌ uncommitted changes present in %s (use Force=true to override)\n", repoName)
		return nil
	}

	gitConfig := config.GetGitConfig(remotes[0])
	remoteName := gitConfig.Remote.Name
	authMethod := auth.GetAuthMethod(remotes[0])

	// Prepare pull options
	pullOpts := &git.PullOptions{
		RemoteName:    remoteName,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", currentBranchName)),
		Force:         force,
		Auth:          authMethod,
		Progress:      os.Stdout,
	}

	fmt.Printf("🔄 %s - Pulling %s from %s...\n", repoName, currentBranchName, remoteName)

	// Perform the pull
	err = worktree.Pull(pullOpts)
	if err != nil {
		if err == git.NoErrAlreadyUpToDate {
			fmt.Printf("✅ %s already up-to-date\n", repoName)
			return nil
		}
		fmt.Fprintf(os.Stderr, "❌ failed to pull the latest changes for %s: %v\n", repoName, err)
		return nil
	}

	fmt.Printf("✅ %s pull successful\n", repoName)
	return nil
}
