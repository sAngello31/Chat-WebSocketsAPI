package services

import (
	"chat_websocket/models"
	"chat_websocket/utils"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	validUsername := isUniqueUsername(c.PostForm("username"))
	if !validUsername {
		log.Println("Username is already taken")
		c.JSON(http.StatusBadRequest, gin.H{"messge": "Username is already taken"})
		return
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error in hashing password")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	newUser := models.CreateUser(c, string(passwordHashed))
	result, err := saveUser(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"StatusCode": http.StatusInternalServerError,
			"Message":    "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"StatusCode": http.StatusOK,
		"InsertedID": result.InsertedID,
	})
}

func LoginUser(c *gin.Context) {
	user, err := getUserByUsername(c.PostForm("username"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.PostForm("password")))
	if err != nil {
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

func saveUser(user *models.User) (*mongo.InsertOneResult, error) {
	collection := utils.GetCollection("users")
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func getUserByUsername(username string) (models.User, error) {
	collection := utils.GetCollection("users")
	var user models.User
	filter := bson.M{"username": username}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Println(err)
		return user, err
	}
	return user, nil
}

func isUniqueUsername(username string) bool {
	collection := utils.GetCollection("users")
	filter := bson.M{"username": username}
	count, err := collection.CountDocuments(context.TODO(), filter)

	if err != nil {
		log.Println(err)
		return false
	}

	return count == 0
}
