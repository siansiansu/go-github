package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	owner       = "siansiansu"
	repo        = "nvim"
	contentPath = "README.md"
	branch      = "main"
	token       = ""
)

func main() {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	c := github.NewClient(tc)

	oct, _, err := c.Octocat(ctx, "hello world!")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(oct)

	_, tree, _, err := c.Repositories.GetContents(ctx, owner, repo, ".", &github.RepositoryContentGetOptions{
		Ref: "main",
	})

	if err != nil {
		panic(err)
	}

	for i := range tree {
		if tree[i].GetType() == "file" {
			fmt.Println(tree[i].GetName())
		}
	}

	// // Get File From Github repositories
	readCloser, err := c.Repositories.DownloadContents(ctx, owner, repo, contentPath, &github.RepositoryContentGetOptions{
		Ref: branch,
	})

	defer readCloser.Close()

	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(readCloser)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
