package main

import (
	"fmt"
	"net/http"
)

func (app *application) logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		app.infoLog.Printf("%s %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			w.Header().Set("Connection", "close")
			app.serverError(w, fmt.Errorf("%s", recover()))
		}()

		next.ServeHTTP(w, r)
	})
}
