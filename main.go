package main

import (
	// Generic interface around SQL in general
	"github.com/odmrs/crud-phonebook/cmd/db"
	"github.com/odmrs/crud-phonebook/cmd/ui"
)

func main() {
	db.RunPostgresql()
	ui.HandleOptions()
}
