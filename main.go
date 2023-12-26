package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Snip-Snap"))
	})

	log.Print("Server started on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
