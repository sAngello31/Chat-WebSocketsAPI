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
	Username      string             `bson:"username"`
	Password      string             `bson:"password"`
}

type UserRepository struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{client: client}
}

func (ctrl *UserRepository) getUserCollection() *mongo.Collection {
	return ctrl.client.Database(os.Getenv("NAME_DATABASE")).Collection(os.Getenv("NAME_USER_COLLECTION"))
}

func (ctrl *UserRepository) RegisterNewUser(c *gin.Context) {
	isUnique := ctrl.isUniqueUsername(c.PostForm("username"))
	if !isUnique {
		log.Println("Este username está ocupado. StatusCode: ", http.StatusConflict)
		c.String(http.StatusConflict, "El username está ocupado")
	}

}

func (ctrl *UserRepository) Login(c *gin.Context) {
	log.Println("Se va a loggear un usuario :D")
	c.JSON(http.StatusOK, gin.H{
		"message": "Te voy a loggear",
	})
}

func (ctrl *UserRepository) GetAllUser(c *gin.Context) []User {
	var users []User
	collection := ctrl.getUserCollection()
	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Println("Error: Bad Query ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error: Bad Query"})
		return nil
	}
	if err = cursor.All(context.TODO(), &users); err != nil {
		log.Println("Error, Decode Error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al decodificar el documento MongoDB"})
		return nil
	}
	//c.JSON(http.StatusOK, users)
	return users
}

// Considerar agregar un error como retorno para manejar otros errores de consultas
func (ctrl *UserRepository) isUniqueUsername(username string) bool {
	collection := ctrl.getUserCollection()
	filter := bson.M{"username": username}
	var user User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return true
		}
	}
	return false
}
