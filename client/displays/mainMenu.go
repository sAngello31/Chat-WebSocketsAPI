package displays

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"main.go/utils"
)

var MenuOptions = map[string]func(){
	"Iniciar Sesión":         ShowLoginMenu,
	"Registrarse":            ShowRegisterMenu,
	"Salir de la Aplicación": SalirApp,
}

func ShowMainMenu() (string, error) {
	options := make([]string, 0, len(MenuOptions))
	for option := range MenuOptions {
		options = append(options, option)
	}
	prompt := promptui.Select{
		Label: "Menú Principal",
		Items: options,
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Error al seleccionar la opción")
		return "", err
	}
	return result, nil
}

func SelectionMainMenu() {
	choice, err := ShowMainMenu()
	if err != nil {
		fmt.Println("Error al seleccionar: ", err)
		return

	}
	action, ok := MenuOptions[choice]
	if ok {
		action()
	}
}

func SalirApp() {
	utils.CleanScreen()
	fmt.Println("\nNos vemos! :D")
	os.Exit(0)
}
