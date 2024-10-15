package displays

import (
	"fmt"

	"main.go/utils"
)

func ShowRegisterMenu() {
	utils.InitDisplay("Registrarse")
	name := utils.PrepareGetText("Nombre")
	last_name := utils.PrepareGetText("Apellido")
	username := utils.PrepareGetText("Username")
	password := utils.PrepareGetText("Password")
	fmt.Println(name, last_name, username, password)
}
