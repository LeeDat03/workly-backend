package postgres

import (
	"context"
	"database/sql"
	"workly-backend/internal/domain"

	"github.com/jmoiron/sqlx"
)

type UserDbRepository struct {
	db *sqlx.DB
}

func NewUserDbRepository(db *sqlx.DB) *UserDbRepository {
	return &UserDbRepository{
		db: db,
	}
}

func (r *UserDbRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (id, email, password, name, role, created_at, updated_at)
		VALUES (:id, :email, :password, :name, :role, :created_at, :updated_at)
	`
	_, err := r.db.NamedExecContext(ctx, query, user)
	return err
}

func (r *UserDbRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	query := `SELECT id, email, password, name, role, created_at, updated_at FROM users WHERE id = $1`
	user := &domain.User{}
	err := r.db.GetContext(ctx, user, query, id)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserDbRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, email, password, name, role, created_at, updated_at FROM users WHERE email = $1`
	user := &domain.User{}
	err := r.db.GetContext(ctx, user, query, email)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserDbRepository) Update(ctx context.Context, user *domain.User) error {
	query := `
		UPDATE users 
		SET email = :email, name = :name, role = :role, updated_at = :updated_at
		WHERE id = :id
	`
	result, err := r.db.NamedExecContext(ctx, query, user)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *UserDbRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return domain.ErrNotFound
	}

	return nil
}
