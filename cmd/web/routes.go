package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	dynamicMiddleware := alice.New(app.session.Enable, noSurf)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/news/about", dynamicMiddleware.ThenFunc(app.aboutPage))
	mux.Get("/news/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createNews))
	mux.Post("/news/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createNews))
	mux.Get("/foods", dynamicMiddleware.ThenFunc(app.foods))
	mux.Get("/foods/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createFood))
	mux.Post("/foods/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createFood))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.ThenFunc(app.logoutUser))

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
