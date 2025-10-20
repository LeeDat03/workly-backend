package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Log request
		log.Printf("Started %s %s", c.Request.Method, c.Request.URL.Path)

		// Process request
		c.Next()

		// Log response
		log.Printf("Completed %s %s in %v with status %d",
			c.Request.Method,
			c.Request.URL.Path,
			time.Since(start),
			c.Writer.Status(),
		)
	}
}

func CORSMiddleware(feHost string) gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowOrigins:     []string{feHost},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
		},
	)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement JWT validation
		// token := c.GetHeader("Authorization")
		// Validate token and extract user info
		// Add user info to context: c.Set("userID", userID)

		c.Next()
	}
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Internal Server Error",
					"message": "An unexpected error occurred",
					"code":    "PANIC_RECOVERED",
				})
				c.Abort()
			}
		}()

		c.Next()
	}
}

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			log.Printf("Error: %v", err.Error())

			if !c.Writer.Written() {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Internal Server Error",
					"message": err.Error(),
					"code":    "HANDLER_ERROR",
				})
			}
		}
	}
}
