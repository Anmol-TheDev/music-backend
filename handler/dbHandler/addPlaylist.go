package dbhandler

import "net/http"

func SetPlaylist(w http.ResponseWriter, r *http.Request) {
	  if r.Method != http.MethodPost {
		http.Error(w,"",http.StatusMethodNotAllowed)
		return
	  }

	  
}
