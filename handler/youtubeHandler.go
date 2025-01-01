package handler

import (
	"encoding/json"
	"fetch-spotify/utils"
	"fmt"
	"net/http"
)

type YoutubeData struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Url         string `json:"url"`
	DownloadURl string `json:"downloadurl"`
}

func FromYouTube(w http.ResponseWriter, r *http.Request) {

	var result YoutubeData
	w.WriteHeader(http.StatusOK)
	// checking is query is right
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "Bad Request Boys 😁", http.StatusBadRequest)
		return
	}
	// fetching data from youtube
	fetchedData := utils.Youtube(name)

	for _, item := range fetchedData.Items {
		result.Name = item.Snippet.Title
		result.Image = item.Snippet.Thumbnails.High.Url
		result.Url = fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.Id.VideoId)
	}
	ytdlp, err := utils.Ytdlp(result.Url)
	if err != nil {
		fmt.Println("err while getting url from ytdlp", err)
	}
	fmt.Println(ytdlp)

	result.DownloadURl = string(ytdlp)

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	w.Write(jsonData)
}
