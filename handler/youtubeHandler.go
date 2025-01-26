package handler

import (
	"context"
	"encoding/json"
	"fetch-spotify/utils"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chromedp/chromedp"
)

type YoutubeData struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	YtUrl       string `json:"yturl"`
	DownloadURl string `json:"downloadurl"`
	YTID        string `json:"ytid"`
}

func FromYouTube(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")        // Allow all origins
	w.Header().Set("Access-Control-Allow-Methods", "GET")                         // Allowed methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allow

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
		result.YTID = item.Id.VideoId
	}
	ytURL := ytcRC(result.YTID)
	if ytURL == "" {
		http.Error(w, "server error", http.StatusInternalServerError)
		fmt.Println("err getting ytdlp")
		return
	}
	// ytdlp, err := utils.Ytdlp(result.YtUrl)
	// if err != nil {
	// 	http.Error(w, "server error", http.StatusInternalServerError)
	// 	fmt.Println("err getting ytdlp", err)
	// 	return
	// }

	result.DownloadURl = ytURL
	fmt.Println(ytURL)

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	w.Write(jsonData)
}

func ytcRC(id string) string {
	// Create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Run chromedp tasks
	var data []map[string]string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://ytc.re/button/mp3/"+id),
		chromedp.WaitVisible(`#down`, chromedp.ByID), // Wait for the button
		chromedp.Click(`#down`, chromedp.ByID),       // Click the button
		chromedp.Sleep(1500*time.Millisecond),
		chromedp.AttributesAll(`#down`, &data), // Get the button attributes
	)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range data {
		if href, ok := d["href"]; ok {
			fmt.Println(href)
			return href
		}
	}
	return ""
}
