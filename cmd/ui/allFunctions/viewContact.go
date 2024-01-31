package ui

import (
	"fmt"

	db "github.com/odmrs/crud-phonebook/cmd/db/all_functions"
)

func ViewOnDataBase() {
	var nameUser string
	fmt.Printf("----------- Show contact Menu ---------\n\n")

	fmt.Print("Enter the name of user: ")
	fmt.Scan(&nameUser)

	userFounded, phoneFounded := db.ViewContacts(nameUser)
	if userFounded == "" {
		fmt.Println("\n--------> User not found, try again <-------")
		return
	}

	fmt.Printf("---> %v Found:\n", userFounded)
	fmt.Printf("Name: %v | Phone: %v\n", userFounded, phoneFounded)
	fmt.Scanln()
}
