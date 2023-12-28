package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// CLI flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")

	flag.Parse()

	// HTTP
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(*staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Server started on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
