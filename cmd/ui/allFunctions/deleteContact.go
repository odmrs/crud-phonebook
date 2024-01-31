package ui

import (
	"fmt"

	db "github.com/odmrs/crud-phonebook/cmd/db/all_functions"
)

func DeleteAUser() {
	allNames := db.ViewAllContactsOnDB()
	var deleteUser string
	fmt.Println("---> See all contacts below <---")

	for name, phone := range allNames {
		fmt.Printf("--> %v | %v\n", name, phone)
	}

	fmt.Print("Enter what contact your want delete: ")
	fmt.Scan(&deleteUser)

	db.DeleteAUser(deleteUser)

	fmt.Printf("---> %v deleted", deleteUser)
}
