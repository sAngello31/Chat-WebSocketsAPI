package services

import (
	"chat_websocket/models"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateJWT(user models.User) (string, error) {
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":             user.ID.Hex(),
		"username":       user.Username,
		"contact_number": user.ContactNumber,
		"exp":            time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("STRING_TOKEN")))
	if err != nil {
		return "nil", err
	}
	return token, nil
}

func ValidJWT(token string) (jwt.Token, error) {
	if token == "" {
		return jwt.Token{}, fmt.Errorf("error: missing JWT")
	}
	tokenString := strings.TrimPrefix(token, "Bearer ")
	realToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("STRING_TOKEN")), nil
	})

	if err != nil || !realToken.Valid {
		return *realToken, fmt.Errorf("unexpected signing method: %v", realToken.Header["alg"])
	}
	return *realToken, nil
}

func GetIDFromJWT(tokenString string) (string, error) {
	token, err := ValidJWT(tokenString)
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["id"].(string), nil
	}

	return "", fmt.Errorf("error: Token expirado o invalido")
}

func GetObjectIDFromJWT(tokenString string) (primitive.ObjectID, error) {
	token, err := GetIDFromJWT(tokenString)
	if err != nil {
		return primitive.NilObjectID, err
	}
	objectToken, err := primitive.ObjectIDFromHex(token)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return objectToken, nil
}
