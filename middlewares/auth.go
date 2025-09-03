package middlewares

import (
	"net/http"
	"test-jwt-go/auth"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			ctx.Abort()
			return
		}
		err := auth.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"eror": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
