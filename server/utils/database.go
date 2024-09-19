package utils

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitDB() {
	var err error
	ctx := context.TODO()
	clOptions := options.Client().ApplyURI(os.Getenv("HOST_DATABASE"))
	Client, err = mongo.Connect(ctx, clOptions)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database is connected")
}

func CloseDB() {
	ctx := context.TODO()
	err := Client.Disconnect(ctx)
	if err != nil {
		log.Fatal("Error: La conexión no se pudo cerrar", err)
	}
}
