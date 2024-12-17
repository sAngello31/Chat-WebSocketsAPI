package services

import (
	"bytes"
	modeldata "client_websockets/model_data"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

func Login(user *modeldata.UserLogin) (int, string) {
	client := &http.Client{}
	formData := url.Values{
		"username": {user.Username},
		"password": {user.Password},
	}.Encode()
	req := makeAuthRequest("/auth/login", bytes.NewBufferString(formData))
	resp, err := client.Do(req)
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
	setHeader(req)
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

func GetAllUsers() *[]modeldata.User {
	var users []modeldata.User
	client := &http.Client{}
	url := os.Getenv("URL_BACKEND") + "/user/getAll"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	setHeader(req)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(body, &users); err != nil {
		panic(err)
	}
	return &users
}

func ConnectChat(userA, userB string) *websocket.Conn {
	url := os.Getenv("URL_CHAT") + GenerateUUID(userA, userB)
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial(url, *makeHeader())
	if err != nil {
		panic(err)
	}
	return conn
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

func makeHeader() *http.Header {
	token := "Bearer " + GetToken()
	header := http.Header{}
	header.Set("User-Agent", "Go-TUI/1.0")
	header.Set("Authorization", token)
	header.Set("Accept", "application/json")
	return &header
}

func setHeader(req *http.Request) {
	req.Header = *makeHeader()
}
