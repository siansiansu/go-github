package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/google/go-github/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

const (
	owner       = "siansiansu"
	repo        = "siansiansu"
	contentPath = "README.md"
	branch      = "main"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read config file: %w \n", err))
	}
	token := viper.GetString("github_token")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	c := github.NewClient(tc)

	// Print Octocat
	fmt.Println(c.Octocat(ctx, "hello world!"))

	_, _, err = c.Repositories.List(ctx, "", nil)
	if err != nil {
		panic(fmt.Errorf("Failed to list repositories"))
	}

	// Get File From Github repositories
	readCloser, err := c.Repositories.DownloadContents(ctx, owner, repo, contentPath, &github.RepositoryContentGetOptions{
		Ref: branch,
	})
	defer readCloser.Close()

	if err != nil {
		panic(fmt.Errorf("Failed to get file"))
	}
	content, err := ioutil.ReadAll(readCloser)
	if err != nil {
		panic(fmt.Errorf("Failed to read file"))
	}
	fmt.Println(string(content))
}
