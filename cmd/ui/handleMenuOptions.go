package ui

import "fmt"

func HandleOptions() {
	var userMenuOption int
	fmt.Print("Choose a options: ")
	fmt.Scan(&userMenuOption)

	switch userMenuOption {
	default:
		fmt.Println("TODO#ERROR")
	case 0:
		fmt.Println("#Todo exiting menu")
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
