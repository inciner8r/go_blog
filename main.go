package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type blog struct {
	Title       string `json:"title"`
	DateTime    string `json:"date"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

func main() {
	r := mux.NewRouter()
	http.ListenAndServe(":4000", r)
}
