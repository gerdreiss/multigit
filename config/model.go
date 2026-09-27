// Package config
/*
Copyright © 2026 Gerd Reiss gerd@reiss.pro
*/
package config

type TokenAuthConfig struct {
	Token string `mapstructure:"token" yaml:"token,omitempty" json:"token,omitempty"`
}

type BasicAuthConfig struct {
	Username string `mapstructure:"username" yaml:"username,omitempty" json:"username,omitempty"`
	Password string `mapstructure:"password" yaml:"password,omitempty" json:"password,omitempty"`
}

type GitAuth struct {
	// Exactly one of the following blocks will be meaningful based on AuthMethod:
	Basic *BasicAuthConfig `mapstructure:"basic" yaml:"basic,omitempty" json:"basic,omitempty"`
	Token *TokenAuthConfig `mapstructure:"token" yaml:"token,omitempty" json:"token,omitempty"`
}

type GitRemote struct {
	Name string `mapstructure:"name" yaml:"name,omitempty" json:"name,omitempty"`
	Host string `mapstructure:"host" yaml:"host,omitempty" json:"host,omitempty"`
}

type GitWorkset struct {
	Id     string `mapstructure:"id" yaml:"id"`
	Branch string `mapstructure:"branch" yaml:"branch"`
}

type GitConfig struct {
	Remote  *GitRemote  `mapstructure:"remote" yaml:"remote,omitempty" json:"remote,omitempty"`
	Auth    *GitAuth    `mapstructure:"auth" yaml:"auth,omitempty" json:"auth,omitempty"`
	Workset *GitWorkset `mapstructure:"workset" yaml:"workset,omitempty" json:"workset,omitempty"`
}

type AppConfig struct {
	Git []GitConfig `mapstructure:"git" yaml:"git,omitempty" json:"git,omitempty"`
}
