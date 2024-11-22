package main

import (
	"chat_websocket/routes"
	"chat_websocket/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	utils.ConnectDB()
	defer utils.CloseDB()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run()
}
