/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package git

import "github.com/go-git/go-git/v5/plumbing/transport"

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
