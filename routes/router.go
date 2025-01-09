package routes

import (
	"fetch-spotify/handler"
	auth "fetch-spotify/handler/Auth"
	"net/http"
)

func Router() {
	http.HandleFunc("/playlist", handler.HandlePlaylist)
	http.HandleFunc("/song", handler.FromYouTube)
	http.HandleFunc("/searchsong", handler.HandleSearch)
	http.HandleFunc("/homesuggestion", handler.HnadleHomeSuggestion)
	http.HandleFunc("/auth/register", auth.Register)
	http.HandleFunc("/auth/login", auth.Login)
	http.HandleFunc("/auth/tokencheck",auth.CheckToken)
}
