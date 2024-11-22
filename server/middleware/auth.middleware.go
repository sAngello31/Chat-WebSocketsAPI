package middleware

import (
	"chat_websocket/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func pathIsContain(path string) bool {
	exclusives_paths := []string{
		"/",
		"/auth/login",
		"/auth/register",
	}
	for _, i := range exclusives_paths {
		if i == path {
			log.Println(i)
			return true
		}
	}
	return false
}

func JWTMiddleware(c *gin.Context) {
	path := c.Request.URL.Path
	log.Println(path)
	isContained := pathIsContain(path)
	if isContained {
		c.Next()
		return
	}
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Error con la cabecera en la petición"})
		c.Abort()
		return
	}
	token_id, err := services.GetObjectIDFromJWT(authHeader) // Cambiar por la logica de negocios, que no sea un ObjectID como se maneje la aplicación
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	c.Set("user_id", token_id)
	c.Next()
}
