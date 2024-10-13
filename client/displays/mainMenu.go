package displays

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

var menuOptions = []string{"Iniciar Sesión", "Registrarse", "Salir de la Aplicación"}

func ShowMainMenu() string {
	prompt := promptui.Select{
		Label: "Menú Principal",
		Items: menuOptions,
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Error al seleccionar la opción")
	}
	return result
}
