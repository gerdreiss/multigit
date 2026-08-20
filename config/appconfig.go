package config

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/cybergodev/json"
	"github.com/gerdreiss/mgit/helpers"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/viper"
	"go.yaml.in/yaml/v3"
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

func Get(key string) (string, error) {
	re := regexp.MustCompile(`^git\.(\d+)(\.\w+)+$`)
	if re.MatchString(key) {
		marshalled, err := json.Marshal(config)
		if err != nil {
			return "", err
		}
		jsonstring := string(marshalled)
		fmt.Println(jsonstring)
		value, err := json.Get(jsonstring, key)
		if err != nil {
			return "", err
		}
		if s, ok := value.(string); ok {
			return s, nil
		}
		if m, ok := value.(map[string]any); ok {
			data, err := yaml.Marshal(m)
			if err != nil {
				return "", err
			}
			return string(data), nil
		}
	}

	return "", errors.New("Invalid configuration key: " + key)
}

func Set(key string, value string) error {
	re := regexp.MustCompile(`^git\.(\d+)(\.\w+)+$`)
	if re.MatchString(key) {
		marshalled, err := json.Marshal(config)
		if err != nil {
			return err
		}
		jsonstring := string(marshalled)
		updated, err := json.Set(jsonstring, key, value)
		if err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(updated), &config); err != nil {
			return err
		}
	}

	return errors.New("Invalid configuration key: " + key)
}
