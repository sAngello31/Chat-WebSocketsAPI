package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	structures "main.go/Structures"
)

func Login(username string, password string) {
	url := os.Getenv("BACKEND_URL") + "/auth/login/"
	creds := structures.UserCredential{Username: username, Password: password}
	data, err := json.Marshal(creds)
	if err != nil {
		fmt.Println("Error al preparar los datos: ", err)
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	fmt.Println(resp)
	if err != nil {
		fmt.Println("Error al presentar la solicitud: ", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {
		fmt.Println(string(body))
	} else {
		fmt.Println("Error al autenticar: ", string(body))
	}
}
