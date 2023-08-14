package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authority() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		expectedToken := "token"
		if authorizationHeader != "Bearer "+expectedToken {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			return
		}

		c.Next()
	}
}
