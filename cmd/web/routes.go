package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/tanerijun/snip-snap/ui"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	router.Handler(http.MethodGet, "/health", http.HandlerFunc(ping))

	dynamicMiddlewares := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	router.Handler(http.MethodGet, "/", dynamicMiddlewares.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/about", dynamicMiddlewares.ThenFunc(app.about))
	router.Handler(http.MethodGet, "/snippet/view/:id", dynamicMiddlewares.ThenFunc(app.snippetView))
	router.Handler(http.MethodGet, "/user/login", dynamicMiddlewares.ThenFunc(app.userLogin))
	router.Handler(http.MethodPost, "/user/login", dynamicMiddlewares.ThenFunc(app.userLoginPost))
	router.Handler(http.MethodGet, "/user/signup", dynamicMiddlewares.ThenFunc(app.userSignup))
	router.Handler(http.MethodPost, "/user/signup", dynamicMiddlewares.ThenFunc(app.userSignupPost))

	protectedMiddlewares := dynamicMiddlewares.Append(app.requireAuthentication)

	router.Handler(http.MethodGet, "/snippet/create", protectedMiddlewares.ThenFunc(app.snippetCreate))
	router.Handler(http.MethodPost, "/snippet/create", protectedMiddlewares.ThenFunc(app.snippetCreatePost))
	router.Handler(http.MethodGet, "/account/view", protectedMiddlewares.ThenFunc(app.accountView))
	router.Handler(http.MethodGet, "/account/password/update", protectedMiddlewares.ThenFunc(app.accountPasswordUpdate))
	router.Handler(http.MethodPost, "/account/password/update", protectedMiddlewares.ThenFunc(app.accountPasswordUpdatePost))
	router.Handler(http.MethodPost, "/user/logout", protectedMiddlewares.ThenFunc(app.userLogout))

	standardMiddlewares := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standardMiddlewares.Then(router)
}
