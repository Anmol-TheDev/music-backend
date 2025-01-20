package utils

import (
	"github.com/zmb3/spotify"
)

func SearchTrack(query string) (*spotify.SearchResult, error) {
	var Client = Token()
	search, err := Client.Search(query, spotify.SearchTypeTrack)

	return search, err
}

func SearchPlaylist(query string) (*spotify.SearchResult, error) {
	var Client = Token()
	search, err := Client.Search(query, spotify.SearchTypePlaylist)

	return search, err
}

func SearchArtist(query string) (*spotify.SearchResult, error) {
	var Client = Token()
	search, err := Client.Search(query, spotify.SearchTypeAlbum)

	return search, err
}

func GetTrackFromSpotifyId(id string) (*spotify.FullTrack, error) {
	var Client = Token()
	track, err := Client.GetTrack(spotify.ID(id))

	return track,err
}

func GetPlaylistFromSpotifyId(id string) (*spotify.PlaylistTrackPage, error) {
	var Client = Token()
	playlist, err := Client.GetPlaylistTracks(spotify.ID(id))

	return playlist,err
}
