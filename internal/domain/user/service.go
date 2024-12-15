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
	GetOneByIDService(string, ctx context.Context) (*UserResponse, *response.ErrorResponse)
	GetOneByEmailService(string, ctx context.Context) (*UserResponse, *response.ErrorResponse)
	GetAllService(ctx context.Context) (*[]UserResponse, *response.ErrorResponse)
	UpdateService(string, req UserRequest, ctx context.Context) *response.ErrorResponse
	DeleteService(string, ctx context.Context) *response.ErrorResponse
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

func (service *userService) GetOneByIDService(string, ctx context.Context) (*UserResponse, *response.ErrorResponse) {
	return nil, nil
}

func (service *userService) GetOneByEmailService(string, ctx context.Context) (*UserResponse, *response.ErrorResponse) {
	return nil, nil
}

func (service *userService) GetAllService(ctx context.Context) (*[]UserResponse, *response.ErrorResponse) {
	return nil, nil
}

func (service *userService) UpdateService(string, req UserRequest, ctx context.Context) *response.ErrorResponse {
	return nil
}

func (service *userService) DeleteService(string, ctx context.Context) *response.ErrorResponse {
	return nil
}
