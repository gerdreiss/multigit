package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("$HOME/.mgit/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			// ignore
		default:
			log.Fatalf("Unable to read config: %v\n", err)
		}
	}

	// Unmarshal the config file into the AppConfig struct
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v\n", err)
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

func Get(key string, displayJson bool) (string, error) {
	key = strings.TrimSpace(key)
	if len(key) < 5 {
		return "", fmt.Errorf("invalid key. It should start with at least 'git.N' where N is a valid index optionally followed by either '.remote.' or '.auth.'")
	}
	path := strings.Split(key, ".")
	if len(path) < 2 {
		return "", fmt.Errorf("invalid key. It should start with at least 'git.N' where N is a valid index optionally followed by either '.remote.' or '.auth.'")
	}
	idx, err := strconv.Atoi(path[1])
	if err != nil {
		return "", fmt.Errorf("invalid key. It should start with git.N where N is a valid index")
	}
	if idx > len(config.Git) {
		return "", fmt.Errorf("invalid key. The index can be maximal %d", len(config.Git))
	}

	yamlbytes, err := yaml.Marshal(config)
	if err != nil {
		return "", err
	}

	value, err := helpers.GetValue(string(yamlbytes), key)
	if err != nil {
		return "", err
	}

	if displayJson {
		prettyJSON, err := json.MarshalIndent(value, "", "  ")
		if err != nil {
			return "", fmt.Errorf("error marshaling config: %v", err)
		}
		return string(prettyJSON), nil
	}

	yamlData, err := yaml.Marshal(value)
	if err != nil {
		return "", fmt.Errorf("error marshaling config to YAML: %v", err)
	}
	return string(yamlData), nil
}

func Set(key string, value string) error {
	key = strings.TrimSpace(key)
	if len(key) < 5 {
		return fmt.Errorf("invalid key. It should start with at least 'git.N' where N is a valid index optionally followed by either '.remote.' or '.auth.'")
	}
	path := strings.Split(key, ".")
	if len(path) < 2 {
		return fmt.Errorf("invalid key. It should start with at least 'git.N' where N is a valid index optionally followed by either '.remote.' or '.auth.'")
	}
	idx, err := strconv.Atoi(path[1])
	if err != nil {
		return fmt.Errorf("invalid key. It should start with git.N where N is a valid index")
	}
	if idx > len(config.Git) {
		return fmt.Errorf("invalid key. The index can be maximal %d", len(config.Git))
	}

	yamlstring, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	updated, err := helpers.SetValue(string(yamlstring), key, value)
	if err != nil {
		return err
	}

	yaml.Unmarshal([]byte(updated), &config)

	file, err := os.OpenFile(viper.ConfigFileUsed(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	viper.Set("git", config.Git)
	viper.WriteConfigTo(file)

	return nil
}
