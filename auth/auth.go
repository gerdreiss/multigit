package auth

import (
	"github.com/gerdreiss/mgit/config"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// GetAuthMethod returns the configured auth method
func GetAuthMethod() transport.AuthMethod {
	appConfig := config.GetAppConfig()
	if appConfig.HasBasicAuth() {
		return &http.BasicAuth{
			Username: appConfig.Git.BasicAuth.Username,
			Password: appConfig.Git.BasicAuth.Password,
		}
	}
	if appConfig.HasTokenAuth() {
		return &http.TokenAuth{
			Token: appConfig.Git.TokenAuth.Token,
		}
	}
	return nil
}
