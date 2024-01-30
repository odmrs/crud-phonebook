package ui

import (
	"fmt"
	"time"
)

func DrawDefaultScreen() {
	fmt.Println("---------- Hey, you type a not found option, try again ----------")
	time.Sleep(2 * time.Second)
}
