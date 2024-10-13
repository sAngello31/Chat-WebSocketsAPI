package utils

import (
	"os"
	"os/exec"
	"runtime"
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
