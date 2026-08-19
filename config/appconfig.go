package config

import (
	"log"
	"strings"

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

	v.AddConfigPath("/etc/mgit/")
	v.AddConfigPath("/usr/local/etc/mgit/")
	v.AddConfigPath("$HOME/.mgit/")
	v.AddConfigPath(".")

	v.SetDefault("git", []GitConfig{{Remote: &GitRemote{Name: "origin", Host: "github.com"}}})

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config: %v\n", err)
	}

	// Unmarshal the config file into the AppConfig struct
	err = v.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}

func GetAppConfig() *AppConfig {
	return &config
}

func GetGitConfig(remoteUrl string) *GitConfig {
	for _, git := range config.Git {
		if strings.Contains(remoteUrl, git.Remote.Host) {
			return &git
		}
	}
	return nil
}

func (host GitConfig) HasBasicAuth() bool {
	return host.Auth != nil && host.Auth.Basic != nil
}

func (host GitConfig) HasTokenAuth() bool {
	return host.Auth != nil && host.Auth.Token != nil
}
