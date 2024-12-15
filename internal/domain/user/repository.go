package user

import (
	"context"
	"fmt"

	"github.com/felipeversiane/go-starter/internal/domain"
	"github.com/felipeversiane/go-starter/internal/infra/config/response"
	"github.com/felipeversiane/go-starter/internal/infra/database"
	"github.com/jackc/pgx/v5"
)

type userRepository struct {
	db database.DatabaseInterface
}

type UserRepositoryInterface interface {
	InsertOneRepository(domain domain.UserInterface, ctx context.Context) (string, *response.ErrorResponse)
	GetOneByIDRepository(id string, ctx context.Context) (*UserResponse, *response.ErrorResponse)
	GetOneByEmailRepository(id string, ctx context.Context) (*UserResponse, *response.ErrorResponse)
	GetOneAllRepository(ctx context.Context) ([]UserResponse, *response.ErrorResponse)
	UpdateRepository(id string, domain domain.UserInterface, ctx context.Context) *response.ErrorResponse
	DeleteRepository(id string, ctx context.Context) *response.ErrorResponse
}

func NewUserRepository(db database.DatabaseInterface) UserRepositoryInterface {
	return &userRepository{db}
}

func (repository *userRepository) InsertOneRepository(domain domain.UserInterface, ctx context.Context) (string, *response.ErrorResponse) {
	search := fmt.Sprintf("%s %s %s", domain.GetEmail(), domain.GetFirstName(), domain.GetLastName())
	query := `INSERT INTO users (id, email, first_name, last_name, password, search, created_at, updated_at, deleted) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	args := []interface{}{
		domain.GetID(),
		domain.GetEmail(),
		domain.GetFirstName(),
		domain.GetLastName(),
		domain.GetPassword(),
		search,
		domain.GetCreatedAt(),
		domain.GetUpdatedAt(),
		domain.GetDeleted(),
	}
	var id string
	err := repository.db.GetDB().QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return "", response.NewInternalServerError(fmt.Sprintf("Unable to insert user: %v", err))
	}
	return id, nil
}

func (repository *userRepository) GetOneByIDRepository(id string, ctx context.Context) (*UserResponse, *response.ErrorResponse) {
	query := `
		SELECT id, email, first_name, last_name, created_at, updated_at
		FROM users
		WHERE id = $1 AND deleted = false`

	var user UserResponse
	err := repository.db.GetDB().QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, response.NewNotFoundError("User not found")
		}
		return nil, response.NewInternalServerError(fmt.Sprintf("Error querying user by ID: %v", err))
	}

	return &user, nil

}

func (repository *userRepository) GetOneByEmailRepository(id string, ctx context.Context) (*UserResponse, *response.ErrorResponse) {
	return nil, nil
}

func (repository *userRepository) GetOneAllRepository(ctx context.Context) ([]UserResponse, *response.ErrorResponse) {
	return nil, nil
}

func (repository *userRepository) UpdateRepository(id string, domain domain.UserInterface, ctx context.Context) *response.ErrorResponse {
	return nil
}

func (repository *userRepository) DeleteRepository(id string, ctx context.Context) *response.ErrorResponse {
	return nil
}
