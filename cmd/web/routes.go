package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes(staticDir string) http.Handler {
	router := httprouter.New()

	fileServer := http.FileServer(http.Dir(staticDir))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/snippet/view", app.snippetView)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreate)

	standardMiddlewares := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standardMiddlewares.Then(router)
}
