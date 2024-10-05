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
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	urlDB := "mongodb+srv://" + os.Getenv("USER_DB") + ":" + os.Getenv("PASSWORD_DB") + "@chatwebsocketapi.uivjm.mongodb.net/?retryWrites=true&w=majority&appName=" + os.Getenv("APP_NAME_DB")
	clOptions := options.Client().ApplyURI(urlDB).SetServerAPIOptions(serverAPI)
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
