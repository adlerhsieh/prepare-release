package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/google/go-github/v30/github"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	repoOwner               string
	repoName                string
	client                  *github.Client
	ignoreMilestoneNotFound bool
}

func NewGitHubClient(ctx context.Context) (*GitHubClient, error) {
	var ignoreMilestoneNotFound bool
	var err error
	ignoreMilestoneNotFoundEnv := os.Getenv("IGNORE_MILESTONE_NOT_FOUND")

	if ignoreMilestoneNotFoundEnv != "" {
		ignoreMilestoneNotFound, err = strconv.ParseBool(ignoreMilestoneNotFoundEnv)
		if err != nil {
			return &GitHubClient{}, err
		}
	}
	return &GitHubClient{
		repoOwner:               os.Getenv("REPO_OWNER"),
		repoName:                os.Getenv("REPO"),
		client:                  newGitHubClient(ctx),
		ignoreMilestoneNotFound: ignoreMilestoneNotFound,
	}, nil
}

func newGitHubClient(ctx context.Context) *github.Client {
	token := os.Getenv("GITHUB_TOKEN")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
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
