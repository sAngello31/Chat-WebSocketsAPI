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
)

func GetAllUsers(c *gin.Context) {
	collection := utils.GetCollection("users")
	filter := bson.M{"username": bson.M{"$ne": c.Keys["username"]}}
	var userList []models.User
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	if err = cursor.All(context.TODO(), &userList); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, userList)
}

func GetUserData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"username":       c.Keys["username"],
		"contact_number": c.Keys["contact_number"],
	})
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
