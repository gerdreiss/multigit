/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package exe

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gerdreiss/mgit/git"
)

func ListAll(rootDir string, branches bool, opts *git.PullOptions) {
	err := filepath.Walk(rootDir, func(repoPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}

		i, e := os.Stat(filepath.Join(repoPath, ".git"))
		if e != nil {
			return nil
		}

		if i.IsDir() {
			return git.PrintRepoWithBranchName(repoPath, branches, opts.RemoteName)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("❌ Error listing GIT repositories in %q, %v\n", rootDir, err)
	}
}
