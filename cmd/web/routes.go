package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/olahol/melody"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	m := melody.New()
	r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		err := m.HandleRequest(w, r)
		if err != nil {
			log.Fatal(err)
		}
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.Recoverer)
		r.Use(app.addIPToContext)
		r.Use(app.Session.LoadAndSave)

		r.Get("/", app.Home)
		r.Post("/", app.Login)

		r.Group(func(r chi.Router) {
			r.Use(app.auth)
			r.Get("/app", app.Dashboard)
		})
	})

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return r
}
