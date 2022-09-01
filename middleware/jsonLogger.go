package middleware

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	utils "github.com/users/utils"
)

// JSONLogMiddleware logs a gin HTTP request in JSON format, with some additional custom key/values
func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process the Next Request
		c.Next()

		// Stop timer
		duration := utils.GetDurationInMilliSeconds(start)

		log.SetFormatter(&log.JSONFormatter{})

		file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Warn("cant create logging file", err)
		}
		log.SetOutput(file)

		entry := log.WithFields(log.Fields{
			"client_ip":    c.ClientIP(),
			"duration(ms)": duration,
			"method":       c.Request.Method,
			"path":         c.Request.RequestURI,
			"status":       c.Writer.Status(),
			"referer":      c.Request.Referer(),
			"request_id":   c.Writer.Header().Get("X-Request-Id"),
		})

		if c.Writer.Status() >= 400 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}
