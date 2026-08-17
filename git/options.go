/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package git

import "github.com/go-git/go-git/v5/plumbing/transport"

// ListOptions configures the git remote and git branch operations
type ListOptions struct {
	RemoteName   string
	Auth         transport.AuthMethod
	ShowProgress bool
}

// DefaultListOptions returns default list options
func DefaultListOptions() *ListOptions {
	return &ListOptions{
		RemoteName:   "origin",
		Auth:         nil,
		ShowProgress: true,
	}
}

// PullOptions configures the git pull operation
type PullOptions struct {
	RemoteName   string
	BranchName   string
	Default      bool
	Force        bool
	Rebase       bool
	Auth         transport.AuthMethod
	ShowProgress bool
}

// DefaultPullOptions returns default pull options
func DefaultPullOptions() *PullOptions {
	return &PullOptions{
		RemoteName:   "origin",
		BranchName:   "",
		Default:      false,
		Force:        false,
		Rebase:       false,
		Auth:         nil,
		ShowProgress: true,
	}
}

func (opts *PullOptions) WithRemote(remote string) *PullOptions {
	opts.RemoteName = remote
	return opts
}

func (opts *PullOptions) WithBranch(branch string) *PullOptions {
	opts.BranchName = branch
	return opts
}

func (opts *PullOptions) WithDefault(def bool) *PullOptions {
	opts.Default = def
	return opts
}

func (opts *PullOptions) WithForce(force bool) *PullOptions {
	opts.Force = force
	return opts
}
