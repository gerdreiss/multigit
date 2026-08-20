package config

import (
	"log"

	"github.com/gerdreiss/mgit/helpers"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/viper"
)

type TokenAuthConfig struct {
	Token string `mapstructure:"token"`
}

type BasicAuthConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type GitAuth struct {
	// Exactly one of the following blocks will be meaningful based on AuthMethod:
	Basic *BasicAuthConfig `mapstructure:"basic"`
	Token *TokenAuthConfig `mapstructure:"token"`
}

type GitRemote struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
}

type GitConfig struct {
	Remote *GitRemote `mapstructure:"remote"`
	Auth   *GitAuth   `mapstructure:"auth"`
}

type AppConfig struct {
	Git []GitConfig `mapstructure:"git"`
}

var config AppConfig

func Load() {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AddConfigPath("$HOME/.mgit/")
	v.AddConfigPath(".")

	_ = v.ReadInConfig()

	// Unmarshal the config file into the AppConfig struct
	err := v.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}

func GetAppConfig() *AppConfig {
	return &config
}

// GetGitConfig return a configuration for the given Remote even if none exists
func GetGitConfig(remote *git.Remote) *GitConfig {
	remoteConfig := remote.Config()

	remoteName := remoteConfig.Name
	// ignore error here because we rely on the the URL being valid here
	hostName, _ := helpers.GetHostName(remoteConfig.URLs[0])

	for _, git := range config.Git {
		if git.Remote.Name == remoteName && git.Remote.Host == hostName {
			return &git
		}
	}

	// if no corresponding configuration was found for the given Remote,
	// a new one is returned without the Auth
	return &GitConfig{Remote: &GitRemote{Name: remoteConfig.Name, Host: hostName}}
}

func (host GitConfig) HasBasicAuth() bool {
	return host.Auth != nil && host.Auth.Basic != nil
}

func (host GitConfig) HasTokenAuth() bool {
	return host.Auth != nil && host.Auth.Token != nil
}
