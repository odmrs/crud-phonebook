package ui

import (
	"fmt"

	db "github.com/odmrs/crud-phonebook/cmd/db/all_functions"
)

func ViewAllContactsOnDB() {
	fmt.Println("----------- All Contacts ----------")

	allContacts := db.ViewAllContactsOnDB()

	for name, phone := range allContacts {
		fmt.Printf("--> %v | %v \n", name, phone)
	}
	fmt.Scanln()
}
