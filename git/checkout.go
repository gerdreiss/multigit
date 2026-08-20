/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package git

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// CheckoutBranch checks out the branch with the given name
func checkoutBranch(repo *git.Repository, branchName string, force bool) error {
	// Get the worktree
	worktree, err := repo.Worktree()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ failed to get worktree for %v: %v\n", repo, err)
		return nil
	}

	// Create branch reference
	branchRef := plumbing.NewBranchReferenceName(branchName)

	// Checkout the branch
	opts := git.CheckoutOptions{
		Branch: branchRef,
		Force:  force,
	}
	err = worktree.Checkout(&opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ failed to checkout branch '%s': %v\n", branchName, err)
		return nil
	}

	fmt.Printf("✅ Switched to branch '%s'\n", branchName)
	return nil
}
