package main

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    hashed_password BLOB NOT NULL, -- Storing as BLOB for byte slice
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

func main() {
	dbName := "users_database.db"

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")

	//createTable(db)
	//fmt.Println("Created table")

	for i := 30; i < 40; i++ {
		name := fmt.Sprintf("Russel_%d", i)
		email := fmt.Sprintf("john%d@gmail.com", i)
		lastID, err := createUser(db, name, email, "Test1234")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Created user", lastID)
	}

	fmt.Println("Created users")
}

func createTable(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func createUser(db *sql.DB, name, email, hashedPassword string) (int64, error) {
	stmt := `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`

	hp, err := bcrypt.GenerateFromPassword([]byte(hashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	res, err := db.Exec(stmt, name, email, string(hp))
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()

}
