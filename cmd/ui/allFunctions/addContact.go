package ui

import (
	"fmt"
)

func InsertOnDataBase() {
	var nameUser, phoneUser string
	//userMap := make(map[string]string)
	fmt.Printf("----------- Add contact Menu ---------\n\n")

	fmt.Print("Enter the name of user: ")
	fmt.Scan(&nameUser)
	fmt.Print("Enter the phone of user: ")
	fmt.Scan(&phoneUser)
	fmt.Printf("%v -> name | %v -> phone\n", nameUser, phoneUser)
}
