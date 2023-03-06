package utils

import (
	"fmt"
	"time"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Cyan = "\033[36m"
var Gray = "\033[37m"

func Log(text string) {
	now := time.Now()

	fmt.Print(Green + "INFO " + Reset)

	fmt.Print(" [")
	fmt.Print(now.Format("01-02-2006 15:04:05"))
	fmt.Print("] ")

	fmt.Println(Cyan + text + Reset)
}

func Error(err error) {
	now := time.Now()

	fmt.Print(Red + "ERROR" + Reset)

	fmt.Print(Yellow + " [")
	fmt.Print(now.Format("01-02-2006 15:04:05"))
	fmt.Print("] " + Reset)

	fmt.Println(Gray + err.Error() + Reset)
}
