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
	repoOwner               string
	repoName                string
	client                  *github.Client
	ignoreMilestoneNotFound bool
}

func NewGitHubClient(ctx context.Context, repoOwner, repoName string, ignoreMilestoneNotFound bool) *GitHubClient {
	return &GitHubClient{
		repoOwner:               repoOwner,
		repoName:                repoName,
		client:                  newGitHubClient(ctx),
		ignoreMilestoneNotFound: ignoreMilestoneNotFound,
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
		return "", errors.New("no release found")
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

	if targetMilestone != nil {
		closedState := "closed"
		targetMilestone.State = &closedState

		_, _, err = g.client.Issues.EditMilestone(ctx, g.repoOwner, g.repoName, *targetMilestone.Number, targetMilestone)
		if err != nil {
			return err
		}

		return nil
	}

	if g.ignoreMilestoneNotFound {
		return nil
	}

	return fmt.Errorf("no milestone is matching the tag name. tagName=%s", title)
}
