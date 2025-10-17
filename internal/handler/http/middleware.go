package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware logs HTTP requests using Gin
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

// CORSMiddleware handles CORS using Gin
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

// AuthMiddleware validates JWT tokens using Gin
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement JWT validation
		// token := c.GetHeader("Authorization")
		// Validate token and extract user info
		// Add user info to context: c.Set("userID", userID)

		c.Next()
	}
}

// RecoveryMiddleware recovers from panics using Gin
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

// ErrorHandlerMiddleware handles errors from handlers
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// Log the error
			log.Printf("Error: %v", err.Error())

			// If response hasn't been written yet
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
