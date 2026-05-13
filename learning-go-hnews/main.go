package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	userRepo UserRepository
}

func main() {

	db, err := connectToDatabase("users_database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
		infoLog:  log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
		userRepo: NewSQLUserRepository(db),
	}

	fmt.Println("Connecting to database...")
	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}

func connectToDatabase(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", name)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
