// Package exe
/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package exe

import (
	"fmt"

	"github.com/spf13/cobra"
)

func AddRepoToWorkset(cmd *cobra.Command, args []string) {
	fmt.Println("Add repo to workset called...")
}

func RemoveRepoFromWorkset(cmd *cobra.Command, args []string) {
	fmt.Println("Remove repo from workset called...")
}

func DeleteWorkset(cmd *cobra.Command, args []string) {
	fmt.Println("Delete workset command called...")
}
