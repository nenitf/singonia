package main

import (
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/nenitf/singonia/pkg/data"
)

var pathToTemplates = "./templates/"

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	td := make(map[string]any)

	if app.Session.Exists(r.Context(), "user") {
		http.Redirect(w, r, "/app", http.StatusSeeOther)
	}

	_ = app.render(w, r, "login.page.tpl", &TemplateData{Data: td})
}

func (app *application) Dashboard(w http.ResponseWriter, r *http.Request) {
	user, ok := app.Session.Get(r.Context(), "user").(data.User)

	if ok {
		_ = app.render(w, r, "app.page.tpl", &TemplateData{User: user})
	}
}

type TemplateData struct {
	IP    string
	Data  map[string]any
	Error string
	Flash string
	User  data.User
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) error {
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t), path.Join(pathToTemplates, "base.layout.tpl"))
	if err != nil {
		log.Fatal(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}

	td.IP = app.ipFromContext(r.Context())

	td.Error = app.Session.PopString(r.Context(), "error")
	td.Flash = app.Session.PopString(r.Context(), "flash")

	// execute the template, passing it data, if any
	err = parsedTemplate.Execute(w, td)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	form := NewForm(r.PostForm)
	form.Required("name")

	if !form.Valid() {
		app.Session.Put(r.Context(), "error", "Invalid login credentials")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	name := r.Form.Get("name")
	user := &data.User{Name: name}

	// validar utilização do nome para evitar duplicidade
	err = nil
	if err != nil {
		app.Session.Put(r.Context(), "error", "Name already in use")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	_ = app.Session.RenewToken(r.Context())
	app.Session.Put(r.Context(), "user", user)
	http.Redirect(w, r, "/app", http.StatusSeeOther)
}
