package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func (app *application) render(w http.ResponseWriter, filename string, data interface{}) {

	fullPath := filepath.Join(app.templateDir, filename)

	tmpl, err := template.ParseFiles(fullPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
