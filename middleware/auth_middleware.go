package middleware

import (
	"go-gin/errorhandler"
	"go-gin/helper"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			errorhandler.HandlerError(c, &errorhandler.UnauthorizedError{
				Message: "Unauthorized",
			})
			c.Abort()
			return
		}

		userId, err := helper.ValidateToken(tokenString)
		if err != nil {
			errorhandler.HandlerError(c, &errorhandler.UnauthorizedError{
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("userID", *userId)
		c.Next()
	}

}
