package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

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
