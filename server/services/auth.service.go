package services

import (
	"chat_websocket/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Falta integrarlo, se hará cuando se tenga internet xd

func RUser(c *gin.Context) {
	validUsername := isUniqueUsername(c.PostForm("username"))
	if !validUsername {
		msg := "Username is already taken"
		log.Println(msg)
		c.JSON(http.StatusBadRequest, gin.H{"messge": msg})
		return
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	newUser := models.CreateUser(c, string(passwordHashed))
	result, err := saveUser(&newUser)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"InsertedID": result.InsertedID,
	})
}

func LUser(c *gin.Context) {
	user, err := getUserByUsername(c.PostForm("username"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.PostForm("password")))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	token, err := GenerateJWT(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Access_Token": token})
}
