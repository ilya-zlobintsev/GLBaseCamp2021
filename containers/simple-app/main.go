package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed static
var static embed.FS

var counter uint64 = 0

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", index_handler)

	err := http.ListenAndServe(":8000", router)

	if err != nil {
		log.Printf("Failed initializing web! %v", err)
	}
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(static, "static/page.html.tmpl")

	if err != nil {
		log.Printf("Error loading template %v", err)
	}

	counter++

	t.Execute(w, counter)
}
