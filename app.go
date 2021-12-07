// Package githubapp provides a convenient interface for handling Github App authentication.
package githubapp

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/go-github/v41/github"
)

// AppsJWTAPI is the interface that is satisfied by the Apps client when authenticated with a JWT.
//
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o fakes/fake_jwt_api.go . AppsJWTAPI
type AppsJWTAPI interface {
	ListInstallations(ctx context.Context, opt *github.ListOptions) ([]*github.Installation, *github.Response, error)
	CreateInstallationToken(ctx context.Context, id int64, opt *github.InstallationTokenOptions) (*github.InstallationToken, *github.Response, error)
}

// AppsTokenAPI is the interface that is satisfied by the Apps client when authenticated with an installation token.
//
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o fakes/fake_token_api.go . AppsTokenAPI
type AppsTokenAPI interface {
	ListRepos(ctx context.Context, opts *github.ListOptions) (*github.ListRepositories, *github.Response, error)
}

// New returns a new App.
func New(client AppsJWTAPI, options ...option) *App {
	a := &App{
		client:         client,
		updateInterval: 1 * time.Minute,
		installsClientFactory: func(token string) AppsTokenAPI {
			return NewInstallationClient(token).V3.Apps
		},
	}
	for _, option := range options {
		option(a)
	}
	return a
}

type option func(*App)

// WithUpdateInterval can be used to override the default update interval for installations and repositories.
func WithUpdateInterval(duration time.Duration) option {
	return func(a *App) {
		a.updateInterval = duration
	}
}

// WithInstallationClientFactory sets the function used to create new installation clients internally, and can be used to inject test fakes.
func WithInstallationClientFactory(f func(token string) AppsTokenAPI) option {
	return func(a *App) {
		a.installsClientFactory = f
	}
}

// App wraps the AppsAPI client and caches the installations and repositories for the installation.
type App struct {
	client                AppsJWTAPI
	installs              []*installation
	installsUpdatedAt     time.Time
	installsClientFactory func(string) AppsTokenAPI
	updateInterval        time.Duration
}

type installation struct {
	ID                    int64
	Owner                 string
	Repositories          []*repository
	RepositoriesUpdatedAt time.Time
}

type repository struct {
	ID   int64
	Name string
}

// Permissions is re-exported to prevent issues with conflicting go-github versions.
type Permissions github.InstallationPermissions

// Token is re-exported to prevent issues with conflicting go-github versions.
type Token struct {
	*github.InstallationToken
}

// CreateInstallationToken returns a new installation token for the given owner, scoped to the provided repositories and permissions.
func (a *App) CreateInstallationToken(owner string, repositories []string, permissions *Permissions) (*Token, error) {
	installationID, err := a.getInstallationID(owner)
	if err != nil {
		return nil, err
	}
	tokenOptions := &github.InstallationTokenOptions{
		Permissions: (*github.InstallationPermissions)(permissions),
	}
	for _, repo := range repositories {
		id, err := a.getRepositoryID(owner, repo)
		if err != nil {
			return nil, err
		}
		tokenOptions.RepositoryIDs = append(tokenOptions.RepositoryIDs, id)
	}
	installationToken, _, err := a.client.CreateInstallationToken(context.TODO(), installationID, tokenOptions)
	if err != nil {
		return nil, err
	}
	return &Token{InstallationToken: installationToken}, nil
}

// getInstallation gets the installation ID for the specified owner.
func (a *App) getInstallationID(owner string) (int64, error) {
	if err := a.updateInstallations(); err != nil {
		return 0, err
	}
	for _, i := range a.installs {
		if i.Owner == owner {
			return i.ID, nil
		}
	}
	return 0, ErrInstallationNotFound(owner)
}

// updateInstallations refreshes the installations on a set interval.
func (a *App) updateInstallations() error {
	if a.installsUpdatedAt.Add(a.updateInterval).After(time.Now()) {
		return nil
	}

	var installs []*installation
	var listOptions = &github.ListOptions{PerPage: 10}

	for {
		list, response, err := a.client.ListInstallations(context.TODO(), listOptions)
		if err != nil {
			return err
		}
		for _, i := range list {
			installs = append(installs, &installation{
				ID:    i.GetID(),
				Owner: strings.ToLower(i.Account.GetLogin()),
			})
		}
		if response.NextPage == 0 {
			break
		}
		listOptions.Page = response.NextPage
	}

	a.installs, a.installsUpdatedAt = installs, time.Now()
	return nil
}

// getInstallation gets the repository ID for the repository.
func (a *App) getRepositoryID(owner, repo string) (int64, error) {
	if err := a.updateRepositories(owner); err != nil {
		return 0, err
	}
	for _, i := range a.installs {
		if i.Owner == owner {
			for _, r := range i.Repositories {
				if r.Name == repo {
					return r.ID, nil
				}
			}
		}
	}

	return 0, ErrInstallationNotFound(fmt.Sprintf("%s/%s", owner, repo))
}

// updateRepositories refreshes the list of repositories for the specified owner on a set interval.
func (a *App) updateRepositories(owner string) error {
	var i *installation
	for _, ii := range a.installs {
		if ii.Owner == owner {
			i = ii
		}
	}

	if i.RepositoriesUpdatedAt.Add(a.updateInterval).After(time.Now()) {
		return nil
	}

	token, err := a.CreateInstallationToken(owner, nil, &Permissions{})
	if err != nil {
		return err
	}

	var (
		repositories []*repository
		listOptions  = &github.ListOptions{PerPage: 100}
		client       = a.installsClientFactory(*token.Token)
	)

	for {
		list, response, err := client.ListRepos(context.TODO(), listOptions)
		if err != nil {
			return err
		}
		for _, r := range list.Repositories {
			repositories = append(repositories, &repository{
				ID:   r.GetID(),
				Name: r.GetName(),
			})
		}
		if response.NextPage == 0 {
			break
		}
		listOptions.Page = response.NextPage
	}

	i.Repositories, i.RepositoriesUpdatedAt = repositories, time.Now()
	return nil
}

// ErrInstallationNotFound is returned if the requested App installation is not found.
type ErrInstallationNotFound string

func (e ErrInstallationNotFound) Error() string {
	return fmt.Sprintf("installation not found: '%s'", string(e))
}

func stringPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
