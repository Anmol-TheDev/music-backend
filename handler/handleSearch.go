package handler

import (
	"encoding/json"
	"fetch-spotify/utils"
	"net/http"
)

type track struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
type searchResult struct {
	Tracks []track  `json:"tracks"`
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	var data searchResult
	if query == "" {
		http.Error(w, "please provide valid query", http.StatusBadRequest)
		return
	}

	resp, err := utils.Search(query)

	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	for _, item := range resp.Tracks.Tracks {
		temp := track{
			Name: item.Name,
			Id:   string(item.ID),
		}
		temp.Name = item.Name
		data.Tracks = append(data.Tracks, temp)
	}
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}
