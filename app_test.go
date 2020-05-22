package githubapp_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/telia-oss/githubapp"
	"github.com/telia-oss/githubapp/fakes"

	"github.com/google/go-github/v29/github"
)

func isEqual(t *testing.T, expected, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("\nexpected:\n%v\n\ngot:\n%v", expected, got)
	}
}

func noError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestGithubApp(t *testing.T) {
	var (
		client        = &fakes.FakeAppsJWTAPI{}
		tokenClient   = &fakes.FakeAppsTokenAPI{}
		clientFactory = func(string) githubapp.AppsTokenAPI { return tokenClient }
		gh            = githubapp.New(client, githubapp.WithInstallationClientFactory(clientFactory))
		expiresAt     = time.Now().Add(1 * time.Hour)
	)

	client.ListInstallationsReturns([]*github.Installation{{
		ID: github.Int64(23),
		Account: &github.User{
			Login: github.String("owner"),
		},
	}}, &github.Response{}, nil)

	client.CreateInstallationTokenReturns(&github.InstallationToken{
		Token:        github.String("token"),
		ExpiresAt:    &expiresAt,
		Permissions:  nil,
		Repositories: nil,
	}, nil, nil)

	tokenClient.ListReposReturns([]*github.Repository{{
		ID:   github.Int64(23),
		Name: github.String("repository"),
	}}, &github.Response{}, nil)

	token, err := gh.CreateInstallationToken(
		"owner",
		[]string{"repository"},
		&githubapp.Permissions{
			Metadata: github.String("read"),
		})
	noError(t, err)
	isEqual(t, "token", token.GetToken())
	isEqual(t, expiresAt, token.GetExpiresAt())

	_, err = gh.CreateInstallationToken("owner", nil, &githubapp.Permissions{})
	noError(t, err)
	isEqual(t, 1, client.ListInstallationsCallCount())
	isEqual(t, 3, client.CreateInstallationTokenCallCount())
	isEqual(t, 1, tokenClient.ListReposCallCount())
}
