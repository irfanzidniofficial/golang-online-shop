package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: get header authorization

		key := os.Getenv("ADMIN_SECRET")

		// TODO: validate header according to admin password
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.JSON(401, gin.H{"error": "Access Unauthorized"})
			c.Abort()
			return
		}

		if auth != key {
			c.JSON(401, gin.H{"error": "Access Unauthorized"})
			c.Abort()
			return
		}
		// TODO: next request to handler
		c.Next()

	}
}
