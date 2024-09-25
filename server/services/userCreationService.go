package services

import (
	"crypto/rand"
	"log"
	"math/big"
)

/*

func CreateUserToSave(ctrl *models.UserRepository, c *gin.Context) (models.UserToSave, error) {
	isUserUnique := ctrl.IsUniqueUsername(c.PostForm("username"))
	if !isUserUnique {
		log.Println("Este username está ocupado. StatusCode: ", http.StatusConflict)
		return models.UserToSave{}, errors.New("error al validar el username")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error al encriptar la contraseña")
		return models.UserToSave{}, errors.New("error al encriptar la contraseña")
	}

	user := models.UserToSave{
		ContactNumber: 12345,
		Name:          c.PostForm("name"),
		LastName:      c.PostForm("last_name"),
		Username:      c.PostForm("username"),
		Password:      string(password),
		CreatedAt:     "Hoy mijito uwu",
	}

	return user, nil
}
*/

func CreateRandomUserCode() int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		log.Println("Error al generar el seed")
		return 0
	}
	code := nBig.Int64() + 100000
	return code
}
