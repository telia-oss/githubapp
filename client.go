package githubapp

import (
	"context"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v45/github"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// NewClient returns a client for the Github V3 (REST) AppsAPI authenticated with a private key.
func NewClient(integrationID int64, privateKey []byte) (AppsJWTAPI, error) {
	transport, err := ghinstallation.NewAppsTransport(http.DefaultTransport, integrationID, privateKey)
	if err != nil {
		return nil, err
	}
	client := github.NewClient(&http.Client{
		Transport: transport,
	})
	return &struct {
		*github.Client
		*github.AppsService
	}{
		client,
		client.Apps,
	}, nil
}

// NewInstallationClient returns a new client.
func NewInstallationClient(token string) *InstallationClient {
	client := oauth2.NewClient(context.TODO(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))
	return &InstallationClient{V3: github.NewClient(client), V4: githubv4.NewClient(client)}
}

// InstallationClient is authenticated with an installation token and includes a client for both the V3 and V4 Github APIs.
type InstallationClient struct {
	V3 *github.Client
	V4 *githubv4.Client
}
