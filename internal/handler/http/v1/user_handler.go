package v1

import (
	"net/http"

	"workly-backend/internal/dto"
	"workly-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// TODO: Add validation
	// TODO: Call use case
	// user, err := h.userUseCase.CreateUser(c.Request.Context(), &req)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	// TODO: Call use case
	// user, err := h.userUseCase.GetUserByID(c.Request.Context(), id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Get user endpoint",
		"id":      id,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// TODO: Call use case
	// user, err := h.userUseCase.UpdateUser(c.Request.Context(), id, &req)

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"id":      id,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// TODO: Call use case
	// err := h.userUseCase.DeleteUser(c.Request.Context(), id)

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
		"id":      id,
	})
}
