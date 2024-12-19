package user

import (
	"context"
	"net/mail"

	"github.com/felipeversiane/go-starter/internal/infra/config/response"
	"github.com/google/uuid"
)

var validationErr *response.ErrorResponse

type userService struct {
	repository UserRepositoryInterface
}

type UserServiceInterface interface {
	InsertOneService(req UserRequest, ctx context.Context) (string, *response.ErrorResponse)
	GetOneByIDService(id string, ctx context.Context) (*UserResponse, *response.ErrorResponse)
	GetOneByEmailService(email string, ctx context.Context) (*UserResponse, *response.ErrorResponse)
	GetAllService(ctx context.Context) (*[]UserResponse, *response.ErrorResponse)
	UpdateService(id string, req UserUpdateRequest, ctx context.Context) *response.ErrorResponse
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
	user, err := service.GetOneByEmailService(domain.GetEmail(), ctx)
	if err == nil && user != nil {
		return "", response.NewBadRequestError("User with this email already exists")
	}
	id, err := service.repository.InsertOneRepository(domain, ctx)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (service *userService) GetOneByIDService(id string, ctx context.Context) (*UserResponse, *response.ErrorResponse) {
	_, parseErr := uuid.Parse(id)
	if parseErr != nil {
		validationErr = response.NewBadRequestError("Invalid ID format")
		return nil, validationErr
	}
	user, err := service.repository.GetOneByIDRepository(id, ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *userService) GetOneByEmailService(email string, ctx context.Context) (*UserResponse, *response.ErrorResponse) {

	if email == "" {
		validationErr = response.NewBadRequestError("Email is required")
		return nil, validationErr
	}
	if _, err := mail.ParseAddress(email); err != nil {
		validationErr = response.NewBadRequestError("Invalid email format")
		return nil, validationErr
	}
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

func (service *userService) UpdateService(id string, req UserUpdateRequest, ctx context.Context) *response.ErrorResponse {
	if id == "" {
		validationErr = response.NewBadRequestError("ID is required")
		return validationErr
	}
	_, parseErr := uuid.Parse(id)
	if parseErr != nil {
		validationErr = response.NewBadRequestError("Invalid ID format")
		return validationErr
	}
	domain := ConvertUpdateRequestToDomain(req)
	user, err := service.GetOneByIDService(id, ctx)
	if err != nil && user == nil {
		return response.NewBadRequestError("User do not exists")
	}
	err = service.repository.UpdateRepository(id, domain, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (service *userService) DeleteService(id string, ctx context.Context) *response.ErrorResponse {
	if id == "" {
		validationErr = response.NewBadRequestError("ID is required")
		return validationErr
	}
	_, parseErr := uuid.Parse(id)
	if parseErr != nil {
		validationErr = response.NewBadRequestError("Invalid ID format")
		return validationErr
	}
	err := service.repository.DeleteRepository(id, ctx)
	if err != nil {
		return err
	}
	return nil
}
