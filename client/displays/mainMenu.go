package displays

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

var MenuOptions = map[string]func(){
	"Iniciar Sesión":         ShowLoginMenu,
	"Registrarse":            func() {},
	"Salir de la Aplicación": func() {},
}

func ShowMainMenu() string {
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
	}
	return result
}
