package usecase

import (
	"context"
	"fmt"
	"time"
	"workly-backend/internal/domain"
	"workly-backend/internal/dto"
	"workly-backend/pkg/utils"
)

type AuthUseCase struct {
	userRepo domain.UserRepository
}

func NewAuthUseCase(userRepo domain.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		userRepo: userRepo,
	}
}

func (uc *AuthUseCase) Register(ctx context.Context, req *dto.RegisterRequest) (*domain.User, error) {
	// // Check if user already exists
	existingUser, err := uc.userRepo.GetByEmail(ctx, req.Email)
	if err != nil && err != domain.ErrNotFound {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}
	if existingUser != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	// // Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	now := time.Now()
	user := &domain.User{
		Email:     req.Email,
		Password:  hashedPassword,
		Name:      req.Name,
		Role:      "user",
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Save to repository
	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}
