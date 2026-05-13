package main

import (
	"html/template"
	"net/http"
)

func (app *application) render(w http.ResponseWriter, filename string, data interface{}) {

	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
}
