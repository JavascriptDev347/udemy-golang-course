package main

import (
	"database/sql"
	"fmt"
	"log"
	"udemy/section_11/6-repository/repository"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS profile (
    user_id INTEGER PRIMARY KEY REFERENCES users(user_id) ON DELETE CASCADE,
    avatar TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

func main() {

	dbName := "users_database.db"

	db, err := connectToDatabase(dbName)
	checkErr(err)

	fmt.Println("connected to database")

	repo := repository.NewSQLUserRepository(db)

	//user, err := repo.CreateUser("Russel", "rull@gmail.com", "123321", "avatar")
	//if err != nil {
	//	return
	//}
	//fmt.Println(user)

	printUsers(repo)
}

func printUsers(repo repository.UserRepository) {
	users, err := repo.GetUsers()
	checkErr(err)
	for _, user := range users {
		fmt.Println(user.ID, user.Email)
	}
}

func checkErr(err error) {
	if err != nil {
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
