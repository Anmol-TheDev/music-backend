package main

import (
	"fetch-spotify/routes"
	"fmt"
	"net/http"
)

const videoURL = "https://youtu.be/5dYirJj0I9M?si=c4deBBZ53Wi4uXJQ"

func main() {

	// playlist := utils.FetchFromSpotify()

	// for _, value := range playlist.Tracks.Tracks {
	// 	fmt.Println(value.Track.Name)
	// }
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println("err while loading dotenv")
	// }
	// utils.Youtube(os.Getenv("YOUTUBE_KEY"))

	// fmt.Println("Download URLs:")

	// downloadURLs, err := utils.Ytdlp(videoURL)
	// for i, url := range downloadURLs {
	// 	fmt.Printf("%d: %s\n", i+1, url)
	// }
	routes.Router()
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println("err while turning on server", err)
	}

}
