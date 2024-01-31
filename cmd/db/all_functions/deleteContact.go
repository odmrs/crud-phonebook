package db

import "github.com/odmrs/crud-phonebook/cmd/db"

func DeleteAUser(userToDelete string) {
	database, err := db.ConnectDataBase()
	db.CheckError(err)

	defer database.Close()

	sqlToDelete := `DELETE FROM contacts WHERE name = $1`
	_, err = database.Exec(sqlToDelete, userToDelete)
	db.CheckError(err)
}
