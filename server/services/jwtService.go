package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateJWT(user_id primitive.ObjectID) (string, error) {
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user_id.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("STRING_TOKEN")))
	if err != nil {
		return "nil", err
	}
	return token, nil
}
