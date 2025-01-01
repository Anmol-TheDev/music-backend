package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
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
func Token() spotify.Client {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("err while loading env", err)
	}
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSECRET := os.Getenv("SPOTIFY_CLIENT_SECRET")
	ctx := context.Background()
	var (
		clientId     = clientID
		clientSecret = clientSECRET
		tokenURL     = "https://accounts.spotify.com/api/token"
	)

	credential := &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
	}
	token, err := credential.Token(ctx)

	if err != nil {
		fmt.Println("err while getting token ", err)
	}

	client := spotify.Authenticator{}.NewClient(token)

	return client
}
