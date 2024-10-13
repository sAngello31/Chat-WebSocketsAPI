package displays

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

func ShowLoginMenu() {
	user := GetUsername()
	password := GetPassword()
	fmt.Println(user, password)
}

func GetUsername() string {
	validate := func(input string) error {
		if input == "" {
			return errors.New("username invalido")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Username",
		Validate: validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Println("Username invalido")
		return ""
	}
	return result
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
