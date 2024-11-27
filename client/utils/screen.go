package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func CleanScreen() {
	clear := make(map[string]func())

	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	}
}
