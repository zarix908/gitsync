package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

type Client interface {
	GetOwneredRepos(context.Context) ([]*github.Repository, error)
}

func NewClient(githubClient *github.Client) Client {
	return &client{client: githubClient}
}

type client struct {
	client *github.Client
}

func (c *client) GetOwneredRepos(ctx context.Context) ([]*github.Repository, error) {
	repositories := make([]*github.Repository, 0)
	nextPage := 0

	for {
		repos, resp, err := c.client.Repositories.List(ctx, "", &github.RepositoryListOptions{
			Affiliation: "owner",
			ListOptions: github.ListOptions{Page: nextPage, PerPage: 50},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to list repositories: %w", err)
		}

		repositories = append(repositories, repos...)
		if resp.NextPage == 0 {
			break;
		}

		nextPage = resp.NextPage
	}

	return repositories, nil
}
