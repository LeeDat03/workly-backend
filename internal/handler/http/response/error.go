package response

import (
	"errors"
	"net/http"

	"workly-backend/internal/domain"
	"workly-backend/pkg/logger"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}

func Error(c *gin.Context, err error) {
	status, response := MapError(err)

	logger.Error("Request error",
		"path", c.Request.URL.Path,
		"method", c.Request.Method,
		"error", err.Error(),
		"status", status,
	)

	c.JSON(status, response)
}

func MapError(err error) (int, ErrorResponse) {
	switch {
	case errors.Is(err, domain.ErrNotFound):
		return http.StatusNotFound, ErrorResponse{
			Error:   "Not Found",
			Message: "The requested resource was not found",
			Code:    "RESOURCE_NOT_FOUND",
		}

	case errors.Is(err, domain.ErrUserAlreadyExists):
		return http.StatusConflict, ErrorResponse{
			Error:   "Conflict",
			Message: "Email already exists",
			Code:    "EMAIL_EXISTS",
		}

	case errors.Is(err, domain.ErrInvalidCredentials):
		return http.StatusUnauthorized, ErrorResponse{
			Error:   "Unauthorized",
			Message: "Invalid email or password",
			Code:    "INVALID_CREDENTIALS",
		}

	case errors.Is(err, domain.ErrUnauthorized):
		return http.StatusUnauthorized, ErrorResponse{
			Error:   "Unauthorized",
			Message: "Authentication required",
			Code:    "UNAUTHORIZED",
		}

	case errors.Is(err, domain.ErrForbidden):
		return http.StatusForbidden, ErrorResponse{
			Error:   "Forbidden",
			Message: "You don't have permission to access this resource",
			Code:    "FORBIDDEN",
		}

	case errors.Is(err, domain.ErrInvalidInput):
		return http.StatusBadRequest, ErrorResponse{
			Error:   "Bad Request",
			Message: err.Error(),
			Code:    "INVALID_INPUT",
		}

	default:
		return http.StatusInternalServerError, ErrorResponse{
			Error:   "Internal Server Error",
			Message: "An unexpected error occurred",
			Code:    "INTERNAL_ERROR",
		}
	}
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, ErrorResponse{
		Error:   "Bad Request",
		Message: message,
		Code:    "BAD_REQUEST",
	})
}

func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, ErrorResponse{
		Error:   "Unauthorized",
		Message: message,
		Code:    "UNAUTHORIZED",
	})
}

func Forbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, ErrorResponse{
		Error:   "Forbidden",
		Message: message,
		Code:    "FORBIDDEN",
	})
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, ErrorResponse{
		Error:   "Not Found",
		Message: message,
		Code:    "NOT_FOUND",
	})
}

func InternalError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Error:   "Internal Server Error",
		Message: message,
		Code:    "INTERNAL_ERROR",
	})
}
