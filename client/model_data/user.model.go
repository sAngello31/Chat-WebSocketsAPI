package modeldata

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username      string `json:"Username"`
	ContactNumber int    `json:"ContactNumber"`
}
