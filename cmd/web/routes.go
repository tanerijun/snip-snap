package main

import "net/http"

func (app *application) routes(staticDir string) http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
