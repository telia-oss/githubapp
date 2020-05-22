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
		client    = &fakes.FakeAppsJWTAPI{}
		gh        = githubapp.New(client)
		expiresAt = time.Now().Add(1 * time.Hour)
	)

	client.ListInstallationsReturns([]*github.Installation{{
		ID: github.Int64(23),
		Account: &github.User{
			Login: github.String("login"),
		},
	}}, &github.Response{NextPage: 0}, nil)

	client.CreateInstallationTokenReturns(&github.InstallationToken{
		Token:        github.String("token"),
		ExpiresAt:    &expiresAt,
		Permissions:  nil,
		Repositories: nil,
	}, nil, nil)

	token, err := gh.CreateInstallationToken("login", nil, &githubapp.Permissions{})
	noError(t, err)
	isEqual(t, "token", *token.Token)
	isEqual(t, expiresAt, *token.ExpiresAt)

	_, err = gh.CreateInstallationToken("login", nil, &githubapp.Permissions{})
	noError(t, err)
	isEqual(t, 1, client.ListInstallationsCallCount())
	isEqual(t, 2, client.CreateInstallationTokenCallCount())
}
