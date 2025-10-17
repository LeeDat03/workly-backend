package domain

import (
	"context"
	"time"
)

type User struct {
	ID        string
	Email     string
	Password  string
	Name      string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserRepository defines the interface for user data access
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
}
