package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	defaultMiddleware := alice.New(app.recover, app.logger)
	secureMiddleware := alice.New(app.session.Enable)
	//staticDir := filepath.Join(".", "learning-go-hnews", "public")
	//fmt.Println("Fayllar qidirilayotgan to'liq yo'l:", staticDir)
	// handle CSS file
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(app.publicPath))))

	mux.Handle("/", secureMiddleware.ThenFunc(app.home))
	mux.Handle("/login", secureMiddleware.ThenFunc(app.login))
	mux.Handle("/submit", secureMiddleware.Append(app.requireAuth).ThenFunc(app.submit))
	mux.Handle("/register", secureMiddleware.ThenFunc(app.register))
	mux.HandleFunc("/about", app.about)
	mux.HandleFunc("/contact", app.contact)

	//handler := app.recover(app.logger(app.session.Enable(mux)))
	return defaultMiddleware.Then(mux)
}
