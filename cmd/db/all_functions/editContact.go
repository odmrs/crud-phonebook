package db

import "github.com/odmrs/crud-phonebook/cmd/db"

func EditAUser(newUser, newPhone, defaultUser string) {
	database, err := db.ConnectDataBase()
	db.CheckError(err)

	defer database.Close()

	sqlUpdate := `UPDATE "contacts" SET "name" = $1, "phone" = $2 WHERE "name" = $3`

	_, err = database.Exec(sqlUpdate, newUser, newPhone, defaultUser)
	db.CheckError(err)
}
