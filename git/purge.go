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

	// Get the local branches including the current branch
	localBranches, err := getLocalBranchNames(repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to determine the local branches for %s: %v\n", repoName, err)
		localBranches = []string{}
	}

	remoteBranches, err := getRemoteBranchNames(repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to determine the remote branches of repository %s: %v\n", repoName, err)
		return nil
	}

	for _, localBranch := range slices.DeleteFunc(localBranches, func(lb string) bool { return lb == currentBranch }) {
		if slices.ContainsFunc(remoteBranches, func(rb string) bool { return strings.HasSuffix(rb, "/"+localBranch) }) {
			continue
		}

		if helpers.SureToProceed("⚠️ The branch %s is about to be deleted. Proceed? (N/y) ", localBranch) {
			err = repo.Storer.RemoveReference(plumbing.NewBranchReferenceName(localBranch))
			if err != nil {
				fmt.Fprintf(os.Stderr, "❌ Failed to delete branch %s of repository %s: %v\n", localBranch, repoName, err)
				return nil
			}
		}

		fmt.Printf("✅ branch %s of repository %s deleted successfully\n", localBranch, repoName)
	}

	return nil
}
