package models

type Message struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Content string `json:"content"`
	UUID    string
}
