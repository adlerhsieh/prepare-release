package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	ctx := context.Background()

	repository := strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")
	if len(repository) != 2 {
		log.Fatalln(errors.New(fmt.Sprintf("GITHUB_REPOSITORY env var should be in the format 'octocat/hello-world', got %#q", os.Getenv("GITHUB_REPOSITORY"))))
	}

	var ignoreMilestoneNotFound bool
	if os.Getenv("IGNORE_MILESTONE_NOT_FOUND") != "" {
		var err error
		ignoreMilestoneNotFound, err = strconv.ParseBool(os.Getenv("IGNORE_MILESTONE_NOT_FOUND"))
		if err != nil {
			log.Fatalln(errors.New(fmt.Sprintf("IGNORE_MILESTONE_NOT_FOUND env var should contain a boolean value, got %#q", os.Getenv("IGNORE_MILESTONE_NOT_FOUND"))))
		}
	}

	client := NewGitHubClient(ctx, repository[0], repository[1], ignoreMilestoneNotFound)

	tagName, err := client.GetLatestReleaseTag(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = client.CloseMilestone(ctx, tagName)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
