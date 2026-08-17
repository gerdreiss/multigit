package auth

import (
	"github.com/gerdreiss/mgit/config"
	"github.com/gerdreiss/mgit/helpers"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// GetAuth returns the configured auth method
func GetAuth() transport.AuthMethod {
	appConfig := config.GetAppConfig()
	return helpers.IfElse(
		appConfig.Git.AuthMethod == "basic",
		&http.BasicAuth{
			Username: appConfig.Git.BasicAuth.Username,
			Password: appConfig.Git.BasicAuth.Password,
		},
		nil,
	)
}
