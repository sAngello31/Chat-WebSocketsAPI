package main

import (
	"chat_websocket/middleware"
	"chat_websocket/routes"
	"chat_websocket/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	utils.ConnectDB()
	defer utils.CloseDB()
	r := gin.Default()
	r.Use(middleware.JWTMiddleware)
	routes.SetupRoutes(r)
	r.Run()
}
