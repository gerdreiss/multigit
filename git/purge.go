/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package git

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/gerdreiss/mgit/config"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func PurgeLocalBranches(repoPath string, exclude []string) error {
	repoName := filepath.Base(repoPath)
	if slices.Contains(exclude, repoName) {
		return nil
	}

	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ failed to open repo %s: %v\n", repoName, err)
		return nil
	}

	// Get current currentBranch name
	currentBranch, err := getCurrentBranchName(repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to determine the current branch name for %s: %v\n", repoName, err)
		return nil
	}

	localBranches, err := getLocalBranchNames(repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to determine the local branches for %s: %v\n", repoName, err)
		localBranches = []string{}
	}

	remoteBranches := []string{}
	remotes, err := repo.Remotes()
	if err == nil {
		gitConfig := config.GetGitConfig(remotes[0])
		remoteBranches, _ = getRemoteBranchNames(repo, gitConfig.Remote.Name)
	}

	for _, localBranch := range slices.DeleteFunc(localBranches, func(b string) bool { return b == currentBranch }) {
		if slices.Contains(remoteBranches, localBranch) {
			continue
		}

		branchRef := plumbing.NewBranchReferenceName(localBranch)
		err = repo.Storer.RemoveReference(branchRef)
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Failed to delete branch %s of repository %s: %v\n", localBranch, repoName, err)
			return nil
		}

		fmt.Printf("✅ branch %s of repository %s deleted successfully\n", localBranch, repoName)
	}

	return nil
}
