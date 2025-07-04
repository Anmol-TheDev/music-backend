package routes

import (
	"fetch-spotify/handler"
	auth "fetch-spotify/handler/Auth"
	"net/http"
)

func Router() {
	http.HandleFunc("/getsptfyplaylist", handler.HandlePlaylist)     // id = spotify id / return tracks inside spotify playlist | sptfy = spotify
	http.HandleFunc("/fchytsong", handler.FromYouTube)               // fetching song that is not available on jio savan
	http.HandleFunc("/srchbar", handler.HandleSearch)                // end point for api debounceing
	http.HandleFunc("/homesuggestion", handler.HnadleHomeSuggestion) // home suggestion like tracks and playlist
	http.HandleFunc("/song", handler.Song)
	http.HandleFunc("/auth/register", auth.Register)     // auth register user
	http.HandleFunc("/auth/login", auth.Login)           // auth login user
	http.HandleFunc("/auth/tokencheck", auth.CheckToken) // check user is authorized
	http.HandleFunc("/auth/logout", auth.Logout)         // user logout
}
