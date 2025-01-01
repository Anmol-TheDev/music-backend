package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func Youtube(query string) youtube.SearchListResponse {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("err while loading env", err)
	}
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
