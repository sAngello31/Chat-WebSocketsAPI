package displays

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
	"main.go/services"
	"main.go/utils"
)

func ShowLoginMenu() {
	utils.InitDisplay("Iniciar Sesión")
	user := utils.PrepareGetText("Username")
	password := GetPassword()
	services.Login(user, password)
	utils.Spinner(100)
}

func GetPassword() string {
	validate := func(input string) error {
		if input == "" {
			return errors.New("password invalida")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Password",
		Validate: validate,
		Mask:     '*',
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Println("Password Invalida")
		return ""
	}
	return result
}
