package models

type Message struct {
	To      string `json:"to" bson:"to"`
	From    string `json:"from" bson:"from"`
	Content string `json:"content" bson:"content"`
	UUID    string
}
