## githubapp

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/telia-oss/githubapp)
[![latest release](https://img.shields.io/github/v/release/telia-oss/githubapp?style=flat-square)](https://github.com/telia-oss/githubapp/releases/latest)
[![build status](https://img.shields.io/github/workflow/status/telia-oss/githubapp/test?label=build&logo=github&style=flat-square)](https://github.com/telia-oss/githubapp/actions?query=workflow%3Atest)
[![code quality](https://goreportcard.com/badge/github.com/telia-oss/githubapp?style=flat-square)](https://goreportcard.com/report/github.com/telia-oss/githubapp)

A small Go package for handling authentication with a Github App using owner login and repository names instead of UUIDs. Installations and repositories are
cached internally and refreshed (lazily) on a set interval, to reduce the number of `List*` API calls against the Apps API.

### Usage

```go
package main

import (
	"github.com/telia-oss/githubapp"

	"github.com/google/go-github/v29/github"
)

func main() {
    client, err := githubapp.NewClient(911, []byte("private-key"))
    if err != nil {
        panic(err)
    }

    app := githubapp.New(client)
    
    token, err := app.CreateInstallationToken(
        "telia-oss",
        []string{"githubapp"},
		&githubapp.Permissions{
            Metadata: github.String("read"),
        },
    )
    if err != nil {
        panic(err)
    }
}
```
