package http

import (
	"net/http"

	"workly-backend/internal/container"
	v1 "workly-backend/internal/handler/http/v1"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(c *container.Container) *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware(c.Config.FeHost))
	// router.Use(RecoveryMiddleware())

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "Server is running",
			})
		})
		setupAuthRoutes(apiV1, c)
		setupUserRoutes(apiV1, c)
	}

	return router
}

func setupAuthRoutes(rg *gin.RouterGroup, c *container.Container) {
	handler := v1.NewAuthHandler(c.AuthUseCase)

	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/register", handler.Register)
		// TODO: Add more auth routes
		// authGroup.POST("/login", handler.Login)
		// authGroup.POST("/logout", handler.Logout)
		// authGroup.POST("/refresh", handler.RefreshToken)
	}
}

func setupUserRoutes(rg *gin.RouterGroup, c *container.Container) {
	handler := v1.NewUserHandler(c.UserUseCase)

	userGroup := rg.Group("/users")
	{
		userGroup.POST("", handler.CreateUser)
		userGroup.GET("/:id", handler.GetUser)
		userGroup.PUT("/:id", handler.UpdateUser)
		userGroup.DELETE("/:id", handler.DeleteUser)
	}
}
