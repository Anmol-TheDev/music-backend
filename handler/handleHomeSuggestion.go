package handler

import (
	"encoding/json"
	"fetch-spotify/utils"
	"net/http"

	"github.com/zmb3/spotify"
)

type trackData struct {
	Name   string          `json:"name"`
	Id     string          `json:"id"`
	Images []spotify.Image `json:"images"`
	Views  int64           `json:"views"`
}

type respDataStr struct {
	Tracks []trackData
}

func HnadleHomeSuggestion(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	if query == "" {
		http.Error(w, "empty query", http.StatusBadRequest)
		return
	}

	var data respDataStr

	trackRes, err := utils.SearchTrack("khatta flow")

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	for _, item := range trackRes.Tracks.Tracks {
		temp := trackData{
			Name:   item.Name,
			Id:     item.ID.String(),
			Images: item.Album.Images,
		}
		data.Tracks = append(data.Tracks, temp)

	}
	jsonData, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}

	w.Write(jsonData)
}
