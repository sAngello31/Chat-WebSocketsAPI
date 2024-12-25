package models

import (
	"crypto/rand"
	"log"
	"math/big"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ContactNumber int64              `bson:"contact_number" json:"contact_number"`
	Username      string             `bson:"username" json:"username"`
	Password      string             `bson:"password" json:"-"`
}

func CreateUser(c *gin.Context, password string) User {
	return User{
		ContactNumber: createRandUserCode(),
		Username:      c.PostForm("username"),
		Password:      password,
	}
}

func createRandUserCode() int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		log.Println("Error: Seed does not generate")
		return 0
	}
	return nBig.Int64() + 100000
}
