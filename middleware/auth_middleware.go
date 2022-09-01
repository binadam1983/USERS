package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	utils "github.com/users/utils"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		_, err := utils.VerifyToken(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
		}
		c.Next()
	}

}
