package utils

import (
	"fmt"
	"log"

	"github.com/zmb3/spotify"
)

func Recommendation() {
	songName := "Shape of You"
	spotifyClient := Token()

	result, err := spotifyClient.Search(songName, spotify.SearchTypeTrack)

	if err != nil {
		log.Fatalln("err while getting track ", err)

	}

	songId := result.Tracks.Tracks[0].ID

	if songId == "" {
		panic("not able to get songID from spotify")
	}

	recommendations, err := spotifyClient.GetRecommendations(spotify.Seeds{
		Tracks: []spotify.ID{songId},
	}, nil, nil)
	if err != nil {
		log.Fatalln("err while getting recommendation", err)
	}
	for _, track := range recommendations.Tracks {
		fmt.Printf("- %s by %s\n", track.Name, track.Artists[0].Name)
	}
}
