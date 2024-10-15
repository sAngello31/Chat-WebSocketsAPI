package main

import (
	"main.go/displays"
	"main.go/utils"
)

func main() {
	utils.LoadEnv()
	utils.CleanScreen()
	utils.WriteTitle()
	displays.SelectionMainMenu()
}
