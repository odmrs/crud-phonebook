package ui

import (
	"os"
	"os/exec"
	"time"
)

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func TimeSleep() {
	time.Sleep(2 * time.Second)
}
