package models

import (
	"chat_websocket/services"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	ContactNumber int64              `bson:"contact_number"`
	Username      string             `bson:"username"`
	Password      string             `bson:"password"`
}

type UserToSave struct {
	ContactNumber int64
	Name          string
	LastName      string
	Username      string
	Password      string
	CreatedAt     string
}

type UserRepository struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{client: client}
}

// Refactorizar (dividir la logica)
func (ctrl *UserRepository) RegisterNewUser(c *gin.Context) {
	isUnique := ctrl.IsUniqueUsername(c.PostForm("username"))
	if !isUnique {
		log.Println("Este username está ocupado. StatusCode: ", http.StatusConflict)
		c.String(http.StatusConflict, "El username está ocupado") // Cambiar de acuerdo al cliente
		return
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error: Hashing error ", err)
		c.String(http.StatusConflict, "Error al cifrar la contraseña")
		return
	}
	userForSaving := UserToSave{
		ContactNumber: services.CreateRandomUserCode(),
		Name:          c.PostForm("name"),
		LastName:      c.PostForm("last_name"),
		Username:      c.PostForm("username"),
		Password:      string(passwordHashed),
		CreatedAt:     time.Now().Format(time.RFC3339),
	}
	result := ctrl.saveUser(userForSaving)
	if result != nil {
		log.Println("Nuevo Usuario Creado Correctamente: ", result)
		c.JSON(http.StatusCreated, result)
		return
	}
	c.String(http.StatusConflict, "Error al crear el nuevo usuario")
}

func (ctrl *UserRepository) saveUser(user UserToSave) *mongo.InsertOneResult {
	collection := ctrl.getUserCollection()
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("Error al guardar el nuevo usuario", err)
	}
	return result
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
	return users
}

// Considerar agregar un error como retorno para manejar otros errores de consultas
func (ctrl *UserRepository) IsUniqueUsername(username string) bool {
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

func (ctrl *UserRepository) getUserCollection() *mongo.Collection {
	return ctrl.client.Database(os.Getenv("NAME_DATABASE")).Collection(os.Getenv("NAME_USER_COLLECTION"))
}
