/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package git

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// getCurrentBranchName determines the current branch of the repository
func getCurrentBranchName(repo *git.Repository) (string, error) {
	// Get current HEAD
	head, err := repo.Head()
	if err != nil {
		return "", fmt.Errorf("failed to get HEAD: %w", err)
	}

	if head.Name().IsBranch() {
		return head.Name().Short(), nil
	}

	return "", fmt.Errorf("Couldn't determine current branch name for repo %v", repo)
}

// getDefaultBranchName determines the default branch of the repository
func getDefaultBranchName(repo *git.Repository) (string, error) {
	// Try to get the default branch from remote
	remote, err := repo.Remote("origin")
	if err != nil {
		// No remote configured, try to determine from local branches
		return getDefaultLocalBranchName(repo)
	}

	// Fetch remote info to get HEAD
	refs, err := remote.List(&git.ListOptions{})
	if err != nil {
		// If can't list remote, try local
		return getDefaultLocalBranchName(repo)
	}

	// Find the symbolic reference for HEAD
	for _, ref := range refs {
		if ref.Name() == plumbing.HEAD {
			// HEAD points to the default branch
			target := ref.Target()
			if target.IsBranch() {
				return target.Short(), nil
			}
		}
	}

	// If all else fails, try local
	return getDefaultLocalBranchName(repo)
}

// getDefaultLocalBranchName tries to determine default branch from local repository
func getDefaultLocalBranchName(repo *git.Repository) (string, error) {
	head, err := repo.Head()
	if err != nil {
		return "", fmt.Errorf("failed to get HEAD: %w", err)
	}

	if head.Name().IsBranch() {
		return head.Name().Short(), nil
	}

	// Try common branch names
	commonBranches := []string{"main", "master", "develop", "dev"}
	for _, branch := range commonBranches {
		ref := plumbing.NewBranchReferenceName(branch)
		if _, err := repo.Reference(ref, true); err == nil {
			return branch, nil
		}
	}

	return "", fmt.Errorf("could not determine default branch")
}
