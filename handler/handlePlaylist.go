package handler

import (
	"encoding/json"
	"fetch-spotify/modles"
	"fetch-spotify/utils"
	"net/http"

	"github.com/zmb3/spotify"
)


type sptfyPlaylist struct {
	Name   string  `json:"name"`
	Tracks []modles.TrackStr `json:"tracks"`
	Images []spotify.Image `json:"images"`
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
	playlistTracks, _ := spotifyClient.GetPlaylistTracks(playlistId)
	playlistRes, err := spotifyClient.GetPlaylist(playlistId)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	var playlist sptfyPlaylist

	playlist.Name = playlistRes.Name
	playlist.Images = playlistRes.Images

	for _, item := range playlistTracks.Tracks {

		var temp modles.TrackStr
		temp.Name = item.Track.Name
		temp.Id = item.Track.ID.String()
		playlist.Tracks = append(playlist.Tracks,temp)

	}

	utils.GetTrackfromJio(&playlist.Tracks)

	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(playlist)
	if err != nil {
		http.Error(w, "server error ", http.StatusInternalServerError)
	}
	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}
