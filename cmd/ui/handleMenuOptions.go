package ui

import (
	"fmt"

	ui "github.com/odmrs/crud-phonebook/cmd/ui/allFunctions"
)

func HandleOptions() {
	var userMenuOption int
	fmt.Print("Choose a options: ")
	fmt.Scan(&userMenuOption)

	switch userMenuOption {
	default:
		ui.DrawDefaultScreen()
	case 0:

	case 1:
		fmt.Println("#TODO ADD CONTACT")
	case 2:
		fmt.Println("#TODO View just one contact")
	case 3:
		fmt.Println("#TODO Edit just one contact")
	case 4:
		fmt.Println("#TODO Delete a contact")
	case 5:
		fmt.Println("#TODO View all contacts")
	}
}
