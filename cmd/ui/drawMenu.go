package ui

import (
	"fmt"
	"os"
	"os/exec"
)

func DrawMenu() {
	var userMenuOption int
	ClearScreen()
	fmt.Println(`---------- PHONEBOOK ----------

Add Contact       -> 1
View Contact      -> 2
Edit Contact      -> 3
Delete Contact    -> 4
View all Contacts -> 5
Exit 		  -> 0

	`)

	fmt.Print("Make your choose: ")
	fmt.Scan(&userMenuOption)

}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
