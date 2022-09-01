package middleware

import (
	"github.com/gin-gonic/gin"
)

type CORSOptions struct {
	Origin string
}

func CORS(options CORSOptions) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		if options.Origin != "" {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", "options.Origin")
		}
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Cotent-Type")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
