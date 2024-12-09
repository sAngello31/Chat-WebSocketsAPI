package services

import (
	"bytes"
	modeldata "client_websockets/model_data"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// MEJORAR EL MANEJO DE ERRORES
func Login(user *modeldata.UserLogin) (int, string) {
	url_login := os.Getenv("URL_BACKEND") + "/auth/login"
	app_type := "application/x-www-form-urlencoded"
	formData := url.Values{
		"username": {user.Username},
		"password": {user.Password},
	}
	resp, err := http.Post(url_login, app_type, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var token Token
	err = json.Unmarshal(body, &token)
	if err != nil {
		panic(err)
	}
	return resp.StatusCode, token.Token
}

func Register(user *modeldata.UserLogin) int {
	client := &http.Client{}
	form_data := url.Values{
		"username": {user.Username},
		"password": {user.Password},
	}.Encode()
	req := makeAuthRequest("/auth/register", bytes.NewBufferString(form_data))
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

func GetUserData() *modeldata.User {
	url := os.Getenv("URL_BACKEND") + "/user/data"
	client := &http.Client{}
	var user modeldata.User
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	makeHeader(req)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(body, &user); err != nil {
		panic(err)
	}
	return &user

}

func GetAllUsers() {
	x, err := http.Get("url" + "/user/getAll")
	if err != nil {
		panic("Error al hacer la peticion")
	}
	println(x)
}

func makeAuthRequest(path string, data *bytes.Buffer) *http.Request {
	url := os.Getenv("URL_BACKEND") + path
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Go-TUI/1.0")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req
}

func makeHeader(req *http.Request) {
	token := "Bearer " + GetToken()
	req.Header.Set("User-Agent", "Go-TUI/1.0")
	req.Header.Set("Authorization", token)
	req.Header.Set("Accept", "application/json")
}
