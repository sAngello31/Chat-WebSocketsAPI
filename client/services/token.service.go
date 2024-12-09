package services

import (
	"bufio"
	"os"
	"strconv"
)

type Token struct {
	Token string `json:"Access_Token"`
}

func SaveToken(token string) {
	pid := os.Getpid()
	file_name := "tmp/token_" + strconv.Itoa(pid)
	file, err := os.Create(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(token)
	if err != nil {
		panic(err)
	}
}

func GetToken() string {
	pid := os.Getpid()
	file_name := "tmp/token_" + strconv.Itoa(pid)
	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text()
	}

	return ""
}
