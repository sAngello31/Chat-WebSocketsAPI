package main

import (
	"fmt"

	"main.go/displays"
	"main.go/utils"
)

func main() {
	utils.CleanScreen()
	fmt.Println("\n\n\t\tCLI WebSocketAPI")
	choice := displays.ShowMainMenu()
	action, ok := displays.MenuOptions[choice]
	if ok {
		action()
	} else {
		fmt.Println("Opcion Invalida")
	}
}
