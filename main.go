package main

import (
	// Generic interface around SQL in general

	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	_ "github.com/odmrs/crud-phonebook/cmd/ui"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "postgres"
)

func main() {
	//	ui.DrawMenu()
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// Closing DB

	defer db.Close()

	// checking db

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
