package handler

import (
	"encoding/json"
	"fetch-spotify/utils"
	"net/http"

	"github.com/zmb3/spotify"
)

type Track struct {
	Name    string        `json:"name"`
	SptfyID string   `json:"id"`
	Images  spotify.Image `json:"images"`
}
type sptfyPlaylist struct {
	Name   string  `json:"name"`
	Tracks []Track `json:"tracks"`
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

	spotifyClient := utils.Token()
	playlistId := spotify.ID(id)
	playlistTracks, err := spotifyClient.GetPlaylistTracks(playlistId)
	// playlistRes, err := spotifyClient.GetPlaylist(playlistId)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	var playlist sptfyPlaylist

	for _, item := range playlistTracks.Tracks {

		var temp track
		temp.Name = item.Track.Name
		temp.Id = item.Track.ID.String()

		playlist.Tracks = append(playlist.Tracks)

	}

	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(playlist)
	if err != nil {
		http.Error(w, "server error ", http.StatusInternalServerError)
	}
	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}
