package main

//1-888-420-9943
import (
	"chat_websocket/routes"
	"chat_websocket/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	utils.InitDB()
	defer utils.CloseDB()
	utils.InitRepositories()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run()
}
