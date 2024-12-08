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

	token, err := services.ValidJWT(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	data := services.GetUserData(token)
	c.Set("user_id", data["user_id"])
	c.Set("username", data["username"])
	c.Set("contact_number", data["contact_number"])
	c.Next()
}
