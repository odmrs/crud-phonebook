package ui

import (
	"fmt"
)

func DrawMenu() {
	ClearScreen()
	fmt.Println(`---------- PHONEBOOK ----------

Add Contact       -> 1
View Contact      -> 2
Edit Contact      -> 3
Delete Contact    -> 4
View all Contacts -> 5
Exit 		  -> 0

	`)
	HandleOptions()
}
