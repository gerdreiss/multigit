/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package exe

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gerdreiss/mgit/git"
)

func PurgeLocalBranches(rootDir string, exclude []string) {
	err := filepath.Walk(rootDir, func(repoPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() || strings.Contains(repoPath, "/.git") {
			return nil
		}

		i, e := os.Stat(filepath.Join(repoPath, ".git"))
		if e != nil {
			return nil
		}

		if i.IsDir() {
			fmt.Printf("Checking repository: %s\n", repoPath)
			return git.PurgeLocalBranches(repoPath, exclude)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("❌ Error purging local branches in GIT repositories in %q, %v\n", rootDir, err)
	}
}
