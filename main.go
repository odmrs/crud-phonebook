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

	// Conection
	e := db.Ping()
	CheckError(e)
	fmt.Println("Created!!")

	// Create table

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS contacts (
	id SERIAL PRIMARY KEY,
	name TEXT,
	phone TEXT
	)
`
	_, err = db.Exec(createTableQuery)
	CheckError(err)
	fmt.Println("Table Created")

	// Insert data

	insetTest := `insert into "contacts"("name", "phone") values('Marcos', '321313')`

	_, err = db.Exec(insetTest)
	CheckError(err)

	fmt.Println("Insert data!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
