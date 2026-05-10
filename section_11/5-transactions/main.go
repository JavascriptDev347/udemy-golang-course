package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id              INTEGER PRIMARY KEY AUTOINCREMENT,
	name            TEXT NOT NULL,
	email           TEXT NOT NULL UNIQUE,
	hashed_password TEXT NOT NULL,
	created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS profile (
	user_id    INTEGER PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
	avatar     TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	Profile        Profile   `json:"profile"`
}

type Profile struct {
	UserID  int       `json:"user_id"`
	Avatar  string    `json:"avatar"`
	Created time.Time `json:"created"`
}

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

	createTable(db)
	fmt.Println("database connection established")

	userID, err := createUser(db, "test with defer", "test@tets.s", "qkjsq", "kjqwkq")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User created with ID: %d\n", userID)
}

func createTable(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func createUser(db *sql.DB, name, email, hashedPassword, avatar string) (int64, error) {
	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// 1. User qoshish
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, name, email, hashedPassword)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// 2. Profile qoshish
	profileStmt, err := tx.PrepareContext(ctx, "INSERT INTO profile (user_id, avatar, created_at) VALUES (?, ?, CURRENT_TIMESTAMP)")
	if err != nil {
		return 0, err
	}
	defer profileStmt.Close()

	_, err = profileStmt.ExecContext(ctx, userID, avatar)
	if err != nil {
		return 0, err
	}

	// 3. Commit
	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return userID, nil
}
