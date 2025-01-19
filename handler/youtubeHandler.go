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
	YtUrl       string `json:"yturl"`
	DownloadURl string `json:"downloadurl"`
}

func FromYouTube(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var result YoutubeData
	w.WriteHeader(http.StatusOK)

	// checking is query is right
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "Bad Request ", http.StatusBadRequest)
		return
	}
	// fetching data from youtube
	fetchedData := utils.Youtube(name)
	for _, item := range fetchedData.Items {
		result.Name = item.Snippet.Title
		result.Image = item.Snippet.Thumbnails.High.Url
		result.YtUrl = fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.Id.VideoId)
	}
	ytdlp, err := utils.Ytdlp(result.YtUrl)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		fmt.Println("err getting ytdlp", err)
	}

	result.DownloadURl = string(ytdlp)

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	w.Write(jsonData)
}
