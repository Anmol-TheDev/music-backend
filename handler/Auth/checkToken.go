package auth

import "net/http"


func CheckToken(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w,"method not allowed",http.StatusMethodNotAllowed)
		return
	}
	_,err := r.Cookie("authToken")

	if err!= nil {
		if err == http.ErrNoCookie {
			http.Error(w,"cookie not found ",http.StatusUnauthorized)
			return
		}
		http.Error(w,"",http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

