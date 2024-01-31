package ui

import (
	"fmt"
	"os"

	ui "github.com/odmrs/crud-phonebook/cmd/ui/allFunctions"
)

func HandleOptions() {
	menu := DrawMenu()
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
			ClearScreen()
			ui.ExitMenu()
			TimeSleep()
			ClearScreen()
			os.Exit(0)
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
}
