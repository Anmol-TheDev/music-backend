package utils

import (
	"fmt"

	"github.com/zmb3/spotify"
)

type Data struct {
	Item struct {
		Name string `json:"name"`
	} `json:"item"`
}

func FetchFromSpotify(id string) *spotify.PlaylistTrackPage {

	spotifyClient := Token()
	playlistId := spotify.ID(id)
	playlist, err := spotifyClient.GetPlaylistTracks(playlistId)
	fmt.Println(playlistId)
	if err != nil {
		fmt.Println("error while getting playlist", err)
	}
	return playlist
}
