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
	GetOneByIDService(id string, ctx context.Context) (*UserResponse, *response.ErrorResponse)
	GetOneByEmailService(email string, ctx context.Context) (*UserResponse, *response.ErrorResponse)
	GetAllService(ctx context.Context) (*[]UserResponse, *response.ErrorResponse)
	UpdateService(id string, req UserRequest, ctx context.Context) *response.ErrorResponse
	DeleteService(id string, ctx context.Context) *response.ErrorResponse
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

func (service *userService) GetOneByIDService(id string, ctx context.Context) (*UserResponse, *response.ErrorResponse) {
	user, err := service.repository.GetOneByIDRepository(id, ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *userService) GetOneByEmailService(email string, ctx context.Context) (*UserResponse, *response.ErrorResponse) {
	user, err := service.repository.GetOneByEmailRepository(email, ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *userService) GetAllService(ctx context.Context) (*[]UserResponse, *response.ErrorResponse) {
	users, err := service.repository.GetAllRepository(ctx)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (service *userService) UpdateService(id string, req UserRequest, ctx context.Context) *response.ErrorResponse {
	return nil
}

func (service *userService) DeleteService(id string, ctx context.Context) *response.ErrorResponse {
	return nil
}
