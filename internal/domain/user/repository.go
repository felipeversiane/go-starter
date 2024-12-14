package user

import (
	"context"
	"fmt"

	"github.com/felipeversiane/go-starter/internal/domain"
	"github.com/felipeversiane/go-starter/internal/infra/config/response"
	"github.com/felipeversiane/go-starter/internal/infra/database"
)

type userRepository struct {
	db database.DatabaseInterface
}

type UserRepositoryInterface interface {
	InsertOneRepository(domain domain.UserInterface, ctx context.Context) (string, *response.ErrorResponse)
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
