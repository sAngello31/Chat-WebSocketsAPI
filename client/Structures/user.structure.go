package structures

type UserCredential struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
