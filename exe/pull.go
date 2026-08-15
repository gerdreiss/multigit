/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package exe

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/gerdreiss/mgit/git"
)

func PullAll(rootDir string, checkoutDefault bool, force bool, exclude []string) {
	err := filepath.Walk(rootDir, func(repoPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		if slices.Contains(exclude, filepath.Base(repoPath)) {
			return nil
		}

		i, e := os.Stat(filepath.Join(repoPath, ".git"))
		if e != nil {
			return nil
		}

		if i.IsDir() {
			opts := git.DefaultPullOptions().WithDefault(checkoutDefault).WithForce(force)
			return git.GitPullWithOptions(repoPath, opts)
		} else {
			return nil
		}
	})

	if err != nil {
		fmt.Printf("❌ Error pulling GIT repositories in %q: %v\n", rootDir, err)
	}
}
