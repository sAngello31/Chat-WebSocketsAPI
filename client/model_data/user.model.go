package modeldata

import (
	"strconv"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username      string `json:"username"`
	ContactNumber int    `json:"contact_number"`
}

func (u User) Title() string {
	return u.Username
}

func (u User) Description() string {
	return strconv.Itoa(u.ContactNumber)
}

func (u User) FilterValue() string {
	return strconv.Itoa(u.ContactNumber)
}
