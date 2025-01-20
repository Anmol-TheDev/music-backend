package handler

import (
	"encoding/json"
	"fetch-spotify/modles"
	"fetch-spotify/utils"
	"fmt"
	"net/http"
)

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
	var data modles.RespDataStr

	trackRes, err := utils.SearchTrack(query)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	for _, item := range trackRes.Tracks.Tracks {
		temp := modles.TrackStr{
			Name:   item.Name,
			Id:     item.ID.String(),
			Images: item.Album.Images,
			Source: "jiosavan",
		}
		for _, value := range item.Artists {
			temp.Artist = append(temp.Artist, value.Name)
		}
		data.Tracks = append(data.Tracks, temp)
	}

	utils.GetTrackfromJio(&data.Tracks)

	for i := range data.Tracks {
		if data.Tracks[i].DownloadUrl == "" {
			data.Tracks[i].Source = "yt"
		}
	}

	// getting playlists

	playlistRes, err := utils.SearchPlaylist(query)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	for i, item := range playlistRes.Playlists.Playlists {
		if i == 10 {
			break
		}
		tempPlaylist := modles.PlaylistStr{
			Name:   item.Name,
			Id:     item.ID.String(),
			Image:  item.Images,
		}
		data.Playlist = append(data.Playlist, tempPlaylist)
	}

	// converting into json
	jsonData, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
	w.Write(jsonData)

	fmt.Println(r.URL, " served!")
}
