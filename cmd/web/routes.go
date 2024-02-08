package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	//mux.HandleFunc("/library", app.library)
	//mux.HandleFunc("/international", app.international)
	//mux.HandleFunc("/events", app.events)
	mux.HandleFunc("/news/about", app.aboutPage)
	mux.HandleFunc("/news/create", app.createNews)
	mux.HandleFunc("/foods", app.foods)
	mux.HandleFunc("/foods/create", app.createFood)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return app.logRequest(secureHeaders(mux))
}
