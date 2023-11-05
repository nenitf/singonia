package main

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/nenitf/singonia/pkg/data"
)

type application struct {
	Session *scs.SessionManager
}

func main() {
	gob.Register(data.User{})
	app := application{}
	app.Session = getSession()

	log.Println("Starting server on port 8000...")

	err := http.ListenAndServe(":8000", app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
