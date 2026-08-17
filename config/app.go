package config

type AppConfig struct {
	Git struct {
		RemoteName string `mapstructure:"remote-name"`
		AuthMethod string `mapstructure:"auth-method"`
		BasicAuth  struct {
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
		} `mapstructure:"basic-auth"`
		Ssl struct {
			PrivateKey string `mapstring:"private-key"`
		}
	} `mapstructure:"git"`
}

var Config AppConfig

func SetConfig(config AppConfig) {
	Config = config
}

func GetConfig() *AppConfig {
	return &Config
}
