package utils

import (
	"context"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func Youtube(query string) youtube.SearchListResponse {

	ctx := context.Background()

	key := os.Getenv("YOUTUBE_KEY")

	service, err := youtube.NewService(ctx, option.WithAPIKey(key))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}
	call := service.Search.List([]string{"id", "snippet"}).
		Q(query).
		Type("video").
		MaxResults(1)

	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making API call: %v", err)
	}

	return *response
}
