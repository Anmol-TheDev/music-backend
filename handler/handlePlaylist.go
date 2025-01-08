package handler

import (
	"encoding/json"
	"fetch-spotify/utils"
	"net/http"
)

type Data struct {
	Tracks []string `json:"tracks"`
}

var Array []string

func HandlePlaylist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is requied", http.StatusBadRequest)
	}

	var spotifyData Data
	resp := utils.FetchFromSpotify(id)
	for _, value := range resp.Tracks {
		Array = append(Array, value.Track.Name)
	}
	spotifyData.Tracks = Array
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(spotifyData)
	if err != nil {
		http.Error(w, "server error ", http.StatusInternalServerError)
	}
	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}
