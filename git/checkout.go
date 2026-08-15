/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package git

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// CheckoutBranch checks out the branch with the given name
func checkoutBranch(repo *git.Repository, branchName string, force bool) error {
	// Get the worktree
	worktree, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	// Create branch reference
	branchRef := plumbing.NewBranchReferenceName(branchName)

	// Checkout the branch
	err = worktree.Checkout(&git.CheckoutOptions{
		Branch: branchRef,
		Force:  force,
	})

	if err != nil {
		return fmt.Errorf("failed to checkout branch '%s': %w", branchName, err)
	}

	fmt.Printf("✅ Switched to branch '%s'\n", branchName)
	return nil
}
