package middlewares

import (
	"net/http"

	"example.com/lib"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	authToken := context.Request.Header.Get("Authorization")

	if authToken == "" {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	jwtAuthorizer := lib.NewJwtAuthorizer()

	userId, err := jwtAuthorizer.ValidateToken(authToken)

	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	context.Set("userId", userId)

	context.Next()
}
