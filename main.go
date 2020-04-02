package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	client := NewGitHubClient(ctx)
	tagName, err := client.GetLatestReleaseTag(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = client.CloseMilestone(ctx, tagName)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
