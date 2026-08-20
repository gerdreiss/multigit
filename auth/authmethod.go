package auth

import (
	"github.com/gerdreiss/mgit/config"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// GetAuthMethod returns the configured auth method
func GetAuthMethod(remote *git.Remote) transport.AuthMethod {
	gitConfig := config.GetGitConfig(remote)
	if gitConfig.HasBasicAuth() {
		return &http.BasicAuth{
			Username: gitConfig.Auth.Basic.Username,
			Password: gitConfig.Auth.Basic.Password,
		}
	}
	if gitConfig.HasTokenAuth() {
		return &http.TokenAuth{
			Token: gitConfig.Auth.Token.Token,
		}
	}
	return nil
}
