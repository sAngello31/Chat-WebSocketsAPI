package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageToSend struct {
	ID               primitive.ObjectID
	To               primitive.ObjectID
	From             primitive.ObjectID
	Msg              string
	Datetime_sending string
}

type MsgRepository struct {
	client *mongo.Client
}

func NewMsgRepository(client *mongo.Client) *MsgRepository {
	return &MsgRepository{client: client}
}
