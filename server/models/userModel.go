package models

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	ContactNumber int32              `bson:"contact_number"`
	Name          string             `bson:"name"`
	LastName      string             `bson:"last_name"`
	Username      string             `bson:"username"`
	Password      string             `bson:"password"`
	CreatedAt     string             `bson:"created_at"`
}

type UserRepository struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{client: client}
}

func (ctrl *UserRepository) RegisterNewUser(c *gin.Context) {
	log.Println("Se va a registar un nuevo usuario :D")
	c.JSON(http.StatusOK, gin.H{
		"message": "Te voy a registar uwu",
	})
}

func (ctrl *UserRepository) Login(c *gin.Context) {
	log.Println("Se va a loggear un usuario :D")
	c.JSON(http.StatusOK, gin.H{
		"message": "Te voy a loggear",
	})
}

func (ctrl *UserRepository) GetAllUser(c *gin.Context) {
	var users []User
	collection := ctrl.client.Database(os.Getenv("NAME_DATABASE")).Collection(os.Getenv("NAME_USER_COLLECTION"))
	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error: Bad Query ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error: Bad Query"})
		return
	}
	if err = cursor.All(context.TODO(), &users); err != nil {
		log.Fatal("Error, Decode Error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al decodificar el documento MongoDB"})
		return
	}
	c.JSON(http.StatusOK, users)
}
