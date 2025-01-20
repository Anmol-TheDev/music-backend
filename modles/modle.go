package modles

import "github.com/zmb3/spotify"

type TrackStr struct {
	Id          string          `json:"id"`
	Name        string          `json:"name"`
	Images      []spotify.Image `json:"images"`
	Artist      []string        `json:"artist"`
	Source      string          `json:"source"`
	DownloadUrl string          `json:"downloadurl"`
}

type PlaylistStr struct {
	Name   string                 `json:"name"`
	Id     string                 `json:"id"`
	Tracks spotify.PlaylistTracks `json:"tracks"`
	Image  []spotify.Image        `json:"images"`
}

type RespDataStr struct {
	Tracks   []TrackStr    `json:"tracks"`
	Playlist []PlaylistStr `json:"playlists"`
}
