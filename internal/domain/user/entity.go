package user

import (
	"time"

	"github.com/felipeversiane/go-starter/internal/domain"
	"github.com/felipeversiane/go-starter/internal/infra/config/response"
)

type UserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	FirstName string `json:"first_name" binding:"required,min=2,max=50"`
	LastName  string `json:"last_name" binding:"required,min=2,max=50"`
	Password  string `json:"password" binding:"required,min=8"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ConvertRequestToDomain(req UserRequest) (domain.UserInterface, *response.ErrorResponse) {
	domain, err := domain.NewUser(req.Email, req.FirstName, req.LastName, req.Password)
	if err != nil {
		return nil, response.NewBadRequestError(err.Error())
	}
	return domain, nil
}

func ConvertDomainToResponse(user domain.UserInterface) *UserResponse {
	return &UserResponse{
		ID:        user.GetID(),
		Email:     user.GetEmail(),
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
		CreatedAt: user.GetCreatedAt(),
		UpdatedAt: user.GetUpdatedAt(),
	}
}
