package main

import (
	"html/template"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

type Book struct {
	Title  string
	Author string
}

func temp(w http.ResponseWriter, r *http.Request) {
	book := Book{"Building Web Apps with Go", "Jeremy Saenz"}
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter()
	r.Path("/").Methods(http.MethodGet).HandlerFunc(temp)
	http.ListenAndServe(":3000", r)
}
