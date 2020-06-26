package main

import (
	"context"
	"errors"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	client, err := NewGitHubClient(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	tagName, exists := os.LookupEnv("RELEASE_NAME")
	if !exists {
		log.Fatalln(errors.New("the RELEASE_NAME env variable needs to be set"))
	}

	err = client.CloseMilestone(ctx, tagName)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
