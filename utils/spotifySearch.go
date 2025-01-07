package utils

import (
	"github.com/zmb3/spotify"
)

func Search(query string) (*spotify.SearchResult, error) {

	client := Token()

	search, err := client.Search(query, spotify.SearchTypeTrack)

	return search, err
}
