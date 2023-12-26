package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Snip-Snap"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TODO: display snippets"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TODO: UI to create snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Server started on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
