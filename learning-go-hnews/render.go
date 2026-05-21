package main

import (
	"errors"
	"net/http"
)

func (app *application) render(w http.ResponseWriter, filename string, data interface{}) {
	if app.tp == nil {
		http.Error(w, errors.New("template render is not rendering").Error(), http.StatusInternalServerError)
	}
	app.tp.Render(w, filename, data)
}
