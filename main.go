package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

type blog struct {
	Title       string `json:"title"`
	DateTime    string `json:"date"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

func setup() {

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	k := os.Getenv("MONGOURI")
	fmt.Println(k)
	r := mux.NewRouter()
	handler := cors.Default().Handler(r)
	http.ListenAndServe(":4000", handler)
}
