package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/manifoldco/promptui"
)

var clear = make(map[string]func())

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CleanScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	}
}

func WriteTitle() {
	fmt.Println("\n\n\t\tCLI WebSocketAPI")
}

func InitDisplay(title string) {
	CleanScreen()
	fmt.Println("\n\t\t" + title)
}

func PrepareGetText(cred string) string {
	validate := func(input string) error {
		if input == "" {
			err := cred + " invalido"
			return errors.New(err)
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    cred,
		Validate: validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Println(cred + " invalido")
		return ""
	}
	return result
}
