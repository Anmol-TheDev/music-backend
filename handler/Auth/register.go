package auth

import "net/http"



func Register(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	password := r.URL.Query().Get("password")

	if password == "" || name == "" {
		http.Error(w,"",http.StatusBadRequest)
	}

	
}