package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageToSend struct {
	To               primitive.ObjectID
	From             primitive.ObjectID
	Msg              string
	Datetime_sending string
}

type Message struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Content string `json:"content"`
}
