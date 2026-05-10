package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

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

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
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

	fmt.Println("Connected to database")

	u1, err := GetUserByEmail(db, "john2@gmail.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Birinchi get qilindi:", u1)
	bs, err := json.MarshalIndent(u1, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Marshal: ", string(bs))

	// get users
	users, err := GetUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	us, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Get users:", string(us))
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

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	stmt := `SELECT id,email,name,hashed_password,created_at FROM users WHERE email=?`

	// bind email to sql
	row := db.QueryRow(stmt, email)
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.HashedPassword, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func GetUsers(db *sql.DB) ([]User, error) {
	stmt := `SELECT id,name,email,hashed_password,created_at FROM users`
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.HashedPassword, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
