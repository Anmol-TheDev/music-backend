package main

import (
	"fetch-spotify/routes"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("err while loading dotenv")
		return
	}
	routes.Router()
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("err while turning on server", err)
	}
}
