package models

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatGroup struct {
	Name       string `bson:"name"`
	Created_at string `bson:"created_at"`
}

type ChatGroupRepository struct {
	client *mongo.Client
}

func NewChatGroupRepository(client *mongo.Client) *ChatGroupRepository {
	return &ChatGroupRepository{client: client}
}

func (ctrl *ChatGroupRepository) GetAllChatsGroups(c *gin.Context) {
	collection := ctrl.GetChatCollection()
	var chats []ChatGroup
	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer el documento MongoDB"})
		return
	}
	if err = cursor.All(context.TODO(), &chats); err != nil {
		log.Println("Error, Decode Error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al decodificar el documento MongoDB"})
		return
	}
	c.JSON(http.StatusOK, chats)
}

func (ctrl *ChatGroupRepository) InsertChatGroup(c *gin.Context) {

}

func (ctrl *ChatGroupRepository) GetChatCollection() *mongo.Collection {
	return ctrl.client.Database(os.Getenv("NAME_DATABASE")).Collection(os.Getenv("NAME_CHATS_COLLECTION"))
}
