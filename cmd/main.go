package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/rs/zerolog"
	xgithub "github.com/zarix908/gitsync/pkg/x/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := xgithub.NewClient(github.NewClient(tc))
	repos, err := client.GetOwneredRepos(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to get ownered repositories")
	}

	for _, repo := range repos {
		fmt.Println(repo.GetFullName())
		fmt.Println(repo.GetURL())
		fmt.Println()
	}

	fmt.Println(len(repos))
}
