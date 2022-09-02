package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	utils "github.com/users/utils"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		user, err := utils.VerifyToken(c)
		if err != nil || user == "" {
			c.String(http.StatusUnauthorized, err.Error())
			log.Info(user, "User not logged in, please login first")
			//calling "Abort" to make sure successive handlers are not called
			c.Abort()
			return
		}
		log.Info("User Valid ", user)
		// allowing for the successive handlers to be called
		c.Next()
	}

}
