package ui

import (
	"fmt"

	db "github.com/odmrs/crud-phonebook/cmd/db/all_functions"
)

func EditAUser() {
	allNames := db.ViewAllContactsOnDB()
	var newUser, newPhone, defaultUser string
	fmt.Println("---> See all contacts below <---")

	for name, phone := range allNames {
		fmt.Printf("--> %v | %v\n", name, phone)
	}

	fmt.Print("Enter which contact your want edit: ")
	fmt.Scan(&defaultUser)
	fmt.Print("Enter new name to contact: ")
	fmt.Scan(&newUser)
	fmt.Print("Enter new phone to contact: ")
	fmt.Scan(&newPhone)

	db.EditAUser(newUser, newPhone, defaultUser)

	fmt.Printf("---> %v edited", defaultUser)
}
