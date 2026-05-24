package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Home page visited")
	app.infoLog.Printf("Session data: %s", app.session.GetString(r, "UserID"))
	app.render(w, r, "index.html", nil)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Login page visited")
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		form := NewForm(r.PostForm)
		form.Required("email", "password").
			MaxLength("email", 255).
			MaxLength("password", 255).
			MinLength("password", 6).
			MinLength("email", 3).
			Matches("email", EmailRX)

		if !form.Valid() {
			app.errorLog.Printf("Form validation failed: %v", form.Errors)
			app.render(w, r, "login.html", nil)
		}
		email := r.FormValue("email")
		password := r.FormValue("password")

		app.infoLog.Printf("Email: %s, Password: %s", email, password)
	}
	app.render(w, r, "login.html", nil)
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Register page visited")
	app.render(w, r, "register.html", nil)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("About page visited")
	app.render(w, r, "about.html", nil)
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Contact page visited")
	app.render(w, r, "contact.html", nil)
}

func (app *application) submit(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Home page visited")
	app.infoLog.Printf("Session data: %s", app.session.GetString(r, "UserID"))
	app.render(w, r, "submit.html", nil)
}
