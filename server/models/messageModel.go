package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID   primitive.ObjectID
	To   string
	From string
	Msg  string
}
