package db

import (
	"github.com/odmrs/crud-phonebook/cmd/db"
)

func AddContactsToDB(name, phone string) {
	database, err := db.ConnectDataBase()
	db.CheckError(err)

	defer database.Close()

	insertData := `insert into "contacts"("name", "phone") values ($1, $2)`

	_, err = database.Exec(insertData, name, phone)
	db.CheckError(err)
}
