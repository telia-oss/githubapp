// +build e2e

package githubapp_test

import (
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/telia-oss/githubapp"
)

var (
	appIntegrationID  = os.Getenv("GITHUB_APP_INTEGRATION_ID")
	appPrivateKeyFile = os.Getenv("GITHUB_APP_PRIVATE_KEY_FILE")
	targetOwner       = os.Getenv("GITHUB_APP_TARGET_ORG")
	targetRepository  = os.Getenv("GITHUB_APP_TARGET_REPOSITORY")
)

func TestGithubAppE2E(t *testing.T) {
	integrationID, err := strconv.ParseInt(appIntegrationID, 10, 64)
	noError(t, err)

	privateKey, err := ioutil.ReadFile(appPrivateKeyFile)
	noError(t, err)

	client, err := githubapp.NewClient(integrationID, privateKey)
	noError(t, err)

	app := githubapp.New(client)

	token, err := app.CreateInstallationToken(
		targetOwner,
		[]string{targetRepository},
		&githubapp.Permissions{},
	)
	noError(t, err)

	for _, r := range token.Repositories {
		isEqual(t, targetRepository, r.GetName())
	}
}
