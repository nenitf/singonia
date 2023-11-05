package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(app.addIPToContext)
	r.Use(app.Session.LoadAndSave)

	r.Get("/", app.Home)
	r.Post("/", app.Login)

	r.Group(func(r chi.Router) {
		r.Use(app.auth)
		r.Get("/app", app.Dashboard)
	})

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return r
}
