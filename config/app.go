package config

type AppConfig struct {
	Git struct {
		RemoteName string `mapstructure:"remote-name"`
		AuthMethod string `mapstructure:"auth-method"`
		BasicAuth  struct {
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
		} `mapstructure:"basic-auth"`
		TokenAuth struct {
			Token string `mapstructure:"token"`
		} `mapstructure:"token-auth"`
	} `mapstructure:"git"`
}

var Config AppConfig

func SetAppConfig(config AppConfig) {
	Config = config
}

func GetAppConfig() *AppConfig {
	return &Config
}

func (config AppConfig) HasBasicAuth() bool {
	return config.Git.AuthMethod == "basic"
}

func (config AppConfig) HasTokenAuth() bool {
	return config.Git.AuthMethod == "token"
}
