package main

import (
	"client_websockets/models"
	"client_websockets/utils"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	utils.CleanScreen()
	p := tea.NewProgram(models.InitModel())
	_, err := p.Run()
	if err != nil {
		fmt.Printf("there's been an error: %v", err)
		os.Exit(1)
	}
}
