package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golangcollege/sessions"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	userRepo    UserRepository
	templateDir string
	publicPath  string
	tp          *TemplatesRenderer
	session     *sessions.Session
}

var secret = []byte(`s3cr3t-k3y-asdas-dasdadasdasdas`)

func main() {

	db, err := connectToDatabase("users_database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	session := sessions.New(secret)
	session.Lifetime = 12 * time.Hour
	session.Secure = true
	session.SameSite = http.SameSiteLaxMode

	app := &application{
		errorLog:    log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
		infoLog:     log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
		userRepo:    NewSQLUserRepository(db),
		templateDir: "./templates",
		publicPath:  "./public",
		session:     session,
	}

	app.tp = NewTemplatesRenderer(app.templateDir, true)

	fmt.Println("Connecting to database...")
	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}

func connectToDatabase(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", name) // registers driver, does NOT connect yet
	if err != nil {
		return nil, err
	}

	err = db.Ping() // actually verifies the connection
	if err != nil {
		return nil, err
	}

	return db, nil
}
