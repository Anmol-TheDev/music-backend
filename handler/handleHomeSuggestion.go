package handler

import (
	"encoding/json"
	"fetch-spotify/utils"
	"net/http"

	"github.com/zmb3/spotify"
)

type trackStr struct {
	Name    string          `json:"name"`
	Id      string          `json:"id"`
	Images  []spotify.Image `json:"images"`
	Views   int64           `json:"views"`
	Artiest []string        `json:"artiest"`
}

type playlistStr struct {
	Name   string                 `json:"name"`
	Id     string                 `json:"id"`
	Tracks spotify.PlaylistTracks `json:"tracks"`
	Image  []spotify.Image        `json:"images"`
}

type respDataStr struct {
	Tracks   []trackStr    `json:"tracks"`
	Playlist []playlistStr `json:"playlists"`
}

func HnadleHomeSuggestion(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")        // Allow all origins
	w.Header().Set("Access-Control-Allow-Methods", "GET")                         // Allowed methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allowed headers

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query().Get("query")

	if query == "" {
		http.Error(w, "empty query", http.StatusBadRequest)
		return
	}
	// getting tracks
	var data respDataStr

	trackRes, err := utils.SearchTrack(query)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	for _, item := range trackRes.Tracks.Tracks {
		temp := trackStr{
			Name:   item.Name,
			Id:     item.ID.String(),
			Images: item.Album.Images,
		}
		for _, value := range item.Artists {
			temp.Artiest = append(temp.Artiest, value.Name)
		}
		data.Tracks = append(data.Tracks, temp)
	}

	// getting playlists

	playlistRes, err := utils.SearchPlaylist(query)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	for _, item := range playlistRes.Playlists.Playlists {
		tempPlaylist := playlistStr{
			Name:   item.Name,
			Id:     item.ID.String(),
			Image:  item.Images,
			Tracks: item.Tracks,
		}
		data.Playlist = append(data.Playlist, tempPlaylist)
	}

	// converting into json
	jsonData, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
	w.Write(jsonData)
}
