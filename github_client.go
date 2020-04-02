package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/google/go-github/v30/github"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	repoOwner string
	repoName  string
	client    *github.Client
}

func NewGitHubClient(ctx context.Context) *GitHubClient {
	return &GitHubClient{
		repoOwner: os.Getenv("REPO_OWNER"),
		repoName:  os.Getenv("REPO"),
		client:    newGitHubClient(ctx),
	}
}

func newGitHubClient(ctx context.Context) *github.Client {
	token := os.Getenv("GITHUB_TOKEN")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func (g *GitHubClient) GetLatestReleaseTag(ctx context.Context) (string, error) {
	latestRelease, _, err := g.client.Repositories.GetLatestRelease(ctx, g.repoOwner, g.repoName)
	if err != nil {
		return "", err
	}
	if latestRelease == nil {
		return "", errors.New("No releae found")
	}

	return *latestRelease.TagName, nil
}

func (g *GitHubClient) CloseMilestone(ctx context.Context, title string) error {
	milestones, _, err := g.client.Issues.ListMilestones(ctx, g.repoOwner, g.repoName, nil)
	if err != nil {
		return err
	}

	var targetMilestone *github.Milestone
	for _, m := range milestones {
		if *m.Title == title {
			targetMilestone = m
		}
	}
	if targetMilestone == nil {
		return fmt.Errorf("No milestone is matching the tag name. tagName=%s", title)
	}

	closedState := "closed"
	targetMilestone.State = &closedState

	_, _, err = g.client.Issues.EditMilestone(ctx, g.repoOwner, g.repoName, *targetMilestone.Number, targetMilestone)
	if err != nil {
		return err
	}

	return nil
}
