package handler

import (
	"fmt"
	"goclean/utils/authutils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorize",
			})
			c.Abort()
			return
		}
		tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
		fmt.Println(tokenString)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorize",
			})
			c.Abort()
			return
		}
		token, err := authutils.VerifyAccessToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorize",
			})
			c.Abort()
			return
		}
		fmt.Println(token)
		if token != "" {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorize",
			})
			c.Abort()
			return
		}
	}
}
