package v1

import (
	"workly-backend/internal/dto"
	"workly-backend/internal/handler/http/response"
	"workly-backend/internal/usecase"
	"workly-backend/pkg/logger"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUseCase *usecase.AuthUseCase
}

func NewAuthHandler(authUseCase *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var req dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	user, err := h.authUseCase.Register(ctx.Request.Context(), &req)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	logger.Info("User", user)

	userResponse := &dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	response.CreatedWithMessage(ctx, userResponse, "User registered successfully")
}
