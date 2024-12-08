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

func GetAllUsers() {
	x, err := http.Get("url" + "/user/getAll")
	if err != nil {
		panic("Error al hacer la peticion")
	}
	println(x)
}
