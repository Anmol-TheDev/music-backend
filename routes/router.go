package routes

import (
	"fetch-spotify/handler"
	"net/http"
)

func Router() {
	http.HandleFunc("/playlist", handler.HandlePlaylist)
	http.HandleFunc("/notfound", handler.FromYouTube)


}
