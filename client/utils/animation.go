package utils

import (
	"fmt"
	"time"
)

func Spinner(delay time.Duration) {
	chars := []rune{'|', '/', '-', '\\'}
	for {
		for _, char := range chars {
			fmt.Printf("\t\r%c Cargando...", char)
			time.Sleep(delay * 1000000)
		}
	}
}
