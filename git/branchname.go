/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package git

import (
	"fmt"

	"github.com/gerdreiss/mgit/auth"
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

	return "", fmt.Errorf("couldn't determine current branch name for repo %v", repo)
}

// getDefaultBranchName determines the default branch of the repository
func getDefaultBranchName(repo *git.Repository) (string, error) {
	remotes, err := repo.Remotes()
	if err != nil || len(remotes) == 0 {
		// No remote configured, try to determine from local branches
		return getDefaultLocalBranchName(repo)
	}

	// Fetch remote info to get HEAD
	opts := git.ListOptions{
		Auth: auth.GetAuthMethod(remotes[0]),
	}
	refs, err := remotes[0].List(&opts)
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
	commonBranches := []string{"main", "master", "develop"}
	for _, branch := range commonBranches {
		ref := plumbing.NewBranchReferenceName(branch)
		if _, err := repo.Reference(ref, true); err == nil {
			return branch, nil
		}
	}

	return "", fmt.Errorf("could not determine default branch")
}

// getLocalBranchNames returns the names of all local braches
func getLocalBranchNames(repo *git.Repository) ([]string, error) {
	var branchNames []string

	branches, err := repo.Branches()
	if err != nil {
		return branchNames, err
	}

	branches.ForEach(func(ref *plumbing.Reference) error {
		branchNames = append(branchNames, ref.Name().Short())
		return nil
	})

	return branchNames, nil
}

// getRemoteBranchNames returns the names of remote branch names
func getRemoteBranchNames(repo *git.Repository) ([]string, error) {
	var branchNames []string

	remotes, err := repo.Remotes()
	if err != nil || len(remotes) == 0 {
		return branchNames, nil
	}

	remoteName := remotes[0].Config().Name

	// List remote references
	refs, err := remotes[0].List(&git.ListOptions{Auth: auth.GetAuthMethod(remotes[0])})
	if err != nil {
		return branchNames, fmt.Errorf("failed to list remote references: %w", err)
	}

	for _, ref := range refs {
		// Filter only branch references (not tags, not HEAD)
		if ref.Name().IsBranch() {
			// Get the branch name without the remote prefix
			// e.g., "refs/remotes/origin/main" -> "main"
			branchNames = append(branchNames, remoteName+"/"+ref.Name().Short())
		}
	}

	return branchNames, nil
}
