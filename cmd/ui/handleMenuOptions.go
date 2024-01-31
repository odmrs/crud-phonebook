package ui

import (
	"fmt"

	ui "github.com/odmrs/crud-phonebook/cmd/ui/allFunctions"
)

func HandleOptions() {
	menu := DrawMenu()
out:
	for {
		ClearScreen()
		fmt.Print(menu)
		var userMenuOption int
		fmt.Print("Choose a options: ")
		fmt.Scan(&userMenuOption)

		switch userMenuOption {
		default:
			HandleOptions()
		case 0:
			fmt.Println("#TODO EXIT MENU")
			break out
		case 1:
			ClearScreen()
			ui.InsertOnDataBase()
		case 2:
			ClearScreen()
			ui.ViewOnDataBase()
		case 3:
			ClearScreen()
			ui.EditAUser()
			TimeSleep()
		case 4:
			ClearScreen()
			ui.DeleteAUser()
			TimeSleep()
		case 5:
			ClearScreen()
			ui.ViewAllContactsOnDB()
		}
	}
	ui.DrawDefaultScreen()
}
