package db

import "github.com/odmrs/crud-phonebook/cmd/db"

func ViewAllContactsOnDB() map[string]string {
	allNamesPhones := make(map[string]string)

	database, err := db.ConnectDataBase()
	db.CheckError(err)

	rows, err := database.Query(`SELECT "name", "phone" FROM "contacts"`)
	db.CheckError(err)

	defer rows.Close()

	for rows.Next() {
		var name, phone string
		err = rows.Scan(&name, &phone)
		db.CheckError(err)
		allNamesPhones[name] = phone
	}

	return allNamesPhones
}
