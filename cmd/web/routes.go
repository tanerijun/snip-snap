package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes(staticDir string) http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	dynamicMiddlewares := alice.New(app.sessionManager.LoadAndSave)

	fileServer := http.FileServer(http.Dir(staticDir))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.Handler(http.MethodGet, "/", dynamicMiddlewares.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/snippet/view/:id", dynamicMiddlewares.ThenFunc(app.snippetView))
	router.Handler(http.MethodGet, "/snippet/create", dynamicMiddlewares.ThenFunc(app.snippetCreate))
	router.Handler(http.MethodPost, "/snippet/create", dynamicMiddlewares.ThenFunc(app.snippetCreatePost))

	standardMiddlewares := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standardMiddlewares.Then(router)
}
