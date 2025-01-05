package utils

import (
	"fmt"
	"log"

	"github.com/zmb3/spotify"
)

type Seeds struct {
	Artists []string
	Tracks  []string
	Genres  []string
}

func Recommendation() {
	client := Token()
	seedArtists := []spotify.ID{"4NHQUGzhtTLFvgF5SZesLK"}
	seedTracks := []spotify.ID{"3n3Ppam7vgaVa1iaRUc9Lp"}
	seedGenres := []string{"pop"}

	limit := 20

	trackAttributes := spotify.NewTrackAttributes().
		MaxAcousticness(0.15).
		TargetDanceability(0.8).
		MinEnergy(0.5).
		TargetTempo(120)

	seeds := spotify.Seeds{
		Artists: seedArtists,
		Tracks:  seedTracks,
		Genres:  seedGenres,
	}

	recommendations, err := client.GetRecommendations(seeds, trackAttributes, &spotify.Options{
		Limit: &limit,
	})
	if err != nil {
		log.Printf("Error fetching recommendations: %v", err)
		return
	}

	fmt.Println("Recommended Tracks:")
	for _, track := range recommendations.Tracks {
		fmt.Println(track.Name, track.URI)
	}
}
