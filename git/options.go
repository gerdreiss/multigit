/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package git

import "github.com/go-git/go-git/v5/plumbing/transport"

// CommonOptions configures the git remote and git branch operations
type CommonOptions struct {
	RemoteName string
	Auth       transport.AuthMethod
}

// DefaultCommonOptions returns default list options
func DefaultCommonOptions() *CommonOptions {
	return &CommonOptions{
		RemoteName: "origin",
		Auth:       nil,
	}
}

// PullOptions configures the git pull operation
type PullOptions struct {
	Commons *CommonOptions
	Default bool
	Force   bool
	Rebase  bool
}

// DefaultPullOptions returns default pull options
func DefaultPullOptions() *PullOptions {
	return &PullOptions{
		Commons: DefaultCommonOptions(),
		Default: false,
		Force:   false,
		Rebase:  false,
	}
}

func (opts *CommonOptions) WithRemote(remote string) *CommonOptions {
	opts.RemoteName = remote
	return opts
}

func (opts *CommonOptions) WithAuth(auth transport.AuthMethod) *CommonOptions {
	opts.Auth = auth
	return opts
}

func (opts *PullOptions) WithCommonOptions(commons *CommonOptions) *PullOptions {
	opts.Commons = commons
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
