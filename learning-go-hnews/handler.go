package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Home page visited")
	app.render(w, "index.html", nil)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Login page visited")
	app.render(w, "login.html", nil)
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Register page visited")
	app.render(w, "register.html", nil)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("About page visited")
	app.render(w, "about.html", nil)
}

func (app *application) contact(writer http.ResponseWriter, request *http.Request) {
	app.infoLog.Println("Contact page visited")
	app.render(writer, "contact.html", nil)
}
