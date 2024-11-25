package models

import (
	"crypto/rand"
	"log"
	"math/big"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	ContactNumber int64              `bson:"contact_number"`
	Username      string             `bson:"username"`
	Password      string             `bson:"password"`
}

type UserRepository struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{client: client}
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

/*

// ------------- Register New User -------------------------
// Refactorizar (dividir la logica)

// --------------------- Login User -------------------------
func (ctrl *UserRepository) Login(c *gin.Context) {
	user, err := ctrl.GetUserByUsername(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error al encontrar el usuario",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.PostForm("password")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Contraseña incorrecta",
		})
		return
	}

	token, err := services.GenerateJWT(user.ID)
	if err != nil {
		log.Println("Error al generar el JWT", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error al generar JWT",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"token": token,
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

// ----------------------- Operation with Database -------------------------

func (ctrl *UserRepository) saveUser(user UserToSave) *mongo.InsertOneResult {
	collection := ctrl.getUserCollection()
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("Error al guardar el nuevo usuario", err)
	}
	return result
}

func (ctrl *UserRepository) GetUserByUsername(c *gin.Context) (User, error) {
	collection := ctrl.getUserCollection()
	filter := bson.M{"username": c.PostForm("username")}
	var user User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ctrl *UserRepository) getUserCollection() *mongo.Collection {
	return ctrl.client.Database(os.Getenv("NAME_DATABASE")).Collection(os.Getenv("NAME_USER_COLLECTION"))
}

// REFACTORING


func IsUniqueUsername(username string) bool {
	collection := utils.GetCollection("users")
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
*/
