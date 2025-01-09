package auth

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	db "fetch-spotify/Db"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var key, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

type User struct {
	Name     string
	Id       uuid.UUID
	Password string
	Time     int64
}

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w,"method not allowed",http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	password := r.URL.Query().Get("password")

	if password == "" || name == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	//setting user info
	var user User
	user.Name = name
	user.Id = uuid.New()
	user.Password = password
	user.Time = time.Now().Unix()

	// getting db
	database, err := db.Db()

	if err != nil {
		http.Error(w, "not able to get firestore client", http.StatusInternalServerError)
		return
	}

	ctx := context.Background()


	// checking if user already exiests
	_, err = database.Collection("users").Doc(user.Name).Get(ctx)

	if err != nil {
		database.Collection("users").Doc(user.Name).Set(ctx, user)
	} else {
		http.Error(w, "user already exiests", http.StatusBadRequest)
		return
	}
	// setting claims for jwt
	claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"name":   user.Name,
		"id":     user.Id,
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
	w.WriteHeader(http.StatusOK)
}
