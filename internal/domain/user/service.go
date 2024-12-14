package user

import (
	"context"

	"github.com/felipeversiane/go-starter/internal/infra/config/response"
)

type userService struct {
	repository UserRepositoryInterface
}

type UserServiceInterface interface {
	InsertOneService(req UserRequest, ctx context.Context) (string, *response.ErrorResponse)
}

func NewUserService(repository UserRepositoryInterface) UserServiceInterface {
	return &userService{repository}
}

func (service *userService) InsertOneService(req UserRequest, ctx context.Context) (string, *response.ErrorResponse) {
	domain, err := ConvertRequestToDomain(req)
	if err != nil {
		return "", err
	}
	id, err := service.repository.InsertOneRepository(domain, ctx)
	if err != nil {
		return "", err
	}
	return id, nil
}
