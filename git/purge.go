package git

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/go-git/go-git/v5"
)

func PurgeLocalBranches(repoPath string, opts *PullOptions) error {
	repoName := filepath.Base(repoPath)
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

	remoteBranches, err := getRemoteBranchNames(repo, opts.RemoteName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to determine the remote branches for %s: %v\n", repoName, err)
		localBranches = []string{}
	}

	for _, localBranch := range slices.DeleteFunc(localBranches, func(b string) bool { return b == currentBranch }) {
		if slices.Contains(remoteBranches, localBranch) {
			continue
		}
		err = repo.DeleteBranch(localBranch)
		if err == nil {
			fmt.Printf("✅ branch %s of repository %s deleted successfully\n", localBranch, repoName)
		} else {
			fmt.Fprintf(os.Stderr, "❌ Failed to delete branch %s of repository %s\n", localBranch, repoName)
		}
	}

	return nil
}
