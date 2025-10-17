package usecase

import (
	"context"

	"workly-backend/internal/domain"
	"workly-backend/internal/dto"
)

type UserUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*domain.User, error) {
	// TODO: Implement user creation logic
	// - Validate input
	// - Hash password
	// - Create user in repository
	return nil, nil
}

func (uc *UserUseCase) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	// TODO: Implement get user logic
	return uc.userRepo.GetByID(ctx, id)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, id string, req *dto.UpdateUserRequest) (*domain.User, error) {
	// TODO: Implement update user logic
	return nil, nil
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id string) error {
	// TODO: Implement delete user logic
	return uc.userRepo.Delete(ctx, id)
}
