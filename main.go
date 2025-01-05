package main

import (
	"fetch-spotify/routes"
	"fetch-spotify/utils"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	// playlist := utils.FetchFromSpotify()

	// for _, value := range playlist.Tracks.Tracks {
	// 	fmt.Println(value.Track.Name)
	// }
	err := godotenv.Load()
	if err != nil {
		fmt.Println("err while loading dotenv")
	}
	// utils.Youtube(os.Getenv("YOUTUBE_KEY"))

	// fmt.Println("Download URLs:")

	// downloadURLs, err := utils.Ytdlp(videoURL)
	// for i, url := range downloadURLs {
	// 	fmt.Printf("%d: %s\n", i+1, url)
	// }
	utils.Recommendation()
	routes.Router()
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("err while turning on server", err)
	}

}
