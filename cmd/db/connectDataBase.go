package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "postgres"
)

func ConnectDataBase() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// check db
	err = db.Ping()
	CheckError(err)
	fmt.Println("Connected")

	// Create Table

	createTableQuery := `
           CREATE TABLE IF NOT EXISTS contacts(
           id SERIAL PRIMARY KEY,
           name TEXT,
           phone TEXT);
 `
	_, err = db.Exec(createTableQuery)
	CheckError(err)

	fmt.Println("Table Created")
	return db, err
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
