package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ilhampraset/testcasebe/auth/utils"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		tokenString := strings.SplitAfter(authHeader, "Bearer")[1]
		token, _ := utils.JWTAuth().ValidateToken(strings.TrimSpace(tokenString))

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unathorized"})

		}
		c.Next()
	}
}
