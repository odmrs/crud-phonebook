package db

import (
	"github.com/odmrs/crud-phonebook/cmd/db"
)

func ViewContacts(name string) (string, string) {
	database, err := db.ConnectDataBase()
	db.CheckError(err)
	rows, err := database.Query(`Select "name", "phone" from "contacts"`)
	db.CheckError(err)
	defer rows.Close()

	for rows.Next() {
		var dbName, dbPhone string

		err = rows.Scan(&dbName, &dbPhone)

		db.CheckError(err)

		if name == dbName {
			return dbName, dbPhone
		}
	}
	db.CheckError(err)
	return "", ""
}
