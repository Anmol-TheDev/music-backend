package auth

import (
	"context"
	db "fetch-spotify/Db"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	name := r.URL.Query().Get("name")
	password := r.URL.Query().Get("password")

	if password == "" || name == "" {
		http.Error(w, "", http.StatusBadRequest)
	}

	database, err := db.Db()

	if err != nil {
		http.Error(w, "while geting database ", http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	Res, err := database.Collection("users").Doc(name).Get(ctx)

	if err != nil {
		http.Error(w, "user does not exiest ", http.StatusNotFound)
		return
	}
	data := Res.Data()

	dbPassword, _ := data["Password"].(string)

	if dbPassword == password {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("true"))
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"name":   data["Name"].(string),
		"id":     data["Id"].(string),
		"date":   time.Now().Unix(),
		"expiry": time.Now().Add(15 * 24 * time.Hour).Unix(),
	})

	// jwt token
	token, err := claims.SignedString(key)

	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
	}

	// setting cookie
	cookie := http.Cookie{
		Name:     "authToken",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
	
}
